package controllers

import (
	"net/http"

	"github.com/subodhqss/test-api/models"
	"github.com/subodhqss/test-api/utils"
)

func TestController(w http.ResponseWriter, r *http.Request) {
	data := models.Employee{}
	utils.WriteResponse(w, data, 200)
}
