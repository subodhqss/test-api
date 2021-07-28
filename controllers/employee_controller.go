package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subodhqss/test-api/config"
	"github.com/subodhqss/test-api/models"
	"github.com/subodhqss/test-api/repository"
	"github.com/subodhqss/test-api/services"
	"github.com/subodhqss/test-api/utils"
)

var userService = services.NewUserService(repository.NewEmployeeRepository())

func GetEmployees(rw http.ResponseWriter, r *http.Request) {
	data, err := userService.GetEmployees()
	if err != nil {
		fmt.Println("error in getting user data ", err)
		utils.WriteResponse(rw, utils.GetErrorResponse(err, http.StatusBadRequest, config.BadRequest), http.StatusBadRequest)
	}
	utils.WriteResponse(rw, data, 200)
}

func GetEmployeeById(rw http.ResponseWriter, r *http.Request) {
	var empNo string

	vars := mux.Vars(r)
	empNo = vars["employeeNumber"]
	log.Print("employee number in contorller ", empNo)
	data, err := userService.GetEmployeeById(empNo)
	if err != nil {
		fmt.Println("error in getting employee by id", err)
	}
	utils.WriteResponse(rw, data, 200)
}
func GetEmployeeInOffice(rw http.ResponseWriter, r *http.Request) {
	var officeCode string

	vars := mux.Vars(r)
	officeCode = vars["officeCode"]
	log.Print("office code in contorller ", officeCode)
	data, err := userService.GetEmployeeInOffice(officeCode)
	if err != nil {
		fmt.Println("error in getting office by office code", err)
	}
	utils.WriteResponse(rw, data, 200)
}

func AddEmployee(rw http.ResponseWriter, r *http.Request) {
	postBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("Error in reading the post body")
		utils.WriteResponse(rw, utils.GetErrorResponse(err, http.StatusBadRequest, config.BadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	emp := &models.Employee{}
	if err = json.Unmarshal(postBody, &emp); err != nil {
		utils.WriteResponse(rw, utils.GetErrorResponse(err, http.StatusBadRequest, config.BadRequest), http.StatusBadRequest)
		return
	}

	employees, err := userService.AddEmployee(emp)
	if err != nil {
		utils.WriteResponse(rw, utils.GetErrorResponse(err, http.StatusBadRequest, config.BadRequest), http.StatusBadRequest)
		return
	}
	utils.WriteResponse(rw, employees, http.StatusCreated)
}

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	qrCode, err := userService.GetQRCodeImage()
	if err != nil {
		fmt.Println("error in generating QR code ", err)
	}
	utils.WriteResponse(w, qrCode, 200)
}
