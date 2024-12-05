package utils

import (
	"github.com/LewisT543/msvc-primefinder-go/internal/model"
	"github.com/google/uuid"
	"math/rand/v2"
	"time"
)

const maxLineItems = 10
const minLineItems = 6

const maxQuantity = 10
const maxPrice = 10000

func GenerateOrders(n int) []model.Order {
	orders := make([]model.Order, n)

	for i := 0; i < n; i++ {
		numLineItems := rand.IntN(maxLineItems-minLineItems+1) + minLineItems
		createdAt := time.Now()

		lineItems := make([]model.LineItem, numLineItems)
		for j := 0; j < numLineItems; j++ {
			lineItems[j] = model.LineItem{
				ItemID:   uuid.New(),
				Quantity: uint(rand.IntN(maxQuantity) + 1),
				Price:    uint(rand.IntN(maxPrice) + 1),
			}
		}

		orders[i] = model.Order{
			CustomerID: uuid.New(),
			LineItems:  lineItems,
			CreatedAt:  &createdAt,
		}
	}

	return orders
}
