package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Set default environments
var Env = struct {
	// Environment settings
	Env          string        `mapstructure:"ENV"`                              // ระบุ environment ที่ใช้งาน เช่น development, staging, production
	Port         string        `mapstructure:"PORT"`                             // พอร์ตที่แอปจะรันอยู่
	Cors         string        `mapstructure:"CORS"`                             // รายการ origin ที่อนุญาต (CORS)
	AppHost      string        `mapstructure:"APP_HOST" validate:"required,uri"` // Host ของแอปพลิเคชัน
	DBURI        string        `mapstructure:"DB_URI" validate:"required"`
	DBName       string        `mapstructure:"DB_NAME" validate:"required"`
	SignatureKey string        `mapstructure:"SIGNATURE_KEY"`
	SignatureExp time.Duration `mapstructure:"SIGNATURE_EXP"`
}{
	Env:     "production",
	Port:    "3000",
	Cors:    "*",
	AppHost: "http://localhost:3000",
}

func NewAppInitEnvironment() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(".env file not found, loading from environment variables only.")
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatalf("Unable to unmarshal environment variables: %s", err)
	}

	log.Println("Environment variables loaded successfully.")
}
