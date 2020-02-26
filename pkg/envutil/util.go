package envutil

import (
	"os"
	"strconv"
)

func Get(key string, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func GetInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return i
}

func GetInt64(key string, fallback int64) int64 {
	s := os.Getenv(key)
	if s == "" {
		return fallback
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fallback
	}
	return v
}
