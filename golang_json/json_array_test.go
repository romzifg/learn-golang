package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Romzi",
		MiddleName: "Farhan",
		LastName:   "Ghozi",
		Age:        25,
		Hobbies:    []string{"Coding", "Gaming", "Music"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Romzi","MiddleName":"Farhan","LastName":"Ghozi","Age":25,"Married":false,"Hobbies":["Coding","Gaming","Music"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Romzi",
		Addresses: []Address{
			{
				Street:     "Jl Ada",
				Country:    "Indonesia",
				PostalCode: "666676",
			},
			{
				Street:     "Jl Belum ada",
				Country:    "Indonesia",
				PostalCode: "66232",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Romzi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl Ada","Country":"Indonesia","PostalCode":"666676"},{"Street":"Jl Belum ada","Country":"Indonesia","PostalCode":"66232"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jl Ada",
			Country:    "Indonesia",
			PostalCode: "666676",
		},
		{
			Street:     "Jl Belum ada",
			Country:    "Indonesia",
			PostalCode: "66232",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl Ada","Country":"Indonesia","PostalCode":"666676"},{"Street":"Jl Belum ada","Country":"Indonesia","PostalCode":"66232"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
