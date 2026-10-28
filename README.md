# BookStoreWithGo

A simple Go REST API for managing books in a MySQL database.

This project is a beginner-friendly bookstore backend built with Go, Gorilla Mux, and MySQL. It supports creating, listing, fetching, deleting books, and returning the total number of records.

## Features

- Create a book
- List all books
- Get a book by ID
- Delete a book
- Count books in the database
- Auto-create database/table if missing

## Tech Stack

- Go
- Gorilla Mux
- MySQL
- Go SQL driver for MySQL

## Project Structure

```text
bookStoreWithGo/
├── main.go
├── go.mod
├── go.sum
├── README.md
└── datafunc/
    ├── app.go
    ├── conennact.go
    ├── CreateBook.go
    ├── GetBooks.go
    ├── GetBookById.go
    ├── GetBookSize.go
    └── DeleteBook.go
```

## Prerequisites

Before running the app, make sure you have:

- Go installed
- MySQL running locally
- a database user with access to the `bookstore` database

The project currently uses this setup in the code:

```go
root:qw@tcp(localhost:3306)
```

So the expected MySQL settings are:

- host: `localhost`
- port: `3306`
- user: `root`
- password: `qw`
- database: `bookstore`

## Installation

Clone the project and run:

```bash
go mod tidy
```

Then start the server:

```bash
go run .
```

The API will run on:

```text
http://localhost:1337
```

## API Endpoints

### GET `/books`
Returns all books.

### GET `/books/{id}`
Returns a single book by ID.

### POST `/books`
Creates a new book.

Example request body:

```json
{
  "title": "Clean Code",
  "author": "Robert C. Martin"
}
```

### DELETE `/books/{id}`
Deletes a book by ID.

### GET `/booksSize`
Returns the total number of books.

Example response:

```json
{
  "count": 5
}
```

## Example Requests

List all books:

```bash
curl http://localhost:1337/books
```

Create a book:

```bash
curl -X POST http://localhost:1337/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Architecture","author":"Robert Martin"}'
```

Get one book:

```bash
curl http://localhost:1337/books/1
```

Delete a book:

```bash
curl -X DELETE http://localhost:1337/books/1
```

Get book count:

```bash
curl http://localhost:1337/booksSize
```

## Database Table

The app creates this table automatically if it does not exist:

```sql
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
```

## Notes

This is a learning project and a good example of a small CRUD API in Go. It is useful for practicing:

- HTTP handlers
- routing
- MySQL integration
- JSON request/response handling
- basic API validation

## Future Improvements

Some ideas for improving this project:

- move DB credentials to environment variables
- add unit/integration tests
- validate IDs more strictly
- use a cleaner project structure with repositories/services
- add proper logging and config management

## License

This project is for learning and personal use.
"} 
