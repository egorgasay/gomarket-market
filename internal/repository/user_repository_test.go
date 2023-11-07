package repository

import (
	"go-rest-api/internal/model"
	"gorm.io/gorm"
	"testing"
)

func Test_userRepository_CreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			if err := ur.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
