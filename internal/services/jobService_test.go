package services

import (
	"errors"
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
		name string
		//s       *Service
		args         args
		want         model.Company
		wantErr      bool
		repoResponse func() (model.Company, error)
	}{
		{name: "success case",
			args:    args{model.CreateCompany{CompanyName: "", Adress: "bangalore", Domain: "it"}},
			want:    model.Company{},
			wantErr: true,
			repoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("invalid input for create company")
			}},
		{name: "success case",
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
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().CreateCompany(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
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
		{name: "success case",
			want:    []model.Company{{CompanyName: "tek"}, {CompanyName: "google"}},
			wantErr: false,
			repoRespons: func() ([]model.Company, error) {
				return []model.Company{{CompanyName: "tek"}, {CompanyName: "google"}}, nil
			},
		},
		{name: "failures case",
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
			if tt.repoRespons != nil {
				mockRepoCompany.EXPECT().GetAllCompany().Return(tt.repoRespons()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
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
		{name: "succes for get company",
			args:    args{id: 12},
			want:    model.Company{CompanyName: "tek", Adress: "bangalore", Domain: "IT"},
			wantErr: false,
			repoResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "tek", Adress: "bangalore", Domain: "IT"}, nil
			},
		},
		{name: "failure for get company",
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
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetCompany(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
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

func TestService_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name         string
		args         args
		want         model.Job
		wantErr      bool
		repoResponse func() (model.Job, error)
	}{
		{name: "success for jobCreate",
			args:    args{nj: model.CreateJob{JobTitle: "software", JobSalary: "5000"}, id: 2},
			want:    model.Job{JobTitle: "software", JobSalary: "5000"},
			wantErr: false,
			repoResponse: func() (model.Job, error) {
				return model.Job{JobTitle: "software", JobSalary: "5000"}, nil
			},
		},
		{name: "failure for jobCreate",
			args:    args{nj: model.CreateJob{JobTitle: "", JobSalary: "5000"}, id: 2},
			want:    model.Job{},
			wantErr: true,
			repoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("invalid enter in job Create")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepoUser := repository.NewMockUsers(mc)
			mockRepoCompany := repository.NewMockCompany(mc)
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().CreateJob(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
			if err != nil {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, err)
				return

			}
			got, err := s.JobCreate(tt.args.nj, tt.args.id)
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
		{name: "succes for get jobs",
			args:    args{id: 2},
			want:    []model.Job{{JobTitle: "software", JobSalary: "3000"}, {JobTitle: "developer", JobSalary: "30000"}},
			wantErr: false,
			repoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "software", JobSalary: "3000"}, {JobTitle: "developer", JobSalary: "30000"}}, nil
			}},
		{name: "failure for get jobs",
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
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetJobs(gomock.Any()).Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
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
		{name: "success for get all jobs",
			want:    []model.Job{{JobTitle: "software", JobSalary: "1999"}, {JobTitle: "technicall", JobSalary: "47575"}},
			wantErr: false,
			repoResponse: func() ([]model.Job, error) {
				return []model.Job{{JobTitle: "software", JobSalary: "1999"}, {JobTitle: "technicall", JobSalary: "47575"}}, nil
			}},
		{name: "failure for get all jobs",
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
			if tt.repoResponse != nil {
				mockRepoCompany.EXPECT().GetAllJobs().Return(tt.repoResponse()).AnyTimes()
			}
			s, err := NewService(mockRepoUser, mockRepoCompany)
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
