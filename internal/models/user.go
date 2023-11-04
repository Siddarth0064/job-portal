package model

import (
	//"project/internal/database"

	"gorm.io/gorm"
)

// ====================== USER TABLE ================================
type User struct {
	gorm.Model
	UserName     string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required" `
	PasswordHash string `json:"-" validate:"required"`
}

// ====================== USER SIGN UP FIELDS ==========================
type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required" `
	Password string `json:"password" validate:"required"`
}

// ====================== USER LOGIN  FIELDS =============================
type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
