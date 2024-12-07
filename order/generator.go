package order

import (
	"github.com/google/uuid"
	"math/rand/v2"
	"time"
)

type GenerateOrderOptions struct {
	MaxLineItems int
	MinLineItems int
	MaxQuantity  int
	MaxPrice     int
}

func NewGenerateOrderOptions(opts ...func(*GenerateOrderOptions)) GenerateOrderOptions {
	// Default values
	options := GenerateOrderOptions{
		MaxLineItems: 10,
		MinLineItems: 6,
		MaxQuantity:  10,
		MaxPrice:     10000,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return options
}

func WithMaxLineItems(max int) func(*GenerateOrderOptions) {
	return func(o *GenerateOrderOptions) {
		o.MaxLineItems = max
	}
}

func WithMinLineItems(min int) func(*GenerateOrderOptions) {
	return func(o *GenerateOrderOptions) {
		o.MinLineItems = min
	}
}

func WithMaxQuantity(max int) func(*GenerateOrderOptions) {
	return func(o *GenerateOrderOptions) {
		o.MaxQuantity = max
	}
}

func WithMaxPrice(max int) func(*GenerateOrderOptions) {
	return func(o *GenerateOrderOptions) {
		o.MaxPrice = max
	}
}

func GenerateOrders(n int, opt GenerateOrderOptions) []Order {
	orders := make([]Order, n)

	for i := 0; i < n; i++ {
		ord := generateOrder(opt)
		orders[i] = ord
	}
	return orders
}

func generateOrder(opt GenerateOrderOptions) Order {
	numLineItems := rand.IntN(opt.MaxLineItems-opt.MinLineItems+1) + opt.MinLineItems
	createdAt := time.Now()

	lineItems := make([]LineItem, numLineItems)
	for j := 0; j < numLineItems; j++ {
		lineItems[j] = LineItem{
			ItemID:   uuid.New(),
			Quantity: uint(rand.IntN(opt.MaxQuantity) + 1),
			Price:    uint(rand.IntN(opt.MaxPrice) + 1),
		}
	}

	return Order{
		OrderID:    rand.Uint64(),
		CustomerID: uuid.New(),
		LineItems:  lineItems,
		CreatedAt:  &createdAt,
	}
}
