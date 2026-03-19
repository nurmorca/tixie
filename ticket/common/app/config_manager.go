package app

import "ticket/common/database"

type ConfigManager struct {
	PostgreSqlConfig database.Config
}

func NewConfigManager() *ConfigManager {
	postgresConfig := getPostgreSqlConfig()
	return &ConfigManager{
		PostgreSqlConfig: postgresConfig,
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
