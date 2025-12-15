package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	RedisAddr    string `mapstructure:"REDIS_ADDR"`
	RedisPass    string `mapstructure:"REDIS_PASS"`
	KafkaBrokers string `mapstructure:"KAFKA_BROKERS"`
}

func LoadConfig() *Config {
	// --- DEBUG SECTION START ---
	// Let's check what the container ACTUALLY sees before Viper touches it.
	log.Println("🔍 --- DEBUG: RAW ENVIRONMENT VARIABLES ---")
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	log.Println("REDIS_ADDR:", os.Getenv("REDIS_ADDR"))
	log.Println("🔍 --- END DEBUG ---")
	// --- DEBUG SECTION END ---

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// 1. BIND ENV VARS EXPLICITLY
	// This ensures Viper looks for these Env Vars even if the .env file is missing
	_ = viper.BindEnv("APP_ENV")
	_ = viper.BindEnv("SERVER_PORT")
	_ = viper.BindEnv("DB_HOST")
	_ = viper.BindEnv("DB_PORT")
	_ = viper.BindEnv("DB_USER")
	_ = viper.BindEnv("DB_PASSWORD")
	_ = viper.BindEnv("DB_NAME")
	_ = viper.BindEnv("REDIS_ADDR")
	_ = viper.BindEnv("REDIS_PASS")
	_ = viper.BindEnv("KAFKA_BROKERS")

	// Try reading file, but don't fail if missing
	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  No .env file found (Expected inside Docker)")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Unable to decode config: %v", err)
	}

	// FALLBACK: If Viper failed to map them, fill them manually
	// This fixes the issue if 'mapstructure' tags are ignored for some reason
	if cfg.DBHost == "" {
		log.Println("⚠️ Viper returned empty config. Attempting manual fallback...")
		cfg.DBHost = os.Getenv("DB_HOST")
		cfg.DBPort = os.Getenv("DB_PORT")
		cfg.DBUser = os.Getenv("DB_USER")
		cfg.DBPassword = os.Getenv("DB_PASSWORD")
		cfg.DBName = os.Getenv("DB_NAME")
		cfg.RedisAddr = os.Getenv("REDIS_ADDR")
		cfg.RedisPass = os.Getenv("REDIS_PASS")
		cfg.KafkaBrokers = os.Getenv("KAFKA_BROKERS")
	}

	return &cfg
}