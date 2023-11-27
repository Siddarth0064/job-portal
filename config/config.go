package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig
	DatabaseConfing
	RedisConfig
}
type AppConfig struct {
	Port    string `env:"APP_PORT,required=true"`
	Private string `env:"PRIVATE_KEY,required=true"`
	Public  string `env:"PUBLIC_KEY,required=true"`
}
type DatabaseConfing struct {
	DBConnection string `env:"DB_DSN,required=true"`
}

type RedisConfig struct {
	Address       string `env:"ADDRESS,required=true"`
	RedisPassword string `env:"PASSWORD,required=true"`
	Db            int    `env:"DB_NUMBER,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
func GetConfig() Config {
	return cfg
}
