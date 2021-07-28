package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", TestController)

	// employee subroute
	employeeSubRoute := mux.NewRouter().PathPrefix("/employees").Subrouter().StrictSlash(true)
	employeeSubRoute.Methods(http.MethodGet).Path("/{employeeNumber}").HandlerFunc(GetEmployeeById)
	employeeSubRoute.Methods(http.MethodGet).HandlerFunc(GetEmployees)
	employeeSubRoute.Methods(http.MethodPost).HandlerFunc(AddEmployee)

	// office subroute
	officeSubRoute := mux.NewRouter().PathPrefix("/office").Subrouter().StrictSlash(true)
	officeSubRoute.Methods(http.MethodGet).Path("/{officeCode}").HandlerFunc(GetEmployeeInOffice)

	router.HandleFunc("/generate-qr-code", GenerateQRCode)

	router.PathPrefix("/employees").Handler(employeeSubRoute)
	router.PathPrefix("/office").Handler(officeSubRoute)

	return router
}
