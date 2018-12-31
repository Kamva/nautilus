package nautilus

import (
	"os"
	"reflect"
)

// GetEnv get the environment variable and set it to given fallback if not present
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

// Empty Check if i is empty (Zero valued)
func Empty(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
