package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	redisHost string
	redisPort string
}

func (c *Config) GetRedisHost() string {
	return c.redisHost
}

func (c *Config) GetRedisPort() string {
	return c.redisPort
}

var config *Config = nil

func GetConfig() *Config {
	if config != nil {
		return config
	}

	initConfig()
	return config
}

func initConfig() {
	loadEnvFile(".env")
	config = &Config{
		redisHost: getEnv("REDIS_HOST", "redis"),
		redisPort: getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func loadEnvFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comments
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("cannot parse line '%s'", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
	return scanner.Err()
}
