package service

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User ...
type User interface {
	Login(ctx context.Context, username, password string) (string, error)
	Get(ctx context.Context, id string) (*pb.User, error)
	List(ctx context.Context) ([]*pb.User, error)
	Create(ctx context.Context, user *pb.User) (string, error)
	Update(ctx context.Context, user *pb.User) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type userService struct {
	db *sql.DB
}

// NewUserService ...
func NewUserService(db *sql.DB) User {
	return &userService{db}
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	if len(username) == 0 || len(password) == 0 {
		return "", status.Error(codes.InvalidArgument, "username and pasword are required")
	}

	rows, err := s.db.Query("SELECT id, username, password_hash FROM user_account WHERE username=$1", username)
	if err != nil {
		return "", status.Errorf(codes.Unknown, "failed to query user_account, username: %s, err: %s", username, err.Error())
	}
	defer rows.Close()

	for !rows.Next() {
		if rows.Err() != nil {
			return "", status.Errorf(codes.Unknown, "failed to retrieve data from user_account, err: %s", err.Error())
		}
		return "", status.Errorf(codes.NotFound, "user_account with username: '%s' is not found", username)
	}

	var u pb.User
	if err := rows.Scan(&u.Id, &u.Username, &u.PasswordHash); err != nil {
		return "", status.Errorf(codes.Unknown, "failed to retrieve field values from user_account, err: %s", err.Error())
	}

	if rows.Next() {
		return "", status.Errorf(codes.Unknown, "found multiple rows with username: '%s'", username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return "", status.Errorf(codes.PermissionDenied, "failed to login, err: %v", err)
	}

	token, err := authutil.SignJWT(&pb.User{
		Id:       u.Id,
		Username: u.Username,
	})
	if err != nil {
		return "", status.Errorf(codes.Unknown, "failed to login, err: %v", err)
	}

	return token, nil
}

func (s *userService) Get(ctx context.Context, id string) (*pb.User, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, updated_at FROM user WHERE id=$1", id)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query user, id: %s, err: %s", id, err.Error())
	}
	defer rows.Close()

	for !rows.Next() {
		if rows.Err() != nil {
			return nil, status.Errorf(codes.Unknown, "failed to retrieve data from user, err: %s", err.Error())
		}
		return nil, status.Errorf(codes.NotFound, "user with ID: '%s' is not found", id)
	}

	var a pb.Anything
	var createdAt, updatedAt time.Time
	if err := rows.Scan(&a.Id, &a.Name, &a.Description, &createdAt, &updatedAt); err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to retrieve field values from user, err: %s", err.Error())
	}
	a.CreatedAt = timestamppb.New(createdAt)
	a.UpdatedAt = timestamppb.New(updatedAt)

	if rows.Next() {
		return nil, status.Errorf(codes.Unknown, "found multiple rows with ID: '%s'", id)
	}

	return nil, nil
}

func (s *userService) List(ctx context.Context) ([]*pb.User, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, updated_at FROM user")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query user, err: %s", err.Error())
	}
	defer rows.Close()

	var createdAt, updatedAt time.Time
	res := []*pb.Anything{}
	for rows.Next() {
		var a pb.Anything
		if err := rows.Scan(&a.Id, &a.Name, &a.Description, &createdAt, &updatedAt); err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to retrieve field values from user, err: %s", err.Error())
		}
		a.CreatedAt = timestamppb.New(createdAt)
		a.UpdatedAt = timestamppb.New(updatedAt)

		res = append(res, &a)
	}

	if rows.Err() != nil {
		return nil, status.Errorf(codes.Unknown, "failed to retrieve data from users, err: %s", err.Error())
	}

	return nil, nil
}

func (s *userService) Create(ctx context.Context, user *pb.User) (string, error) {
	now := time.Now().In(time.UTC)

	var id int
	err := s.db.QueryRow("INSERT INTO user(username,password,created_at,updated_at) VALUES($1,$2,$3,$4) returning id;", user.Username, user.Password, now, now).Scan(&id)
	if err != nil {
		return "", status.Errorf(codes.Unknown, "failed to insert into user, err: %s", err.Error())
	}

	return strconv.Itoa(id), nil
}

func (s *userService) Update(ctx context.Context, user *pb.User) (bool, error) {
	stmt, err := s.db.Prepare("UPDATE user SET name=$1, description=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to prepare update user, err: %s", err.Error())
	}

	res, err := stmt.Exec(user.Username, user.Password, time.Now().In(time.UTC), user.Id)
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to update user, err: %s", err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to retrieve rows affected value, err: %s ", err.Error())
	}

	if rows == 0 {
		return false, status.Errorf(codes.NotFound, "user with ID: '%s' is not found", user.Id)
	}

	return true, nil
}

func (s *userService) Delete(ctx context.Context, id string) (bool, error) {
	stmt, err := s.db.Prepare("DELETE FROM user WHERE id=$1")
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to prepare delete user, err: %s", err.Error())
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to delete user, err: %s", err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to retrieve rows affected value, err: %s ", err.Error())
	}

	if rows == 0 {
		return false, status.Errorf(codes.NotFound, "user with ID: '%s' is not found", id)
	}

	return true, nil
}
