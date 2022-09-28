package data

import "testing"

func TestChecksValidation(t *testing.T) {
    p := &Product{
        Name: "hi",
        Price: 1.00,
        SKU: "ads-dsa-af",
    }

    err := p.Validate()

    if err != nil {
        t.Fatal(err)
    }
}
