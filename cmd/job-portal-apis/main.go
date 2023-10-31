package main

import (
	"fmt"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/database"
	"job-portal-api/internal/handlers"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/services"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
}
func startApp() error {
	log.Info().Msg("started main")
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 4\Desktop\coding\job-portal-api\private.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file private.pem %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 4\Desktop\coding\job-portal-api\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file pubkey.pem %w", err)
	}
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
	repo, err := repository.NewRepo(db)
	if err != nil {
		return err
	}

	se, err := services.NewService(repo, repo)

	if err != nil {
		return err
	}

	api := http.Server{
		Addr:    ":8085",
		Handler: handlers.Api(a, *se),
	}
	err = api.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("error in creating the server")
	}

	return nil

}
