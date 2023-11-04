package repository

import (
	"errors"
	model "job-portal-api/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

// ===================== NEW REPO FUNC IS USED TO INITIALIZE TO REPO STRUCT =================

func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("db connection not given")
	}

	return &Repo{db: db}, nil

}

// ====================== USERS INTERFACE =====================================
//
//go:generate mockgen -source=userDao.go -destination=userDao_mock.go -package=repository
type Users interface {
	CreateUser(model.User) (model.User, error)
	FetchUserByEmail(string) (model.User, error)
}

// ===================== CREATE USER FUNC IS USED TO CREATE USER INFORMATION ===========================
func (r *Repo) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

// ======================= FETCH USER BY EMAIL FUNC IS USED TO FETCH USER DATA BY EMAIL ======================
func (r *Repo) FetchUserByEmail(s string) (model.User, error) {
	var u model.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.User{}, nil
	}
	return u, nil

}
