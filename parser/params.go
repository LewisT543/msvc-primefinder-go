package parser

import (
	"fmt"
	"net/http"
	"strconv"
)

func ParseQueryParam[T any](r *http.Request, key string, defaultValue *T, parser func(string) (T, error)) (T, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		if defaultValue != nil {
			return *defaultValue, nil
		}
		var zero T
		return zero, fmt.Errorf("missing required query parameter: %s", key)
	}

	parsedValue, err := parser(value)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("invalid value for %s: %s (%v)", key, value, err)
	}
	return parsedValue, nil
}

func IntParser(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func StringParser(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("empty string")
	}
	return value, nil
}
