package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// User ...
type User interface {
	Login(ctx context.Context, username, password string) (string, error)
	Me(ctx context.Context) (*pb.User, error)
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
		return "", status.Error(codes.PermissionDenied, fmt.Sprintf("failed to login, err: %v", err))
	}

	token, err := authutil.SignJWT(&pb.User{
		Id:       u.Id,
		Username: u.Username,
	})
	if err != nil {
		return "", status.Error(codes.Unknown, fmt.Sprintf("failed to login, err: %v", err))
	}

	return token, nil
}

func (s *userService) Me(ctx context.Context) (*pb.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unknown, "metadata is required")
	}
	auth := md.Get("authorization")
	if len(auth[0]) == 0 {
		return nil, status.Error(codes.InvalidArgument, "authorization is required")
	}
	_, claims, err := authutil.ValidateJWT(auth[0])
	if err != nil {
		log.Print(err.Error())
	}
	return &pb.User{
		Id:       claims["id"].(string),
		Username: claims["username"].(string),
	}, nil
}