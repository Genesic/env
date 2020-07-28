package env

import (
	"os"
	"strconv"
)

func Get(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err == nil {
			return v
		}
	}

	return fallback
}
