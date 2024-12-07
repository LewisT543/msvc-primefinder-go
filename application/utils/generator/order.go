package generator

import (
	"github.com/LewisT543/msvc-primefinder-go/application/model"
	"github.com/google/uuid"
	"math/rand/v2"
	"time"
)

func GenerateOrders(n int, opt GenerateOrderOptions) []model.Order {
	orders := make([]model.Order, n)

	for i := 0; i < n; i++ {
		order := generateOrder(opt)
		orders[i] = order
	}
	return orders
}

func generateOrder(opt GenerateOrderOptions) model.Order {
	numLineItems := rand.IntN(opt.MaxLineItems-opt.MinLineItems+1) + opt.MinLineItems
	createdAt := time.Now()

	lineItems := make([]model.LineItem, numLineItems)
	for j := 0; j < numLineItems; j++ {
		lineItems[j] = model.LineItem{
			ItemID:   uuid.New(),
			Quantity: uint(rand.IntN(opt.MaxQuantity) + 1),
			Price:    uint(rand.IntN(opt.MaxPrice) + 1),
		}
	}

	return model.Order{
		OrderID:    rand.Uint64(),
		CustomerID: uuid.New(),
		LineItems:  lineItems,
		CreatedAt:  &createdAt,
	}
}
