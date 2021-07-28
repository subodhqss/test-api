package server

import (
	"fmt"
	"net/http"

	"github.com/subodhqss/test-api/controllers"
)

func StartServer() {
	r := controllers.NewRouter()
	server := &http.Server{
		Addr:    ":4900",
		Handler: r,
	}
	fmt.Println("Server is running at ", server.Addr)
	server.ListenAndServe()
}
