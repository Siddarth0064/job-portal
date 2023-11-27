package services

import (
	"errors"
	"job-portal-api/internal/cache"
	model "job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_CompanyCreate(t *testing.T) {
	type args struct {
		nc model.CreateCompany
	}
	tests := []struct {
		name         string
		args         args
		want         model.Company
		wantErr      bool
		repoResponse func() (model.Company, error)
	}{
		{name: "FAILURE======CASE 1",
			args:    args{model.CreateCompany{CompanyName: "", Adress: "bangalore", Domain: "it"}},
			want:    model.Company{},
			wantErr: true,
			repoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("invalid input for create company")
			}},
		{name: "SUCCESS====CASE 1",
			args:    args{model.CreateCompany{CompanyName: "tek", Adress: "bangalore", Domain: "it"}},
			want:    model.Company{CompanyName: "tek", Adress: "bangs", Domain: "it"},
			wantErr: false,
			repoResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "tek", Adress: "bangs", Domain: "it"}, nil
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().CreateCompany(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.CompanyCreate(tt.args.nc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CompanyCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CompanyCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllCompanies(t *testing.T) {
	tests := []struct {
		name string

		want        []model.Company
		wantErr     bool
		repoRespons func() ([]model.Company, error)
	}{
		{name: "SUCCESS====CASE 1",
			want:    []model.Company{{CompanyName: "tek"}, {CompanyName: "google"}},
			wantErr: false,
			repoRespons: func() ([]model.Company, error) {
				return []model.Company{{CompanyName: "tek"}, {CompanyName: "google"}}, nil
			},
		},
		{name: "FAILURE======CASE 1",
			want:    nil,
			wantErr: true,
			repoRespons: func() ([]model.Company, error) {
				return nil, errors.New("records not founds")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)

			if tt.repoRespons != nil {
				mockRepoCompany.EXPECT().GetAllCompany().Return(tt.repoRespons()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.GetAllCompanies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetCompany(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string

		args         args
		want         model.Company
		wantErr      bool
		repoResponse func() (model.Company, error)
	}{
		{name: "SUCCESS====CASE 1",
			args:    args{id: 12},
			want:    model.Company{CompanyName: "tek", Adress: "bangalore", Domain: "IT"},
			wantErr: false,
			repoResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "tek", Adress: "bangalore", Domain: "IT"}, nil
			},
		},
		{name: "FAILURE======CASE 1",
			args:    args{},
			want:    model.Company{},
			wantErr: true,
			repoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("invalid id ")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)

			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetCompany(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.GetCompany(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestService_JobCreate(t *testing.T) {
// 	type args struct {
// 		nj model.NewJobRequest
// 		id uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		want         model.Job
// 		wantErr      bool
// 		repoResponse func() (model.Job, error)
// 	}{
// 		{name: "SUCCESS====CASE 1",
// 			args:    args{nj: model.NewJobRequest{JobTitle: "software", Salary: "5000"}, id: 2},
// 			want:    model.Job{JobTitle: "software", Salary: "5000"},
// 			wantErr: false,
// 			repoResponse: func() (model.Job, error) {
// 				return model.Job{JobTitle: "software", Salary: "5000"}, nil
// 			},
// 		},
// 		{name: "FAILURE======CASE 1",
// 			args:    args{nj: model.NewJobRequest{JobTitle: "", Salary: "5000"}, id: 2},
// 			want:    model.Job{},
// 			wantErr: true,
// 			repoResponse: func() (model.Job, error) {
// 				return model.Job{}, errors.New("invalid enter in job Create")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepoUser := repository.NewMockUsers(mc)
// 			mockRepoCompany := repository.NewMockCompany(mc)
// 			if tt.repoResponse != nil {
// 				mockRepoCompany.EXPECT().PostJob(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
// 			}
// 			s, err := NewService(mockRepoUser, mockRepoCompany)
// 			if err != nil {
// 				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
// 				return

// 			}
// 			got, err := s.JobCreate(tt.args.nj, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.JobCreate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestService_GetJobs(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name         string
		args         args
		want         []model.Job
		wantErr      bool
		repoResponse func() ([]model.Job, error)
	}{
		{name: "SUCCESS====CASE 1",
			args:    args{id: 2},
			want:    []model.Job{{JobTitle: "software", Salary: "3000"}, {JobTitle: "developer", Salary: "30000"}},
			wantErr: false,
			repoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "software", Salary: "3000"}, {JobTitle: "developer", Salary: "30000"}}, nil
			}},
		{name: "FAILURE======CASE 1",
			args:    args{id: 2},
			want:    nil,
			wantErr: true,
			repoResponse: func() ([]model.Job, error) {
				return nil, errors.New("invalid jobs id")
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)

			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetJobs(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.GetJobs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllJobs(t *testing.T) {
	tests := []struct {
		name         string
		want         []model.Job
		wantErr      bool
		repoResponse func() ([]model.Job, error)
	}{
		{name: "SUCCESS====CASE 1",
			want:    []model.Job{{JobTitle: "software", Salary: "1999"}, {JobTitle: "technicall", Salary: "47575"}},
			wantErr: false,
			repoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "software", Salary: "1999"}, {JobTitle: "technicall", Salary: "47575"}}, nil
			}},
		{name: "FAILURE======CASE 1",
			want:    nil,
			wantErr: true,
			repoResponse: func() ([]model.Job, error) {
				return nil, errors.New("records not founds")
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)

			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetAllJobs().Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.GetAllJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_JobCreate(t *testing.T) {

	type args struct {
		newJob model.NewJobRequest
		id     uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Response
		wantErr          bool
		mockRepoRespones func() (model.Response, error)
	}{
		{name: "success case 1",
			args: args{model.NewJobRequest{JobTitle: "software developer",
				Salary:              "4lpa",
				MinimumNoticePeriod: 0,
				MaximumNoticePeriod: 3,
				Budget:              500.00,
				JobDescription:      "go lang backend developer",
				MinExperience:       2.3,
				MaxExperience:       4.5,
				QualificationIDs:    []uint{uint(1), uint(2)},
				LocationIDs:         []uint{uint(1), uint(2)},
				SkillIDs:            []uint{uint(1), uint(2)},
				WorkModeIDs:         []uint{uint(1), uint(2)},
				ShiftIDs:            []uint{uint(1), uint(2)},
				JobTypeIDs:          []uint{uint(1), uint(2)},
			}, 45},
			want:    model.Response{ID: 1},
			wantErr: false,
			mockRepoRespones: func() (model.Response, error) {
				return model.Response{ID: 1}, nil
			},
		},
		{name: "failure  case 1",
			args: args{model.NewJobRequest{JobTitle: "software developer",
				Salary:              "4lpa",
				MinimumNoticePeriod: 0,
				MaximumNoticePeriod: 3,
				Budget:              500.00,
				JobDescription:      "go lang backend developer",
				MinExperience:       2.3,
				MaxExperience:       4.5,
				QualificationIDs:    []uint{uint(1), uint(2)},
				LocationIDs:         []uint{uint(1), uint(2)},
				SkillIDs:            []uint{uint(1), uint(2)},
				WorkModeIDs:         []uint{uint(1), uint(2)},
				ShiftIDs:            []uint{uint(1), uint(2)},
				JobTypeIDs:          []uint{uint(1), uint(2)},
			}, 45},
			want:    model.Response{},
			wantErr: true,
			mockRepoRespones: func() (model.Response, error) {
				return model.Response{}, errors.New("error in creation of job")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			mockRepoRedis := cache.NewMockCaching(mc)

			mockRepoCompany.EXPECT().PostJob(gomock.Any()).Return(tt.mockRepoRespones()).AnyTimes()
			s, err := NewService(mockRepoUser, mockRepoCompany, mockRepoRedis)
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := s.JobCreate(tt.args.newJob, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.JobCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
