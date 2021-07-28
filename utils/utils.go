package utils

import (
	"encoding/json"
	"net/http"

	"github.com/subodhqss/test-api/models"
)

func WriteResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, _ := json.Marshal(data)
	WriteJsonData(w, jsonData, statusCode)
}

func WriteJsonData(w http.ResponseWriter, data []byte, statusHeader int) {
	//Write out the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusHeader)
	w.Write(data)
}

func GetErrorResponse(err error, statusCode int, errType string) *models.ErrorResponse {
	errRes := &models.ErrorResponse{Message: err.Error(), Type: errType, StatusCode: statusCode}
	return errRes
}
