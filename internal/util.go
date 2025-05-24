package internal

import "encoding/json"

func PrettyJson(data any) string {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonData)
}

func PrettyPrintJson(data any) {
	println(PrettyJson(data))
}