package models

type ErrorResponse struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
