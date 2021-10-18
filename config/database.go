package config

import (
	"workwx/pkg/config"
)

func init() {

	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{

			// 数据库连接信息
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", "wxdb"),
			"username": config.Env("DB_USERNAME", ""),
			"password": config.Env("DB_PASSWORD", ""),
			"prefix":   config.Env("DB_PREFIX", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 60*60),
		},
		"redis": map[string]interface{}{

			// 数据库连接信息
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"database": config.Env("REDIS_DATABASE", "0"),
			"password": config.Env("REDIS_PASSWORD", ""),


			// 连接池配置
			"max_idle_connections": config.Env("REDIS_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("REDIS_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Env("REDIS_MAX_LIFE_SECONDS", 60*60),
		},
	})

}
