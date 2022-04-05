package osx

import "os"

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will default to the provided default (`def`)
// if not the variable is not found.
func Getenv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return def
}
