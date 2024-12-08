package setup

import (
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/aoc"
	"github.com/LewisT543/msvc-primefinder-go/order"
	"github.com/LewisT543/msvc-primefinder-go/primes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", a.loadOrderRoutes)

	router.Route("/find-primes", a.loadPrimeRoutes)

	router.Route("/aoc", a.loadAOCRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &order.OrderHandler{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Post("/", orderHandler.Create)
	router.Post("/generate", orderHandler.Generate)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}

func (a *App) loadPrimeRoutes(router chi.Router) {
	primeHandler := &primes.PrimeHandler{
		Repo: &primes.RedisRepo{
			Client: a.rdb,
		},
		Algo: primes.SegmentedSieveCalculator{},
	}

	router.Get("/", primeHandler.FindPrimes)
}

func (a *App) loadAOCRoutes(router chi.Router) {
	aocHandler, err := aoc.NewAOCHandler()
	if err != nil {
		fmt.Printf("failed to construct new AOCHandler: %v\n", err)
	}

	fmt.Println("AOC Problems Loaded:")
	for _, problem := range aocHandler.Problems {
		fmt.Printf("\t%s\n", problem.Filename)
	}

	router.Get("/{day}", aocHandler.HandleAOC)
}
