package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Romzi",
		MiddleName: "Farhan",
		LastName:   "Kaka",
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
