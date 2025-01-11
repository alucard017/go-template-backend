package main

import (
	"fmt"
	"go-todo-backend/routes"
	"net/http"
)

func main() {
	r := routes.Router()
	fmt.Println("Starting the Server on Port 3000")
	log.fatal(http.ListenAndServer(":9000", r))
}
