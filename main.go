package main

import (
	"UserService/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler.HealthCheck)
	fmt.Println("User service running on :8000")
	http.ListenAndServe(":8000", nil)
}
