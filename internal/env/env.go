package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if ok {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return valInt
		}
	}
	return fallback
}
