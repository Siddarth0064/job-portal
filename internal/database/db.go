package database

import (
	"context"
	"fmt"
	"job-portal-api/config"
	"job-portal-api/internal/models"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// // ============open func acts as a connecting to the database==========================
// func Open() (*gorm.DB, error) {
// 	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// // ====================connection func do's all the connection to the database===================
// func Connection() (*gorm.DB, error) {
// 	log.Info().Msg("main : Started : Initializing db support")
// 	db, err := Open()
// 	if err != nil {
// 		return nil, fmt.Errorf("connecting to db %w", err)
// 	}
// 	pg, err := db.DB()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get database instance %w ", err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()

//		err = pg.PingContext(ctx)
//		if err != nil {
//			return nil, fmt.Errorf("database is not connected: %w ", err)
//		}
//		//	db.Migrator().DropTable(&model.Job{}, &model.Company{}, &model.User{}, &model.Shift{}, &model.Location{})
//		err = db.Migrator().AutoMigrate(&model.User{})
//		if err != nil {
//			return nil, fmt.Errorf("auto migration failed: %w ", err)
//		}
//		err = db.Migrator().AutoMigrate(&model.Company{})
//		if err != nil {
//			return nil, fmt.Errorf("auto migration failed: %w ", err)
//		}
//		err = db.Migrator().AutoMigrate(&model.Job{})
//		if err != nil {
//			return nil, fmt.Errorf("auto migration failed: %w ", err)
//		}
//		return db, nil
//	}
func Open() (*gorm.DB, error) {
	//dsn := os.Getenv("DB_DSN")
	cfg := config.GetConfig()
	dsn := cfg.DatabaseConfing.DBConnection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connection() (*gorm.DB, error) {
	log.Info().Msg("main : Started : Initializing db support")
	db, err := Open()
	if err != nil {
		return nil, fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("database is not connected: %w ", err)
	}
	// db.Migrator().DropTable(&model.User{}, &model.Company{}, &model.Job{})
	// err = db.Migrator().AutoMigrate(&model.User{})
	// if err != nil {
	// 	return nil, fmt.Errorf("auto migration failed: %w ", err)
	// }
	// err = db.Migrator().AutoMigrate(&model.Company{})
	// if err != nil {
	// 	return nil, fmt.Errorf("auto migration failed: %w ", err)
	// }
	err = db.Migrator().AutoMigrate(&model.Job{}, &model.User{}, &model.Company{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	return db, nil
}
