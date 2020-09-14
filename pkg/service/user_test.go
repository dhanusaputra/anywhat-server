package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"github.com/dhanusaputra/anywhat-server/util/testutil"
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
					AddRow("1", "username", "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO")
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "mockToken", nil
				}
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
					AddRow("1", "username", "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO").
					AddRow("1", "username", "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO")
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
					AddRow("1", "username", "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO")
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
					AddRow("1", "username", "$2y$10$e2d/bL85VdUak2nyPdQA/uGUW6p6s1iT4Q5lPdU00slPvp6wddssO")
				mock.ExpectQuery("SELECT (.+) FROM user_account").WithArgs("username").WillReturnRows(rows)
				authutil.SignJWT = func(user *pb.User) (string, error) {
					return "", errors.New("err")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer testutil.NewPtrs([]interface{}{&authutil.SignJWT}).Restore()
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
