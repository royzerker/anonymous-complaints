package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type IConfigService interface {
	Get(key string) string
	GetAndCheck(key string) string
	GetNumber(key string) int
	GetNumberAndCheck(key string) int
	GetBool(key string) bool
	GetBoolAndCheck(key string) bool
}

type configService struct {
	once sync.Once
}

func ConfigService() IConfigService {
	c := &configService{}
	c.loadEnv()
	return c
}

func (c *configService) loadEnv() {
	c.once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found")
		}
	})
}

func (c *configService) Get(key string) string {
	return os.Getenv(key)
}

func (c *configService) GetAndCheck(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return val
}

func (c *configService) GetNumber(key string) int {
	valStr := os.Getenv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0
	}
	return val
}

func (c *configService) GetNumberAndCheck(key string) int {
	valStr := os.Getenv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatalf("Invalid number for environment variable %s: %v", key, err)
	}
	return val
}

func (c *configService) GetBool(key string) bool {
	val := os.Getenv(key)
	return val == "true" || val == "1"
}

func (c *configService) GetBoolAndCheck(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required boolean environment variable: %s", key)
	}
	return val == "true" || val == "1"
}

/**
cfg := config.NewConfigService()

port := cfg.GetAndCheck("PORT")
debugMode := cfg.GetBool("DEBUG_MODE")
timeout := cfg.GetNumber("REQUEST_TIMEOUT")
*/
