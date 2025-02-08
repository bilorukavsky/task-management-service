package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	//Env      string         `json:"env"` // Можно добавить поле для окружения (local, development, stage, production)
	AuthJWT  AuthJWTConfig  `json:"auth_jwt"`
	Database DatabaseConfig `json:"database"`
	Http     HttpConfig     `json:"http"`
	Redis    RedisConfig    `json:"redis"`
}

type HttpConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"-"`
	Password string `json:"-"`
	DBName   string `json:"name"`
	SSLMode  string `json:"sslmode"`
}

type RedisConfig struct {
	Addr     string        `json:"addr"`
	Password string        `json:"-"`
	DB       int           `json:"db"`
	Timeout  time.Duration `json:"timeout"`
}

type AuthJWTConfig struct {
	PrivateKey                string `json:"-"`
	PublicKey                 string `json:"-"`
	Algorithm                 string `json:"algorithm"`
	AccessTokenExpireMinutes  int    `json:"access_token_expire_minutes"`
	RefreshTokenExpireMinutes int    `json:"refresh_token_expire_minutes"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("config")

	viper.AutomaticEnv()

	_ = viper.BindEnv("database.password", "DB_PASSWORD")
	_ = viper.BindEnv("database.user", "DB_USER")
	_ = viper.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = viper.BindEnv("auth_jwt.private_key", "JWT_PRIVATE")
	_ = viper.BindEnv("auth_jwt.public_key", "JWT_PUBLIC")
	_ = viper.BindEnv("auth_jwt.algorithm", "JWT_ALGORITHM")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Обновляем критичные поля из переменных окружения
	config.Database.Password = viper.GetString("database.password")
	config.Database.User = viper.GetString("database.user")
	config.Redis.Password = viper.GetString("redis.password")
	config.AuthJWT.PrivateKey = viper.GetString("auth_jwt.private_key")
	config.AuthJWT.PublicKey = viper.GetString("auth_jwt.public_key")
	config.AuthJWT.Algorithm = viper.GetString("auth_jwt.algorithm")

	// Проверяем критичные переменные
	if config.Database.Password == "" {
		return nil, errors.New("Ошибка: DB_PASSWORD не установлен")
	}
	if config.Database.User == "" {
		return nil, errors.New("Ошибка: DB_USER не установлен")
	}
	if config.Redis.Password == "" {
		return nil, errors.New("Ошибка: REDIS_PASSWORD не установлен")
	}
	if config.AuthJWT.PrivateKey == "" {
		return nil, errors.New("Ошибка: JWT_PRIVATE не установлен")
	}
	if config.AuthJWT.PublicKey == "" {
		return nil, errors.New("Ошибка: JWT_PUBLIC не установлен")
	}
	if config.AuthJWT.Algorithm == "" {
		return nil, errors.New("Ошибка: JWT_ALGORITHM не установлен")
	}

	return &config, nil
}
