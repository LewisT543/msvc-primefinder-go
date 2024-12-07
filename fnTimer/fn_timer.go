package fnTimer

import "time"

func Timer[Args any, R any](fn func(args ...Args) R) func(args ...Args) (R, time.Duration) {
	return func(args ...Args) (R, time.Duration) {
		start := time.Now()
		result := fn(args...)
		duration := time.Since(start)
		return result, duration
	}
}
