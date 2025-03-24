# Book Management API

## Overview
The Book Management API is a RESTful service built using Go and Gin framework, designed to manage books in a bookstore system. It supports CRUD operations on books and includes OpenAPI documentation powered by Swagger.

## Features
- Retrieve all books
- Fetch book details by ID
- Add a new book
- Update book details
- Delete a book
- Swagger UI for API documentation

## Technologies Used
- **Go** (Gin framework)
- **PostgreSQL** (Database)
- **Swagger** (API Documentation)
- **Kafka** (Event Streaming)
- **Redis** (Caching)

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- Go (>=1.18)
- PostgreSQL
- Redis
- Kafka (optional for event streaming)

### Clone the Repository
```sh
git clone https://github.com/your-repo/book-management.git
cd book-management
```

### Install Dependencies
```sh
go mod tidy
```

### Database Configuration
Update the database connection settings in `config.yaml`.

### Run the Application
```sh
go run main.go
```

### Generate Swagger Documentation
```sh
swag init
```

## API Documentation
Once the server is running, visit the following URL to access Swagger UI:
```
http://localhost:8080/swagger/index.html
```

## API Endpoints
### 1. Get All Books
**GET** `/books`

### 2. Get a Book by ID
**GET** `/books/{id}`

### 3. Create a New Book
**POST** `/books`

### 4. Update a Book
**PUT** `/books/{id}`

### 5. Delete a Book
**DELETE** `/books/{id}`

## Contributing
Feel free to contribute by submitting issues or pull requests.

## Some Useful Commands
### 1. docker-compose up -d
### 2. go install github.com/go-delve/delve/cmd/dlv@latest
### 3. dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./main.go
### 4. choco install yq
### 5. yq eval -o=json docs/swagger.yaml > docs/swagger.json
