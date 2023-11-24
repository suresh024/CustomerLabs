package utils

import (
	"bytes"
	"encoding/json"
	"github.com/suresh024/CustomerLabs/models"
	"log"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, statusCode int, status interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	en := json.NewEncoder(w)
	_ = en.Encode(status)
}

func ErrorResponse(w http.ResponseWriter, responseErrorMessage string, statusCode int, logError error) {
	w.Header().Set("Content-Type", "application/json")
	var buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	if logError != nil {
		log.Printf("error: %s", logError)
	}
	_ = encoder.Encode(models.ErrResponse{Message: responseErrorMessage})
	w.WriteHeader(statusCode)
	_, _ = w.Write(buf.Bytes())
}
