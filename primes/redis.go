package primes

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func (r RedisRepo) Insert(ctx context.Context, primes []int) error {
	fmt.Println("INSERT IN PRIMES-REDIS")
	return nil
}
