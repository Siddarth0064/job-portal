package services

import (
	"errors"
	model "job-portal-api/internal/models"
	"reflect"
	"testing"
)

type MockJobs struct{}

func (m *MockJobs) CreateJob(j model.Job) (model.Job, error) {
	if j.JobTitle == "" {
		return model.Job{}, errors.New("job creation failed")
	}
	if j.JobSalary == "" {
		return model.Job{}, errors.New("job creation failed")
	}

	return model.Job{JobTitle: j.JobTitle, JobSalary: j.JobSalary, Uid: j.Uid}, nil
}
func (m *MockJobs) GetJobs(id int) ([]model.Job, error) {
	if id == 5 {
		return []model.Job{{JobTitle: "developer", JobSalary: "20k", Uid: 5}}, nil
	}

	return nil, errors.New("job retreval failed")
}
func (m *MockJobs) CreateCompany(model.Company) (model.Company, error) { return model.Company{}, nil }
func (m *MockJobs) GetAllCompany() ([]model.Company, error)            { return nil, nil }
func (m *MockJobs) GetCompany(id int64) (model.Company, error)         { return model.Company{}, nil }
func (m *MockJobs) GetAllJobs() ([]model.Job, error)                   { return nil, nil }

// func (m *MockJobs) CreateJob(model.Job)(model.Job,error){return model.Job{},nil}
func TestService_CreateJob(t *testing.T) {
	type args struct {
		nc model.CreateJob
		id uint64
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    model.Job
		wantErr bool
	}{
		{name: "test case 1",
			s:       &Service{c: &MockJobs{}},
			args:    args{nc: model.CreateJob{JobTitle: "developer", JobSalary: "20k"}, id: 5},
			want:    model.Job{JobTitle: "developer", JobSalary: "20k", Uid: 5},
			wantErr: true,
		},
		{name: "test case 2",
			s:       &Service{c: &MockJobs{}},
			args:    args{nc: model.CreateJob{JobTitle: "", JobSalary: ""}, id: 5},
			want:    model.Job{},
			wantErr: false,
		},
		{name: "test case 3",
			s:       &Service{c: &MockJobs{}},
			args:    args{nc: model.CreateJob{JobTitle: "", JobSalary: "20k"}, id: 5},
			want:    model.Job{},
			wantErr: false,
		},
		{name: "test case 4",
			s:       &Service{c: &MockJobs{}},
			args:    args{nc: model.CreateJob{JobTitle: "developer", JobSalary: ""}, id: 5},
			want:    model.Job{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.JobCreate(tt.args.nc, tt.args.id)
			if (err == nil) != tt.wantErr {
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
		name    string
		s       *Service
		args    args
		want    []model.Job
		wantErr bool
	}{
		{
			name:    "test case 1",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 5},
			want:    []model.Job{{JobTitle: "developer", JobSalary: "20k", Uid: 5}},
			wantErr: true,
		},
		{name: "test case 2",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 1},
			want:    nil,
			wantErr: false,
		},
		{name: "test case 3",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 4},
			want:    nil,
			wantErr: false,
		},
		{name: "test case 4",
			s:       &Service{c: &MockJobs{}},
			args:    args{id: 8},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetJobs(tt.args.id)
			if (err == nil) != tt.wantErr {
				t.Errorf("Service.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
