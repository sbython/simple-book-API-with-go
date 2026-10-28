package datafunc

import "database/sql"

type Book struct {
	DB *sql.DB
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookContent struct {
	ID     int    
	Title  string 
	Author string 	
}