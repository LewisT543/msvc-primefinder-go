package aoc

import (
	"encoding/json"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type AOCHandler struct {
	Problems []*AOCProblem
}

type SolverFn func(string) (string, error)

func NewAOCHandler() (*AOCHandler, error) {
	problemData := []struct {
		Filename         string
		ShortDescription string
		Solver           SolverFn
	}{
		{Filename: "1_historian_hysteria", ShortDescription: "short-desc", Solver: func(string) (string, error) { return "5", nil }},
	}

	var problems []*AOCProblem

	for _, data := range problemData {
		problem, err := NewAOCProblem(data.Filename, data.ShortDescription, data.Solver)
		if err != nil {
			return nil, fmt.Errorf("error creating AOCProblem for %s: %v", data.Filename, err)
		}
		problems = append(problems, problem)
	}

	return &AOCHandler{
		Problems: problems,
	}, nil
}

func (h *AOCHandler) HandleAOC(w http.ResponseWriter, r *http.Request) {
	day := chi.URLParam(r, "day")
	if day == "" {
		http.Error(w, "Invalid route", http.StatusBadRequest)
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil || dayInt < 1 || dayInt > len(h.Problems) {
		http.Error(w, "Invalid day parameter", http.StatusBadRequest)
		return
	}

	problem := h.Problems[dayInt-1]

	input, err := utils.ReadFromFile(problem.Filename, ".txt", "inputs")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read input at: %v", err), http.StatusInternalServerError)
		return
	}

	start := time.Now()
	solved, err := problem.SolverFn(input)
	duration := time.Since(start)
	if err != nil {
		http.Error(w, fmt.Sprintf("error solving problem: %v", err), http.StatusInternalServerError)
		return
	}

	var response struct {
		Title            string        `json:"title"`
		Day              int           `json:"day"`
		ShortDescription string        `json:"description"`
		Solution         string        `json:"solution"`
		Duration         time.Duration `json:"duration"`
	}

	response.Title = problem.Title
	response.Day = problem.Day
	response.ShortDescription = problem.ShortDescription
	response.Solution = solved
	response.Duration = duration

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		fmt.Println("failed to write:", err)
		return
	}
}
