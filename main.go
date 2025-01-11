package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alucard017/go-template-backend/routes"
)

func main() {
	r := routes.Router()
	fmt.Println("Starting the Server on Port 3000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
