# CRUD Operations with Go and MySQL

Simple CRUD (Create, Read, Update, Delete) operations for managing books using Go and MySQL.

## Prerequisites

- Go: [Download Go](https://go.dev/dl/)
- MySQL: [Download MySQL](https://dev.mysql.com/downloads/installer/)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/Crud-Go-MySQL.git
cd Crud-Go-MySQL
```

2. Create `.env` file:
```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
SERVER_PORT=
```

3. Install dependencies:
```bash
go mod tidy
```

4. Setup MySQL Database:
```sql
CREATE DATABASE crud_golang;

USE crud_golang;

CREATE TABLE books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255),
  author VARCHAR(255),
  price DECIMAL(10, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

5. Run:
```bash
go run cmd/main.go
```

## API Endpoints

### POST /books
Create a book:
```json
{
  "title": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "price": 40.99
}
```

### GET /books
Get all books

### GET /books/{id}
Get book by ID

### PUT /books/{id}
Update a book:
```json
{
  "title": "The Go Programming Language - Updated",
  "author": "Alan A. A. Donovan",
  "price": 45.00
}
```

### DELETE /books/{id}
Delete a book

## Project Structure
```
Crud-Go-MySQL/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── database/
│   │   └── db.go
│   ├── handlers/
│   │   └── book_handler.go
│   ├── repository/
│   │   └── book_repository.go
│   └── routes/
│       └── route.go
├── models/
│   └── book.go
├── .env
└── README.md
```

## Dependencies
```go
module github.com/boysimon10/go-mysql-crud

go 1.23.2

require (
    github.com/go-sql-driver/mysql v1.8.1
    github.com/gorilla/mux v1.8.1
    github.com/joho/godotenv v1.5.1
)
```