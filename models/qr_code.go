package models

type QRCode struct {
	ID     int    `json:"id"`
	Image  string `json:"image_base64"`
	Secret string `json:"secret"`
}
