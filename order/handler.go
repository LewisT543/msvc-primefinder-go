package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type OrderHandler struct {
	Repo *RedisRepo
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID  `json:"customer_id"`
		Lineitems  []LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	newOrder := Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.Lineitems,
		CreatedAt:  &now,
	}

	err := h.Repo.Insert(r.Context(), newOrder)
	if err != nil {
		fmt.Println("failed to insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newOrder)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(res); err != nil {
		fmt.Println("failed to write: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) Generate(w http.ResponseWriter, r *http.Request) {
	quantityStr := r.URL.Query().Get("quantityStr")
	if quantityStr == "" {
		quantityStr = "10"
	}

	const base = 10
	const bitSize = 8
	quant, err := strconv.ParseInt(quantityStr, base, bitSize)
	if err != nil {
		fmt.Println("failed to parse:", err)
	}

	options := NewGenerateOrderOptions()
	orders := GenerateOrders(int(quant), options)

	_, err = json.Marshal(orders)
	if err != nil {
		fmt.Println("failed to marshal:", err)
	}

	fmt.Println("Inserting orders")
	fmt.Println(orders)

	for _, ord := range orders {
		err := h.Repo.Insert(r.Context(), ord)
		if err != nil {
			fmt.Println("failed to write: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		fmt.Println("failed to write:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := h.Repo.FindAll(r.Context(), FindAllPage{
		Offset: cursor,
		Size:   size,
	})
	if err != nil {
		fmt.Println("failed to find all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []Order `json:"items"`
		Next  uint64  `json:"next,omitempty"` // omit the field in case there is an empty value
	}
	response.Items = res.Orders
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(data); err != nil {
		fmt.Println("failed to write:", err)
		return
	}
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ord, err := h.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ord); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	retrievedOrder, err := h.Repo.FindByID(r.Context(), orderID)
	if errors.Is(ErrNotExist, err) {
		fmt.Println("order not found:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	const completedStatus = "completed"
	const shippedStatus = "shipped"
	now := time.Now().UTC()

	switch body.Status {
	case shippedStatus:
		if retrievedOrder.ShippedAt != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		retrievedOrder.ShippedAt = &now
	case completedStatus:
		if retrievedOrder.CompletedAt != nil || retrievedOrder.ShippedAt == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		retrievedOrder.CompletedAt = &now
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Repo.Update(r.Context(), retrievedOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(retrievedOrder); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Repo.DeleteByID(r.Context(), orderID)
	if errors.Is(ErrNotExist, err) {
		fmt.Println("order not found:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
