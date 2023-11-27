package repository

import (
	"errors"
	//"fmt"
	model "job-portal-api/internal/models"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=jobDao.go -destination=jobDao_mock.go -package=repository

// ==================== COMPANY INTERFACE ========================================
type Company interface {
	CreateCompany(model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(id int64) (model.Company, error)
	//CreateJob(j model.Job) (model.Job, error)
	PostJob(nj model.Job) (model.Response, error)
	GetJobs(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)
	//Applyjobs(cid uint64, jobid uint64) (model.Job, error)
	FetchJobData(jid uint64) (model.Job, error)
	GetTheJobData(jobid uint) (model.Job, error)
}

//========================== CREATE COMPANY FUNC DO'S CREATING COMPANY IN THE DATABASE ================

func (r *Repo) CreateCompany(u model.Company) (model.Company, error) {
	r.db.Preload("Company")
	err := r.db.Create(&u).Error
	if err != nil {
		return model.Company{}, err
	}
	return u, nil
}

//======================== GET ALL COMPANY FUNC DO'S TO GET ALL COMPANY IN THE DATABASE ==================

func (r *Repo) GetAllCompany() ([]model.Company, error) {
	var s []model.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

//================ GET COMPANY FUNC DO'S TO GET COMPANY IN THE DATABASE =================

func (r *Repo) GetCompany(id int64) (model.Company, error) {
	var m model.Company

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Company{}, err
	}
	return m, nil

}

//	func (r *Repo) CreateJob(j model.Job) (model.Job, error) {
//		err := r.db.Create(&j).Error
//		if err != nil {
//			return model.Job{}, err
//		}
//		return j, nil
//	}
func (r *Repo) PostJob(nj model.Job) (model.Response, error) {

	res := r.db.Create(&nj).Error
	if res != nil {
		log.Info().Err(res).Send()
		return model.Response{}, errors.New("job creation failed")
	}
	return model.Response{ID: uint64(nj.ID)}, nil
}

func (r *Repo) GetJobs(id int) ([]model.Job, error) {
	var m []model.Job
	result := r.db.Preload("Company").
		Preload("Qualifications").
		Preload("Shift").
		Preload("Locations").
		Preload("JobTypes").
		Where("id=?", id).Find(&m)
	// if err != nil {
	// 	return nil, errors.New("error in preload")
	// }
	//
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, result.Error
	}
	return m, nil

}
func (r *Repo) GetAllJobs() ([]model.Job, error) {
	var s []model.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ApplyJobID takes an ApplyJob struct and compares it with job criteria.
// If the candidate matches more than 50% of the criteria, it returns a successful message.

// func (s *Repo) Applyjobs(cid, jid uint64) (model.Job, error) {

//		var job model.Job // Replace model.Job with your actual data model type.
//		tx := s.db.Preload("Qualifications").Preload("JobLocations").Preload("Technology").Preload("Shift").Preload("JobType").Preload("MaxNoticePeriod").Where("uid = ?", jid)
//		errr := tx.Find(&job).Error
//		if errr != nil {
//			return model.Job{}, errors.New("error in database to fecth the job")
//		}
//		// If the match percentage is not greater than 50%, return an error message.
//		return job, nil
//	}
func (r *Repo) FetchJobData(jid uint64) (model.Job, error) {
	var j model.Job
	result := r.db.Preload("Comp").
		Preload("Locations").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Where("id = ?", jid).
		Find(&j)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return model.Job{}, result.Error
	}

	return j, nil
}
func (r *Repo) GetTheJobData(jobid uint) (model.Job, error) {
	var jobData model.Job

	// Preload related data using GORM's Preload method
	result := r.db.Preload("Comp").
		Preload("Locations").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Where("id = ?", jobid).
		Find(&jobData)

	if result.Error != nil {

		log.Info().Err(result.Error).Send()
		return model.Job{}, result.Error
	}

	return jobData, nil
}
