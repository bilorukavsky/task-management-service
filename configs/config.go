package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	//Env      string         `json:"env"` // Можно добавить поле для окружения (local, development, stage, production)
	AuthJWT  AuthJWTConfig  `json:"auth_jwt"`
	Database DatabaseConfig `json:"database"`
	Http     HttpConfig     `json:"http"`
	Redis    RedisConfig    `json:"redis"`
	Logger   LoggerConfig   `json:"logger"`
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
	PrivateKeyPath            string `json:"-"`
	PublicKeyPath             string `json:"-"`
	Algorithm                 string `json:"algorithm"`
	AccessTokenExpireMinutes  int    `json:"access_token_expire_minutes"`
	RefreshTokenExpireMinutes int    `json:"refresh_token_expire_minutes"`
}

type LoggerConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	_ = viper.BindEnv("database.password", "DB_PASSWORD")
	_ = viper.BindEnv("database.user", "DB_USER")
	_ = viper.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = viper.BindEnv("auth_jwt.private_key_path", "JWT_PRIVATE_KEY_PATH")
	_ = viper.BindEnv("auth_jwt.public_key_path", "JWT_PUBLIC_KEY_PATH")
	_ = viper.BindEnv("auth_jwt.algorithm", "JWT_ALGORITHM")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Не удалось загрузить config.yaml: %v. Используем только переменные окружения.", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Ошибка при парсинге конфигурации: %v", err)
		return nil, err
	}

	// Обновляем критичные поля из переменных окружения
	config.Database.Password = viper.GetString("database.password")
	config.Database.User = viper.GetString("database.user")
	config.Redis.Password = viper.GetString("redis.password")
	config.AuthJWT.PrivateKeyPath = viper.GetString("auth_jwt.private_key_path")
	config.AuthJWT.PublicKeyPath = viper.GetString("auth_jwt.public_key_path")
	config.AuthJWT.Algorithm = viper.GetString("auth_jwt.algorithm")

	// Проверяем критичные переменные
	if config.Database.Password == "" {
		log.Fatal("Ошибка: DB_PASSWORD не установлен")
	}
	if config.Database.User == "" {
		log.Fatal("Ошибка: DB_USER не установлен")
	}
	if config.Redis.Password == "" {
		log.Fatal("Ошибка: REDIS_PASSWORD не установлен")
	}
	if config.AuthJWT.PrivateKeyPath == "" {
		log.Fatal("Ошибка: JWT_PRIVATE_KEY_PATH не установлен")
	}
	if config.AuthJWT.PublicKeyPath == "" {
		log.Fatal("Ошибка: JWT_PUBLIC_KEY_PATH не установлен")
	}
	if config.AuthJWT.Algorithm == "" {
		log.Fatal("Ошибка: JWT_ALGORITHM не установлен")
	}

	return &config, nil
}
