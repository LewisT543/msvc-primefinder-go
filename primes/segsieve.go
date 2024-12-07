package primes

import "math"

func SegmentedSieve(low, high int) []int {
	if low > high || low < 2 {
		return []int{}
	}

	limit := int(math.Sqrt(float64(high)))
	smallPrimes := generatePrimesUpTo(limit)

	segment := createSegment(low, high)
	for _, p := range smallPrimes {
		markMultiples(p, low, high, segment)
	}

	return extractPrimes(segment, low)
}

func generatePrimesUpTo(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	return collectPrimes(isPrime)
}

func collectPrimes(isPrime []bool) []int {
	primes := []int{}
	for i, prime := range isPrime {
		if i >= 2 && prime {
			primes = append(primes, i)
		}
	}
	return primes
}

func createSegment(low, high int) []bool {
	size := high - low + 1
	isPrime := make([]bool, size)
	for i := range isPrime {
		isPrime[i] = true
	}
	return isPrime
}

func markMultiples(p, low, high int, isPrime []bool) {
	start := max(p*p, (low+p-1)/p*p) // Ensure marking starts within the segment
	for multiple := start; multiple <= high; multiple += p {
		isPrime[multiple-low] = false
	}
}

// extractPrimes extracts the prime numbers from a segment
func extractPrimes(isPrime []bool, low int) []int {
	primes := []int{}
	for i, prime := range isPrime {
		if prime {
			primes = append(primes, low+i)
		}
	}
	return primes
}
