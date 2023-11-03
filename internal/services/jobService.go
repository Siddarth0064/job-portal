package services

import (
	"errors"
	model "job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
)

type CompanyService interface {
	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompany(id int64) (model.Company, error)
	GetAllJobs() ([]model.Job, error)
	GetJobs(id int) ([]model.Job, error)
	JobCreate(nj model.CreateJob, id uint64) (model.Job, error)
}

func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Company{}, errors.New("user creation failed")
	}

	return cu, nil
}

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Service) GetCompany(id int64) (model.Company, error) {

	AllCompanies, err := s.c.GetCompany(id)
	if err != nil {
		return model.Company{}, err
	}
	return AllCompanies, nil

}
func (s *Service) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	job := model.Job{JobTitle: nj.JobTitle, JobSalary: nj.JobSalary, Uid: id}
	cu, err := s.c.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Job{}, errors.New("user creation failed")
	}

	return cu, nil
}
func (s *Service) GetJobs(id int) ([]model.Job, error) {
	AllCompanies, err := s.c.GetJobs(id)
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil
}
func (s *Service) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}
