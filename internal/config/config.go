package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string `mapstructure:"app_env"`
	ServerPort int    `mapstructure:"server_port"`

	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`

	RedisAddr    string `mapstructure:"redis_addr"`
	RedisPass    string `mapstructure:"redis_pass"`
	KafkaBrokers string `mapstructure:"kafka_brokers"`
}

func LoadConfig() (*Config, error) {
	// _ = godotenv.Load()

	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println("❌ godotenv error:", err)
	// } else {
	// 	fmt.Println("✅ .env loaded")
	// }
	v := viper.New()

	// -------------------------------
	// Read from .env (local dev)
	// -------------------------------
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	// -------------------------------
	// Read from ENV (Docker / CI)
	// -------------------------------
	v.AutomaticEnv()

	// 🔑 KEY LINE:
	// server_port  <-> SERVER_PORT
	// db_port      <-> DB_PORT
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// -------------------------------
	// Defaults (safe fallbacks)
	// -------------------------------
	// v.SetDefault("app_env", "local")
	// v.SetDefault("server_port", 8080)
	// v.SetDefault("db_port", 5432)

	// -------------------------------
	// Read .env if present
	// -------------------------------
	_ = v.ReadInConfig() // ignore error (expected in Docker)

	// -------------------------------
	// Unmarshal into struct
	// -------------------------------
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
