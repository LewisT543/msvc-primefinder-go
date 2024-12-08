package primes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/parser"
	"net/http"
	"time"
)

type FindPrimesResult struct {
	Result         []int         `json:"result"`
	NumberOfPrimes int           `json:"numberOfPrimes"`
	AlgorithmName  string        `json:"algorithmName"`
	Duration       time.Duration `json:"duration"`
}

type PrimeCalculator interface {
	Calculate(low int, high int) []int
}

type PrimesRepo interface {
	Insert(ctx context.Context, primes []int) error
}

type PrimeHandler struct {
	Repo PrimesRepo
	Algo PrimeCalculator
}

const lowHighErrorMessage = "'low' must be >= 2 and 'high' must be geater than 'low'"

func (h PrimeHandler) FindPrimes(w http.ResponseWriter, r *http.Request) {
	highLimit, err := parser.ParseQueryParam(r, "high", nil, parser.IntParser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lowLimit, err := parser.ParseQueryParam(r, "low", nil, parser.IntParser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if lowLimit < 2 || highLimit <= lowLimit {
		http.Error(w, lowHighErrorMessage, http.StatusBadRequest)
		return
	}

	start := time.Now()
	primes := h.Algo.Calculate(int(lowLimit), int(highLimit))
	duration := time.Since(start)

	result := FindPrimesResult{
		Result:         primes,
		NumberOfPrimes: len(primes),
		AlgorithmName:  "Segmented-Sieve",
		Duration:       duration,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
