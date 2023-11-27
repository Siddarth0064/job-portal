package database

import (
	model "job-portal-api/internal/models"
	"log"
)

// ====================createiing table in the database func =======================
// func CreateTable() {
// 	// Drop the table student if it exists
// 	db, err := Open()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = db.Migrator().AutoMigrate(&model.Company{}, &model.Job{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = db.Migrator().AutoMigrate(&model.User{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// err = db.Migrator().AutoMigrate(&model.User{}, &model.NewJobRequest{}, &model.Company{}, &model.CreateCompany{}, &model.Job{}, &model.NewJobRequest{}, &model.JobType{}, &model.Location{}, &model.Qualification{}, &model.Shift{}, &model.WorkMode{})
// 	err = db.Migrator().AutoMigrate(&model.User{}, &model.Company{}, &model.Job{})
// 	if err != nil {
// 		log.Fatalln(err)
// 		// fmt.Println("error in database", err)
// 		// return

//		}
//	}
func CreateTable() {

	db, err := Open()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().DropTable(&model.User{})
	if err != nil {
		log.Fatalln(err)
	}
}
