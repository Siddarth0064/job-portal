package model

import (
	//"project/internal/database"

	"gorm.io/gorm"
)

// ====================== USER TABLE ================================
type User struct {
	gorm.Model
	UserName     string `json:"name" validate:"required"`
	DateOfBorn   string `json:"dob" validate:"required"`
	Email        string `json:"email" validate:"required" `
	PasswordHash string `json:"-" validate:"required"`
}

// ====================== USER SIGN UP FIELDS ==========================
type UserSignup struct {
	UserName   string `json:"name" validate:"required"`
	DateOfBorn string `json:"dob" validate:"required"`
	Email      string `json:"email" validate:"required" `
	Password   string `json:"password" validate:"required"`
}

// ====================== USER LOGIN  FIELDS =============================
type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgetPass struct {
	Email      string `json:"email" validate:"required"`
	DateOfBorn string `json:"dob" `
}
type ChnagePass struct {
	Email           string `json:"email" validate:"required"`
	Otp             string `json:"otp" validate:"required"`
	NewPassword     string `json:"newpassword" validate:"required"`
	ComfirmPassword string `json:"comfirmpassword" validate:"required"`
}
