# Stocks API with PostgreSQL

This is a simple Stocks API built with Go, Gorilla Mux, and PostgreSQL. The API allows you to create, read, update, and delete stock information.

## Features

- List all stocks
- Get a specific stock by ID
- Create a new stock
- Update an existing stock
- Delete a stock by ID

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Git

## Getting Started

### Clone the Repository

```
git clone

cd stocksapi
```

### Install Dependencies
```
go mod tidy
```
### Set Up PostgreSQL
- Install PostgreSQL if you haven't already.
- Create a new database:

```
createdb stocksdb
```

- Create a .env file in the root directory and add your PostgreSQL connection details:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=stocksdb

```

### Run the Application
```
go run main.go
The server will start on http://localhost:8888
```
## API Endpoints

### List All Stocks
```
Endpoint: GET /stocks
```

### Get Stock by ID
```
Endpoint: GET /stocks/{id}
```

### Create a New Stock
```
Endpoint: POST /stocks
```
### Request Body:

```
{
  "id": 1,
  "name": "Stock Name",
  "price": 100.0
}
```
### Update an Existing Stock
```
Endpoint: PUT /stocks/{id}
```

### Request Body:
```

{
  "id": 1,
  "name": "Updated Stock Name",
  "price": 150.0
}
```
### Delete a Stock
```
Endpoint: DELETE /stocks/{id}
```
### Project Structure
```
.
├── handlers
│   └── handlers.go
├── models
│   └── stock.go
├── service
│   └── stock_service.go
├── store
│   └── store.go

```