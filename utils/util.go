package utils

import (
	"encoding/json"
	"net/http"
)

// Message status in JSON format
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"status": status, "message": message,
	}
}

// Respond in JSON format
func Respond(write http.ResponseWriter, data map[string]interface{}) {
	write.Header().Add("Content-Type", "application/json")
	json.NewEncoder(write).Encode(data)
}
