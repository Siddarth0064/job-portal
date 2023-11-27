package main

import (
	"fmt"
	"job-portal-api/config"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/database"
	"job-portal-api/internal/database/redis"
	"job-portal-api/internal/handlers"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/services"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	//"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// ====================== main func ======================================================
func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
}

// ================== starting connection to the database and APi and reading private.pem and public.pem ========
func startApp() error {
	cfg := config.GetConfig()
	log.Info().Interface("cfg", cfg).Msg("config")

	//==================================================================

	log.Info().Msg("started main")
	//privateKey:=os.Getenv("PRIVATEKEY")
	// privatePEM, err := os.ReadFile(`.private.env`)
	// if err != nil {
	// 	return fmt.Errorf("cannot find file private.pem %w", err)
	// }
	privatePEM := []byte(cfg.AppConfig.Private)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	// publicPEM, err := os.ReadFile(`.pubkey.env`)
	// if err != nil {
	// 	return fmt.Errorf("cannot find file pubkey.pem %w", err)
	// }
	publicPEM := []byte(cfg.AppConfig.Public)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("cannot create auth instance %w", err)
	}

	db, err := database.Connection()
	if err != nil {
		return err
	}

	//======================== redis connection ====================
	rdb := redis.Connection()
	redisLayer := cache.NewRDBLayer(rdb)
	//=====================================================

	repo, err := repository.NewRepo(db)
	if err != nil {
		return err
	}

	se, err := services.NewService(repo, repo, redisLayer)

	if err != nil {
		return err
	}

	//========================server port 8085======================================
	api := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppConfig.Port),
		Handler: handlers.Api(a, se),
	}
	err = api.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("error in creating the server")
	}

	return nil

}
