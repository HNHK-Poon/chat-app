package config

import (
	"os"
)

var (
	RedisAddr = getEnv("REDIS_ADDR", "redis")
	RedisPort = getEnv("REDIS_PORT", ":6379")
	RedisPass = getEnv("REDIS_PASS", "password")
	WsAddr    = getEnv("WS_ADDR", "server")
	WsPort    = getEnv("WS_PORT", ":8080")
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
