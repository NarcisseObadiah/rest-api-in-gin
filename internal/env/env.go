package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaulValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaulValue
}

func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultValue
}
