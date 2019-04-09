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
func Empty(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		data, _ := json.Marshal(i)
		return string(data) == "[]" || string(data) == "{}"
	default:
		return reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface())
	}
}
