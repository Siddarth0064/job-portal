package services

import (
	"errors"
	model "job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"
	"time"

	//"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func TestService_UserSignup(t *testing.T) {
	type args struct {
		nu model.UserSignup
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             model.User
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{name: "SUCCESS====CASE 1",
			args:    args{model.UserSignup{UserName: "siddarth", Email: "Siddarth@gmail.com", Password: "sidd@12"}},
			want:    model.User{UserName: "siddarth", Email: "Siddarth@gmail.com"},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{UserName: "siddarth", Email: "Siddarth@gmail.com"}, nil
			}},
		{name: "SUCCESS====CASE 2",
			args:    args{model.UserSignup{UserName: "siddarth", Email: "Siddarth@gmail.com", Password: "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"}},
			want:    model.User{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("error in password")
			}},
		{name: "FAILURE======CASE 1",
			args:    args{model.UserSignup{UserName: "", Email: "Siddarth@gmail.com", Password: "sidd@12"}},
			want:    model.User{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("invalid user input for signup")
			}},
		{name: "FAILURE======CASE 2",
			args:    args{model.UserSignup{UserName: "siddarth", Email: "Siddarth@gmail.com"}},
			want:    model.User{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("invalid user input for signup")
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUsers(mc)
			mockRepoCom := repository.NewMockCompany(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepo, mockRepoCom)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.UserSignup(tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Userlogin(t *testing.T) {
	type args struct {
		l model.UserLogin
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{name: "FAILURE======CASE 1",
			args:    args{model.UserLogin{Email: "siddarth@gmail.com", Password: ""}},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("error")
			}},
		{name: "FAILURE======CASE2",
			args:    args{model.UserLogin{Email: "siddarthAgmail.com", Password: ""}},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{Email: "siddarthAgmail.com", PasswordHash: "$2a$10$votXUqKwkXe6l5.2aVKSU.08QEPzZYuXy47OP7JuHebrZSppBlYSW"}, nil
			}},

		{name: "SUCCESS====CASE 1",
			args:    args{model.UserLogin{Email: "siddarthAgmail.com", Password: "hfhhfhfh"}},
			want:    jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{Email: "siddarthAgmail.com", PasswordHash: "$2a$10$votXUqKwkXe6l5.2aVKSU.08QEPzZYuXy47OP7JuHebrZSppBlYSW"}, nil
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRespondUser := repository.NewMockUsers(mc)
			mockRespondCom := repository.NewMockCompany(mc)
			if tt.mockRepoResponse != nil {
				mockRespondUser.EXPECT().FetchUserByEmail(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, err := NewService(mockRespondUser, mockRespondCom)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.Userlogin(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Userlogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Userlogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
