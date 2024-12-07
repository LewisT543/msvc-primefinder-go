package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/application/model"
	"github.com/LewisT543/msvc-primefinder-go/application/repository/order"
	"github.com/LewisT543/msvc-primefinder-go/application/utils/generator"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID        `json:"customer_id"`
		Lineitems  []model.LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	newOrder := model.Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.Lineitems,
		CreatedAt:  &now,
	}

	err := o.Repo.Insert(r.Context(), newOrder)
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

func (o *Order) Generate(w http.ResponseWriter, r *http.Request) {
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

	options := generator.NewGenerateOrderOptions()
	orders := generator.GenerateOrders(int(quant), options)

	_, err = json.Marshal(orders)
	if err != nil {
		fmt.Println("failed to marshal:", err)
	}

	fmt.Println("Inserting orders")
	fmt.Println(orders)

	for _, ord := range orders {
		err := o.Repo.Insert(r.Context(), ord)
		if err != nil {
			fmt.Println("failed to write: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
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
	res, err := o.Repo.FindAll(r.Context(), order.FindAllPage{
		Offset: cursor,
		Size:   size,
	})
	if err != nil {
		fmt.Println("failed to find all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []model.Order `json:"items"`
		Next  uint64        `json:"next,omitempty"` // omit the field in case there is an empty value
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

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ord, err := o.Repo.FindByID(r.Context(), orderID)
	if errors.Is(err, order.ErrNotExist) {
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

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
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

	retrievedOrder, err := o.Repo.FindByID(r.Context(), orderID)
	if errors.Is(order.ErrNotExist, err) {
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

	err = o.Repo.Update(r.Context(), retrievedOrder)
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

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = o.Repo.DeleteByID(r.Context(), orderID)
	if errors.Is(order.ErrNotExist, err) {
		fmt.Println("order not found:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
