package main

import (
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
		Repo: primes.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Get("/", primeHandler.FindPrimes)
}
