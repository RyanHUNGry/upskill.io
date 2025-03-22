package utils

import (
	"log"
	"os"
)

// Generic map function
func FunctionMap[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// Convert typed slice to an empty interface slice
func AnySliceConverter[S any](slice []S) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func GetWorkingDirectory() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to get working directory")
	}

	return dir
}
