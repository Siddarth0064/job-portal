package model

import "gorm.io/gorm"

//====================== JOB TABLE =====================================
type Job struct {
	gorm.Model
	JobTitle  string  `json:"job_title" validate:"required"`
	JobSalary string  `json:"job_salary" validate:"required"`
	Company   Company `gorm:"ForeignKey:uid"`
	Uid       uint64  `JSON:"uid, omitempty"`
}

//============================= COMPANY TABLE =============================
type Company struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
	//	Job         []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyID"`
}

//=================== CREATE COMPANY FIELDS ==========================
type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

//================== CREATE JOB FIELDS ==============================
type CreateJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
}
