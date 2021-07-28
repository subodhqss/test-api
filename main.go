package main

import (
	"fmt"

	"github.com/subodhqss/test-api/repository"
	"github.com/subodhqss/test-api/server"
)

func main() {
	repository.InitDBConnection()
	server.StartServer()
	fmt.Println("Hello world")
}
