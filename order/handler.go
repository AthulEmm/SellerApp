package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateOrderHandler creates a new order in the database
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Save the order in the database
	err = SaveOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateOrderStatusHandler updates the status of an existing order
func UpdateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the order ID from the URL path
	vars := mux.Vars(r)
	orderID := vars["id"]
	var orderStatus struct {
		Status string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&orderStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update the order status in the database
	err = UpdateOrderStatus(orderID, orderStatus.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetOrdersHandler retrieves orders from the database with pagination, filtering, and sorting options
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters for pagination, filtering, and sorting
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 10
	}
	filterField := r.FormValue("filterField")
	filterValue := r.FormValue("filterValue")
	sortField := r.FormValue("sortField")
	sortOrder := r.FormValue("sortOrder")
	// Get orders from the database with pagination, filtering, and sorting
	orders, err := GetOrders(page, limit, filterField, filterValue, sortField, sortOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert orders to JSON
	jsonData, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
