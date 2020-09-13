package service

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Anywhat ...
type Anywhat interface {
	Get(ctx context.Context, id string) (*pb.Anything, error)
	List(ctx context.Context) ([]*pb.Anything, error)
	Create(ctx context.Context, anything *pb.Anything) (string, error)
	Update(ctx context.Context, anything *pb.Anything) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type anywhatService struct {
	db *sql.DB
}

// NewAnywhatService ...
func NewAnywhatService(db *sql.DB) Anywhat {
	return &anywhatService{db}
}

func (s *anywhatService) Get(ctx context.Context, id string) (*pb.Anything, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, updated_at FROM anywhat WHERE id=$1", id)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query anything, id: %s, err: %s", id, err.Error())
	}
	defer rows.Close()

	for !rows.Next() {
		if rows.Err() != nil {
			return nil, status.Errorf(codes.Unknown, "failed to retrieve data from anything, err: %s", err.Error())
		}
		return nil, status.Errorf(codes.NotFound, "anything with ID='%s' is not found", id)
	}

	var a *pb.Anything
	var createdAt, updatedAt time.Time
	if err := rows.Scan(&a.Id, &a.Name, &a.Description, &createdAt, &updatedAt); err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to retrieve field values from anything, err: %s", err.Error())
	}
	a.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "createdAt field has invalid format, err: %s", err.Error())
	}
	a.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "updatedAt field has invalid format, err: %s", err.Error())
	}

	if rows.Next() {
		return nil, status.Errorf(codes.Unknown, "found multiple rows with ID='%s'", id)
	}

	return a, nil
}

func (s *anywhatService) List(ctx context.Context) ([]*pb.Anything, error) {
	rows, err := s.db.Query("SELECT id, name, description, created_at, updated_at FROM anywhat")
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to query anythings, err: %s", err.Error())
	}
	defer rows.Close()

	var createdAt, updatedAt time.Time
	res := []*pb.Anything{}
	for rows.Next() {
		var a pb.Anything
		if err := rows.Scan(&a.Id, &a.Name, &a.Description, &createdAt, &updatedAt); err != nil {
			return nil, status.Errorf(codes.Unknown, "failed to retrieve field values from anything, err: %s", err.Error())
		}
		a.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "createdAt field has invalid format, err: %s", err.Error())
		}
		a.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "updatedAt field has invalid format, err: %s", err.Error())
		}
		res = append(res, &a)
	}

	if rows.Err() != nil {
		return nil, status.Errorf(codes.Unknown, "failed to retrieve data from anythings, err: %s", err.Error())
	}

	return res, nil
}

func (s *anywhatService) Create(ctx context.Context, anything *pb.Anything) (string, error) {
	now := time.Now().In(time.UTC)

	var id int
	err := s.db.QueryRow("INSERT INTO anywhat(name,description,created_at,updated_at) VALUES($1,$2,$3,$4) returning id;", anything.Name, anything.Description, now, now).Scan(&id)
	if err != nil {
		return "", status.Errorf(codes.Unknown, "failed to insert into anything, err: %s", err.Error())
	}

	return strconv.Itoa(id), nil
}

func (s *anywhatService) Update(ctx context.Context, anything *pb.Anything) (bool, error) {
	stmt, err := s.db.Prepare("UPDATE anywhat SET name=$1, description=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to prepare update anything, err: %s", err.Error())
	}

	res, err := stmt.Exec(anything.Name, anything.Description, time.Now().In(time.UTC), anything.Id)
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to update anything, err: %s", err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to retrieve rows affected value, err: %s ", err.Error())
	}

	if rows == 0 {
		return false, status.Errorf(codes.NotFound, "anywhat with ID='%s' is not found", anything.Id)
	}

	return true, nil
}

func (s *anywhatService) Delete(ctx context.Context, id string) (bool, error) {
	stmt, err := s.db.Prepare("DELETE FROM anywhat WHERE id=$1")
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to prepare delete anything, err: %s", err.Error())
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to delete anything, err: %s", err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, status.Errorf(codes.Unknown, "failed to retrieve rows affected value, err: %s ", err.Error())
	}

	if rows == 0 {
		return false, status.Errorf(codes.NotFound, "anywhat with ID='%s' is not found", id)
	}

	return true, nil
}
