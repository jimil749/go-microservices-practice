package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Jim",
		Price: 1.00,
		SKU:   "abc-abc-acb",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
