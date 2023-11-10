package repository

import (
	"context"
	"errors"
	"github.com/egorgasay/dockerdb/v3"
	_ "github.com/lib/pq"
	"go-rest-api/config"
	"go-rest-api/internal/db"
	"go-rest-api/internal/model"
	"testing"
)

func Test_userRepository_CreateUser(t *testing.T) {
	invalidErr := errors.New("invalid")
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Valid user",

			args: args{
				user: model.User{
					Username: "dima",
					Password: "test",
					Session:  "ahjsuiwlght-12",
				},
			},
			wantErr: nil,
		},
		{
			name: "BAD",
			args: args{
				user: model.User{
					Username: "anton",
				},
			},
			wantErr: invalidErr,
		},
	}

	ctx := context.TODO()
	vdb, err := dockerdb.New(ctx, dockerdb.PostgresConfig("market").Build())
	if err != nil {
		t.Fatal(err)
	}
	defer vdb.Clear(ctx)

	gormDB := db.NewDB(config.Config{DB: vdb.GetSQLConnStr()})
	if err != nil {
		t.Fatal(err)
	}
	db := userRepository{
		db: gormDB,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = db.CreateUser(tt.args.user); err != nil {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
