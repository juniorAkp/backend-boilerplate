package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(data))
}
