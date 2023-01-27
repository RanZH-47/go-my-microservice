package coffee

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdateOn    string
	DeletedOn   string
}

var products = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "xyz456",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Americano",
		Description: "Slightly diluted coffee",
		Price:       1.35,
		SKU:         "def123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
