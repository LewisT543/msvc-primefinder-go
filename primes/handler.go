package primes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/fnTimer"
	"net/http"
	"strconv"
	"time"
)

type FindPrimesResult struct {
	Result         []int
	NumberOfPrimes int
	AlgorithmName  string
	Duration       time.Duration
}

type PrimesRepo interface {
	Insert(ctx context.Context, primes []int) error
}

type PrimeHandler struct {
	Repo RedisRepo
}

func (h PrimeHandler) FindPrimes(w http.ResponseWriter, r *http.Request) {
	const base = 10
	const bitSize = 64
	var lowLimit int64 = 0
	var highLimit int64 = 0

	high := r.URL.Query().Get("high")
	if high == "" {
		fmt.Println("no high query parameter")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	highLimit, err := strconv.ParseInt(high, base, bitSize)
	if err != nil {
		fmt.Println("failed to parse high limit: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	low := r.URL.Query().Get("low")
	lowLimit, err = strconv.ParseInt(low, base, bitSize)
	if err != nil {
		fmt.Println("failed to parse low limit: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Calculating primes low: %d, high: %d\n", lowLimit, highLimit)
	timedCalculate := fnTimer.Timer(func(args ...int) []int {
		return SegmentedSieve(args[0], args[1])
	})

	primes, duration := timedCalculate(int(lowLimit), int(highLimit))

	result := FindPrimesResult{
		Result:         primes,
		NumberOfPrimes: len(primes),
		AlgorithmName:  "Segmented-Sieve",
		Duration:       duration,
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
