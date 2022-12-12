package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoderStream(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)
	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
