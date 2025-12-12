package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	PostgreSQL PostgreSQL
	JWT        JWT
	App        Fiber
	Supabase   Supabase
	Mail       Mail
	Chat       Chat
	CORS       CORS
}

type Fiber struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

type JWT struct {
	Secret string
}

type Mail struct {
	Host   string
	Port   string
	Sender string
	Key    string
}

type Supabase struct {
	URL    string
	Key    string
	Bucket string
}

type Chat struct {
	URL string
}

type CORS struct {
	URL string
}

func LoadConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	requireEnv := func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			log.Fatalf("CRITICAL ERROR: Missing required environment variable '%s'", key)
		}
		return value
	}

	return &Configs{
		PostgreSQL: PostgreSQL{
			Host:     requireEnv("DB_HOST"),
			Port:     requireEnv("DB_PORT"),
			Username: requireEnv("DB_USER"),
			Password: requireEnv("DB_PASSWORD"),
			Database: requireEnv("DB_NAME"),
			SSLMode:  requireEnv("SSL_MODE"),
		},
		App: Fiber{
			Host: requireEnv("APP_HOST"),
			Port: requireEnv("APP_PORT"),
		},
		JWT: JWT{
			Secret: requireEnv("JWT_SECRET"),
		},
		Supabase: Supabase{
			URL:    requireEnv("SUPABASE_URL"),
			Key:    requireEnv("SUPABASE_KEY"),
			Bucket: requireEnv("BUCKET_NAME"),
		},
		Mail: Mail{
			Host:   requireEnv("EMAIL_HOST"),
			Port:   requireEnv("EMAIL_PORT"),
			Sender: requireEnv("EMAIL_USER"),
			Key:    requireEnv("EMAIL_PASS"),
		},
		Chat: Chat{
			URL: requireEnv("CHAT_API_URL"),
		},
		CORS: CORS{
			URL: requireEnv("CORS"),
		},
	}
}
