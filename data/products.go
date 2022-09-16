package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    SKU         string  `json:"sku"`
    CreatedOn   string  `json:"-"`
    UpdatedOn   string  `json:"-"`
    DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p*Products) ToJSON(w io.Writer) error {
    e := json.NewEncoder(w)
    return e.Encode(p)
}

func GetProducts() Products {
    return productList
}

var productList = []*Product{
    {
        ID:             1,
        Name:           "Latte",
        Description:    "Coffee wit milk",
        Price:          3.99,
        SKU:            "ssl125",
        CreatedOn:      time.Now().UTC().String(),
        UpdatedOn:      time.Now().UTC().String(),
    },
    {
        ID:             2,
        Name:           "Espresso",
        Description:    "Short and strong coffee",
        Price:          2.99,
        SKU:            "lka799",
        CreatedOn:      time.Now().UTC().String(),
        UpdatedOn:      time.Now().UTC().String(),
    },
}

