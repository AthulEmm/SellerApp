package main

import (
	"go-mysql-http/db"
	"go-mysql-http/order"
	"log"
	"net/http"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.GetConnection()
	// Create a new HTTP server with Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/orders", order.CreateOrderHandler).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}/status", order.UpdateOrderStatusHandler).Methods(http.MethodPut)
	router.HandleFunc("/orders", order.GetOrdersHandler).Methods(http.MethodGet)

	// Start the server on port 8080
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
