# SellerApp : Order Management Service

This project is an order management service implemented in Go with a MySQL database and JSON-HTTP REST APIs. 
It allows you to add orders with different status values, update the status of existing orders, and provides 
support for pagination, filtering, and sorting on all order fields, including item fields.


## Getting started

## Requirements

To run this project, you need to have the following software installed:

1. Go 
2. MySQL database server
3. Git (for cloning the repository)

To get a local copy up and running follow these simple steps.


### Clone repository

    git clone  https://github.com/AthulEmm/SellerApp.git

### Setting up database

**MySql**

Create DB name: CREATE DATABASE order_management

Create table  : CREATE TABLE `orders` (

  `id` varchar(255) DEFAULT NULL,

  `status` varchar(100) DEFAULT NULL,

  `items` varchar(100) DEFAULT NULL,

  `total` double DEFAULT NULL,

  `currency_unit` varchar(20) DEFAULT NULL)


- Update dbUsername , dbPassword and other const values in db.go with your connection values


### APIs


##  Create a new order
/orders/                 - POST  - Creates orders

Description  : Creates a new order with the provided payload

    {
    "id": "Id_123456",
    "status": "FRESH_INVOICE",
    "items": [{
        "id": "123456 ",
        "description": "a product description",
        "price": 12.40,
        "quantity": 1
    }],
    "total": 100.40,
    "currencyUnit": "USD"
    }
## Get a list of orders
Endpoint: GET /orders?page={value}&limit={value}&filterField={value}&filterValue={value}&sortField={value}&sortOrder={asc/desc}

Description: Retrieves a list of orders with support for pagination, filtering, and sorting.

Query Parameters:

page            : Specifies the page number for pagination.

limit           : Specifies the number of orders to retrieve per page .

filterField     : Specifies the field to filter

filterValue     : Specifies the value to filter the orders                  
                  Example:  filterField=status filterValue=PENDING_INVOICE                    

sortField       : Specifies the field to sort the orders. Example: sortField=total.

sortOrder       : Specifies the order of sorting 

Example         :  http://localhost:8080/orders?page=1&limit=5&filterField=status&filterValue=PENDING_INVOICE

Example         :  http://localhost:8080/orders?page=1&limit=5&sortField=id&sortOrder=desc

Response body:

    [
    {
        "id": "A-123456",
        "status": "PENDING_INVOICE",
        "items": [
            {
                "id": "123456 ",
                "description": "a product description",
                "price": 12.4,
                "quantity": 1
            }
        ],
        "total": 12.4,
        "currencyUnit": "USD"
    }   
]


##  Get a specific order by ID
Endpoint: GET /orders?filterField={value}&filterValue={value}

Description: Retrieves a specific order by its ID.

Path Parameters:
    filterField : specify as id    
    filterValue : id of the order to retrive    
    Example:  http://localhost:8080/orders?filterField=id&filterValue=abcdef-123456    

Response:


    {
        "id": "abcdef-123456",
        "status": "Shipped_INVOICE",
        "items": [
            {
                "id": "123456 ",
                "description": "a product description",
                "price": 12.4,
                "quantity": 1
            }
        ],
        "total": 12.4,
        "currencyUnit": "USD"
    }
    

## Update the status of an existing order

Endpoint: PUT /orders/{ID}/status

Description: Updates the status of an existing order.

Path Parameters:

`id`: The ID of the order to update.

Request Body:

    {
    "status": "Shipped_INVOICE"
    }

Response:
    
    {
        "id": "abcdef-123456",
        "status": "Shipped_INVOICE",
        "items": [
            {
                "id": "123456 ",
                "description": "a product description",
                "price": 12.4,
                "quantity": 1
            }
        ],
        "total": 12.4,
        "currencyUnit": "USD"
    },




