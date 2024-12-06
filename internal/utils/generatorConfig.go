package utils

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
