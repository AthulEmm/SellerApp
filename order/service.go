package order

import (
	"encoding/json"
	"fmt"
	"go-mysql-http/db"
)

// SaveOrder saves the order in the database
func SaveOrder(order Order) error {
	// Convert order items to JSON string
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	db, err := db.GetConnection()
	if err != nil {
		return err
	}
	// Insert the order into the database
	_, err = db.Exec("INSERT INTO orders (id, status, items, total, currency_unit) VALUES (?, ?, ?, ?, ?)",
		order.ID, order.Status, string(itemsJSON), order.Total, order.CurrencyUnit)
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrderStatus updates the status of an existing order in the database
func UpdateOrderStatus(orderID, status string) error {
	db, err := db.GetConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderID)
	if err != nil {
		return err
	}

	return nil
}

// GetOrders retrieves orders from the database with pagination, filtering, and sorting
func GetOrders(page, limit int, filterField, filterValue, sortField, sortOrder string) ([]Order, error) {
	// Calculate the offset for pagination
	offset := (page - 1) * limit

	// Prepare the SQL query based on the filter and sort options
	query := "SELECT id, status, items, total, currency_unit FROM orders"

	if filterValue != "" && filterField != "" {
		query += " WHERE " + filterField + " = '" + filterValue + "'"
	}

	if sortField != "" {
		query += " ORDER BY " + sortField
		if sortOrder == "desc" {
			query += " DESC"
		} else {
			query += " ASC"
		}
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}
	// Execute the SQL query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		var itemsJSON string
		err := rows.Scan(&order.ID, &order.Status, &itemsJSON, &order.Total, &order.CurrencyUnit)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON string to order items
		err = json.Unmarshal([]byte(itemsJSON), &order.Items)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
