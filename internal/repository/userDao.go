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
	UpdateUser(email string, updatedUser model.User) (model.User, error)
	FetchUserByEmail(string) (model.User, error)
	FetchUserEmail(s string) (model.ForgetPass, error)
	FetchUserByDob(string) error
}

// ===================== CREATE USER FUNC IS USED TO CREATE USER INFORMATION ===========================
func (r *Repo) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

// ==================== UPDATE USER =============================================
func (r *Repo) UpdateUser(email string, updatedUser model.User) (model.User, error) {
	// Check if the user with the given email exists
	existingUser := model.User{}
	err := r.db.Where("email = ?", email).First(&existingUser).Error
	if err != nil {
		return model.User{}, err
	}

	existingUser.PasswordHash = updatedUser.PasswordHash
	err = r.db.Save(&existingUser).Error
	if err != nil {
		return model.User{}, err
	}

	return existingUser, nil
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

// ============================================================================================
func (r *Repo) FetchUserEmail(s string) (model.ForgetPass, error) {
	var u model.ForgetPass
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.ForgetPass{}, errors.New("email is not found or not match")
	}
	return u, nil

}

// ============================ FETCH USER BY DOB =======================================
func (r *Repo) FetchUserByDob(s string) error {
	var u string
	tx := r.db.Where("dob=?", s).First(&u)
	if tx.Error != nil {
		return errors.New("DOB not found in database")
	}
	return nil

}
