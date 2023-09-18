package usecase

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"

	"sample/pkg/domain/model"
	"sample/pkg/domain/repository"
)

type mockUserRepository struct {
	repository.UserRepository
}

func newMockUserRepository() repository.UserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) List(c echo.Context, limit int, offset int) (*model.Users, error) {
	println("mockUserRepository list")
	return &model.Users{}, nil
}
func (m *mockUserRepository) FindById(c echo.Context, userID int) (*model.User, error) {
	return &model.User{}, nil
}
func (m *mockUserRepository) Create(c echo.Context, user *model.User) (*model.User, error) {
	return &model.User{}, nil
}

func Test_userUseCase_GetUserList(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		c      echo.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Users
		wantErr bool
	}{
		{name: "success",
			fields: fields{
				UserRepository: newMockUserRepository(),
			},
			args: args{
				c:      nil,
				limit:  10,
				offset: 0,
			},
			want:    &model.Users{},
			wantErr: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userUseCase{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.GetUserList(tt.args.c, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.GetUserList(%v, %v, %v) error = %v, wantErr %v", tt.args.c, tt.args.limit, tt.args.offset, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.GetUserList(%v, %v, %v) = %v, want %v", tt.args.c, tt.args.limit, tt.args.offset, got, tt.want)
			}
		})
	}
}

func Test_userUseCase_GetUserDetails(t *testing.T) {
	println("success")

}
