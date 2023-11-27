package services

import (
	//"context"
	"context"
	"encoding/json"
	"errors"

	model "job-portal-api/internal/models"

	//reflect "reflect"

	//"slices"

	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// ================= COMPANY SERVICE INTERFACE ================================

//go:generate mockgen -source=jobService.go -destination=jobService_mock.go -package=services
type CompanyService interface {
	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompany(id int64) (model.Company, error)
	GetAllJobs() ([]model.Job, error)
	GetJobs(id int) ([]model.Job, error)
	JobCreate(newJob model.NewJobRequest, id uint64) (model.Response, error)
	ProccessApplication(ctx context.Context, applicationData []model.NewUserApplication) ([]model.NewUserApplication, error)
}

// ======================= COMAPANY CREATE FUNC IS USED TO CREATE COMPANY INFORMATION IN DATABASE ===============
func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Company{}, errors.New("user creation failed")
	}

	return cu, nil
}

//=============== GET ALL COMPANIES IS USED TO GET ALL COMPANIES IN THE DATABASE ====================

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil

}

// ===================== GET COMPANY FUNC IS USED TO GET COMPANY DATA IN THE DATRABASE =======================
func (s *Service) GetCompany(id int64) (model.Company, error) {

	AllCompanies, err := s.c.GetCompany(id)
	if err != nil {
		return model.Company{}, err
	}
	return AllCompanies, nil

}

// ================= job create func ======================================================
func (s *Service) JobCreate(newJob model.NewJobRequest, id uint64) (model.Response, error) {

	app := model.Job{
		CompanyId:           id,
		JobTitle:            newJob.JobTitle,
		Salary:              newJob.Salary,
		MinimumNoticePeriod: newJob.MinimumNoticePeriod,
		MaximumNoticePeriod: newJob.MaximumNoticePeriod,
		Budget:              newJob.Budget,
		JobDescription:      newJob.JobDescription,
		MinExperience:       newJob.MinExperience,
	}
	for _, v := range newJob.QualificationIDs {
		tempData := model.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Qualifications = append(app.Qualifications, tempData)
	}
	for _, v := range newJob.LocationIDs {
		tempData := model.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Locations = append(app.Locations, tempData)
	}
	for _, v := range newJob.SkillIDs {
		tempData := model.Skill{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Skills = append(app.Skills, tempData)
	}
	for _, v := range newJob.WorkModeIDs {
		tempData := model.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.WorkModes = append(app.WorkModes, tempData)
	}
	for _, v := range newJob.ShiftIDs {
		tempData := model.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.Shifts = append(app.Shifts, tempData)
	}
	for _, v := range newJob.JobTypeIDs {
		tempData := model.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		app.JobTypes = append(app.JobTypes, tempData)
	}
	jobData, err := s.c.PostJob(app)
	if err != nil {
		return model.Response{}, err
	}
	return jobData, nil
}

// ========================= GET JOBS FUNC IS USED TO GET JOBS IN THE SINGLE COMPANYES =========================
func (s *Service) GetJobs(id int) ([]model.Job, error) {
	AllCompanies, err := s.c.GetJobs(id)
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil
}

// =================== GET ALL JOBS FUNC IS USED TO GET ALL JOBS IN THE ALL COMPANIES ==========================
func (s *Service) GetAllJobs() ([]model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

// ================ process job  applications  func ==============================================
func (s *Service) ProccessApplication(ctx context.Context, applicationData []model.NewUserApplication) ([]model.NewUserApplication, error) {
	var wg = new(sync.WaitGroup)
	ch := make(chan model.NewUserApplication)
	var finalData []model.NewUserApplication

	for _, v := range applicationData {
		wg.Add(1)
		go func(v model.NewUserApplication) {
			defer wg.Done()
			var jobData model.Job

			val, err := s.rdb.GetCacheData(ctx, uint(v.ID))

			if err != nil {

				dbData, err := s.c.GetTheJobData(uint(v.ID))
				if err != nil {

					return
				}

				err = s.rdb.AddToCache(ctx, uint(v.ID), dbData)
				if err != nil {

					return
				}
				jobData = dbData
			} else {

				err = json.Unmarshal([]byte(val), &jobData)

				if err == redis.Nil {

					return
				}
				if err != nil {

					return
				}
			}

			check := Compare(v, jobData)
			if check {

				ch <- v
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {

		finalData = append(finalData, v)
	}

	return finalData, nil
}
func Compare(applicationData model.NewUserApplication, val model.Job) bool {

	if applicationData.Jobs.Experience < val.MinExperience {
		return false
	}

	if applicationData.Jobs.NoticePeriod < val.MinimumNoticePeriod {
		return false
	}
	var count int

	count = compareLocations(applicationData.Jobs.Location, val.Locations)
	if count == 0 {
		return false
	}

	count = compareQualifications(applicationData.Jobs.Qualifications, val.Qualifications)
	if count == 0 {
		return false
	}

	count = compareTechStack(applicationData.Jobs.WorkModeIDs, val.WorkModes)
	if count == 0 {
		return false
	}

	count = compareShifts(applicationData.Jobs.Shift, val.Shifts)
	if count == 0 {
		return false
	}

	return true
}

func compareLocations(locationsID []uint, val []model.Location) int {

	count := 0
	for _, v := range locationsID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareQualifications(qualificationID []uint, val []model.Qualification) int {
	count := 0
	for _, v := range qualificationID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareTechStack(stackID []uint, val []model.WorkMode) int {
	count := 0
	for _, v := range stackID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareShifts(shiftID []uint, val []model.Shift) int {
	count := 0
	for _, v := range shiftID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}
