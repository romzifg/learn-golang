package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Macbook",
		ImageUrl: "http://google",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Macbook","image_url":"http://google"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
