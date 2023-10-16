package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	input := `{
        "name": "John Doe",
        "age": 15,
        "hobbies": ["climbing", "cycling", "running"]
    }`

	var target map[string]any

	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	for k, v := range target {
		fmt.Printf("K: %s, V: %v\n", k, v)
	}
}
