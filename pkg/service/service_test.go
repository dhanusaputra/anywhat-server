package service

import (
	"context"
	"database/sql/driver"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/golang/protobuf/ptypes"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	now := time.Now().In(time.UTC)
	tnow, err := ptypes.TimestampProto(now)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when format time", err)
	}
	s := NewAnywhatService(db)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       Anywhat
		args    args
		mock    func()
		want    *pb.Anything
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
					AddRow("1", "name", "description", now, now)
				mock.ExpectQuery("SELECT (.+) FROM anywhat").WithArgs("1").WillReturnRows(rows)
			},
			want: &pb.Anything{
				Id:          "1",
				Name:        "name",
				Description: "description",
				CreatedAt:   tnow,
				UpdatedAt:   tnow,
			},
		},
		{
			name: "select failed",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM ToDo").WithArgs("1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "not found",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
				mock.ExpectQuery("SELECT (.+) FROM anywhat").WithArgs("1").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "multiple rows",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
					AddRow("1", "name", "description", now, now).
					AddRow("1", "name", "description", now, now)
				mock.ExpectQuery("SELECT (.+) FROM anywhat").WithArgs("1").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("anywhat.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anywhat.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	now1 := time.Now().In(time.UTC)
	tnow1, err := ptypes.TimestampProto(now1)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when format time", err)
	}
	now2 := time.Now().In(time.UTC)
	tnow2, err := ptypes.TimestampProto(now2)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when format time", err)
	}
	s := NewAnywhatService(db)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       Anywhat
		args    args
		mock    func()
		want    []*pb.Anything
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
					AddRow("1", "name1", "description1", now1, now1).
					AddRow("2", "name2", "description2", now2, now2)
				mock.ExpectQuery("SELECT (.+) FROM anywhat").WillReturnRows(rows)
			},
			want: []*pb.Anything{
				{
					Id:          "1",
					Name:        "name1",
					Description: "description1",
					CreatedAt:   tnow1,
					UpdatedAt:   tnow1,
				},
				{
					Id:          "2",
					Name:        "name2",
					Description: "description2",
					CreatedAt:   tnow2,
					UpdatedAt:   tnow2,
				},
			},
		},
		{
			name: "empty",
			s:    s,
			args: args{
				ctx: ctx,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
				mock.ExpectQuery("SELECT (.+) FROM anywhat").WillReturnRows(rows)
			},
			want: []*pb.Anything{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("anywhat.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anywhat.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewAnywhatService(db)
	type args struct {
		ctx      context.Context
		anything *pb.Anything
	}
	tests := []struct {
		name    string
		s       Anywhat
		args    args
		mock    func()
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Name:        "name",
					Description: "description",
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("INSERT INTO anywhat").WithArgs("name", "description", AnyTime{}, AnyTime{}).WillReturnRows(rows)
			},
			want: "1",
		},
		{
			name: "query failed",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Name:        "name",
					Description: "description",
				},
			},
			mock: func() {
				mock.ExpectQuery("INSERT INTO anywhat").WithArgs("name", "description", AnyTime{}, AnyTime{}).WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Create(tt.args.ctx, tt.args.anything)
			if (err != nil) != tt.wantErr {
				t.Errorf("anywhat.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anywhat.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewAnywhatService(db)
	type args struct {
		ctx      context.Context
		anything *pb.Anything
	}
	tests := []struct {
		name    string
		s       Anywhat
		args    args
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Id:          "1",
					Name:        "new name",
					Description: "new description",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE anywhat").ExpectExec().WithArgs("new name", "new description", AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: true,
		},
		{
			name: "prepare failed",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Id:          "1",
					Name:        "new name",
					Description: "new description",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE anywhat").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "exec failed",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Id:          "1",
					Name:        "new name",
					Description: "new description",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE anywhat").ExpectExec().WithArgs("new name", "new description", AnyTime{}, "1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "rowsAffected failed",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Id:          "1",
					Name:        "new name",
					Description: "new description",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE anywhat").ExpectExec().WithArgs("new name", "new description", AnyTime{}, "1").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("err")))
			},
			wantErr: true,
		},
		{
			name: "not found",
			s:    s,
			args: args{
				ctx: ctx,
				anything: &pb.Anything{
					Id:          "1",
					Name:        "new name",
					Description: "new description",
				},
			},
			mock: func() {
				mock.ExpectPrepare("UPDATE anywhat").ExpectExec().WithArgs("new name", "new description", AnyTime{}, "1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Update(tt.args.ctx, tt.args.anything)
			if (err != nil) != tt.wantErr {
				t.Errorf("anywhat.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anywhat.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewAnywhatService(db)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       Anywhat
		args    args
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "happy path",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM anywhat").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: true,
		},
		{
			name: "prepare failed",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM anywhat").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "exec failed",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM anywhat").ExpectExec().WithArgs("1").
					WillReturnError(errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "rowsAffected failed",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM anywhat").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("err")))
			},
			wantErr: true,
		},
		{
			name: "not found",
			s:    s,
			args: args{
				ctx: ctx,
				id:  "1",
			},
			mock: func() {
				mock.ExpectPrepare("DELETE FROM anywhat").ExpectExec().WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("anywhat.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anywhat.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
