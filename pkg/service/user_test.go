package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"github.com/dhanusaputra/anywhat-server/util/testutil"
	"github.com/golang/protobuf/ptypes"
)

const (
	passwordHash = "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			want: "mockToken",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "user nill",
			args: args{
				ctx:      ctx,
				username: "",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"})
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
			},
		},
		{
			name: "multiple rows",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
			},
		},
		{
			name: "query failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM user_account").WillReturnError(errors.New("err"))
			},
		},
		{
			name: "bcrypt failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password2",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
			},
		},
		{
			name: "signJWT failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "", errors.New("err")
				}
			},
		},
		{
			name: "rows failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("2", "username", passwordHash).
					RowError(0, errors.New("err"))
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
			},
		},
		{
			name: "prepare failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
				mock.ExpectPrepare("UPDATE user_account").WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "exec failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(AnyTime{}, "1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "rowsAffected failed",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(AnyTime{}, "1").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("err")))
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				ctx:      ctx,
				username: "username",
				password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password_hash"}).
					AddRow("1", "username", passwordHash)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserGet(t *testing.T) {
	ctx := context.Background()
	now := time.Now().In(time.UTC)
	tnow, err := ptypes.TimestampProto(now)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when format time", err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.User
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			want: &pb.User{
				Id:          "1",
				Username:    "username",
				CreatedAt:   tnow,
				UpdatedAt:   tnow,
				LastLoginAt: tnow,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", now, now, now)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("1").WillReturnRows(rows)
			},
		},
		{
			name: "user nill",
			args: args{
				ctx: ctx,
				id:  "id",
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"})
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("1").WillReturnRows(rows)
			},
		},
		{
			name: "multiple rows",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", now, now, now).
					AddRow("1", "username", now, now, now)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("1").WillReturnRows(rows)
			},
		},
		{
			name: "query failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM user_account").
					WillReturnError(errors.New("err"))
			},
		},
		{
			name: "scan failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", "now", "now", "now")
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("1").WillReturnRows(rows)
			},
		},
		{
			name: "rows failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("2", "username", now, now, now).
					RowError(0, errors.New("err"))
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("1").WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserList(t *testing.T) {
	ctx := context.Background()
	now := time.Now().In(time.UTC)
	tnow, err := ptypes.TimestampProto(now)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when format time", err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*pb.User
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
			},
			want: []*pb.User{&pb.User{
				Id:          "1",
				Username:    "username",
				CreatedAt:   tnow,
				UpdatedAt:   tnow,
				LastLoginAt: tnow,
			}},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", now, now, now)
				mock.ExpectQuery("SELECT (.+) FROM user_account").WillReturnRows(rows)
			},
		},
		{
			name: "user nill",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
		{
			name: "query failed",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM user_account").
					WillReturnError(errors.New("err"))
			},
		},
		{
			name: "scan failed",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", "now", "now", "now")
				mock.ExpectQuery("SELECT (.+) FROM user_account").WillReturnRows(rows)
			},
		},
		{
			name: "rows failed",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at", "last_login_at"}).
					AddRow("1", "username", now, now, now).
					AddRow("2", "username", now, now, now).
					RowError(1, errors.New("err"))
				mock.ExpectQuery("SELECT (.+) FROM user_account").WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Username: "username",
					Password: "password",
				},
			},
			want: "1",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow("1")
				mock.ExpectQuery("INSERT INTO user_account").WithArgs("username", sqlmock.AnyArg(), AnyTime{}, AnyTime{}).
					WillReturnRows(rows)
			},
		},
		{
			name: "query failed",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Username: "username",
					Password: "password",
				},
			},
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("INSERT INTO user_account").WithArgs("username", sqlmock.AnyArg(), AnyTime{}, AnyTime{}).
					WillReturnError(errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUpdate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Id:       "1",
					Username: "new username",
					Password: "new password",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(sqlmock.AnyArg(), AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: true,
		},
		{
			name: "prepare failed",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Id:       "1",
					Username: "new username",
					Password: "new password",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE user_account").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "exec failed",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Id:       "1",
					Username: "new username",
					Password: "new password",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(sqlmock.AnyArg(), AnyTime{}, "1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "rowsAffected failed",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Id:       "1",
					Username: "new username",
					Password: "new password",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(sqlmock.AnyArg(), AnyTime{}, "1").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("err")))
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				ctx: ctx,
				user: &pb.User{
					Id:       "1",
					Username: "new username",
					Password: "new password",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE user_account").ExpectExec().WithArgs(sqlmock.AnyArg(), AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDelete(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewUserService(db)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
		mock    func()
	}{
		{
			name: "happy path",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM user_account").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: true,
		},
		{
			name: "prepare failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM user_account").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "exec failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM user_account").ExpectExec().WithArgs("1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "rowsAffected failed",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM user_account").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("err")))
			},
			wantErr: true,
		},
		{
			name: "not found",
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM user_account").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT, &mock}).Restore()
			if tt.mock != nil {
				tt.mock()
			}
			got, err := s.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
