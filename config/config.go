package config

import (
	"os"
	"strconv"
	"sync"
)

// Singleton
var c *config
var once sync.Once

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	env := getEnv(key, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return result
}

func getConnConfig() connConfig {
	return connConfig{
		ConnHost: getEnv("CONN_HOST", "localhost"),
		ConnPort: getEnvAsInt("CONN_PORT", 42069),
		ConnType: getEnv("CONN_TYPE", "tcp"),
	}
}

// New - Initialize Configuration
func New() *config {
	once.Do(func() {
		c = &config{
			ConnConfig: getConnConfig(),
		}
	})
	return c
}
