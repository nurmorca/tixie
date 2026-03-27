package app

import (
	"ticket/common/cache"
	"ticket/common/database"
)

type ConfigManager struct {
	PostgreSqlConfig database.Config
	RedisConfig      cache.RConfig
}

func NewConfigManager() *ConfigManager {
	postgresConfig := getPostgreSqlConfig()
	redisConfig := getRedisConfig()
	return &ConfigManager{
		PostgreSqlConfig: postgresConfig,
		RedisConfig:      redisConfig,
	}
}

func getPostgreSqlConfig() database.Config {
	return database.Config{
		Host:                  "localhost",
		Port:                  "5434",
		Username:              "admin",
		Password:              "admin123",
		DBname:                "ticketdb",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "30s",
	}
}

func getRedisConfig() cache.RConfig {
	return cache.RConfig{
		Address:  "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	}
}
