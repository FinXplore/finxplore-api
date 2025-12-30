package config

import (
	"log"
	"reflect"
	"strings"

	// "github.com/joho/godotenv"
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

	JWTSecret string `mapstructure:"jwt_secret"`

	DataWorkerURL string `mapstructure:"data_worker_url"`
	DataWorkerApiKey string `mapstructure:"data_worker_api_key"`

	
}

func BindEnvs(v *viper.Viper, cfg interface{}) {
	val := reflect.ValueOf(cfg)
	typ := reflect.TypeOf(cfg)

	// handle pointer
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		tag := field.Tag.Get("mapstructure")
		if tag == "" || tag == "-" {
			continue
		}

		// mapstructure: "db_host" -> DB_HOST
		envKey := strings.ToUpper(tag)

		// Bind using mapstructure key
		_ = v.BindEnv(tag, envKey)
	}
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	// -------------------------------
	// Read from .env (local dev)
	// -------------------------------
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

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
	// v.SetDefault("db_user", "postgres")
	// v.SetDefault("db_name", "finxplore")
	// v.SetDefault("db_port", 5432)
	BindEnvs(v, &Config{})

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
