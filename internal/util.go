package internal

import (
	"encoding/json"
	"fmt"
)

func PrettyJson(data any) string {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonData)
}

func PrettyPrintJson(data any) {
	fmt.Println(PrettyJson(data))
}
