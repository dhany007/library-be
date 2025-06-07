package env

import "os"

func GetEnv(key string, defult string) string {
	value := os.Getenv(key)
	if value == "" {
		return defult
	}

	return value
}
