package datafunc

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func checkbookdatabase(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS bookstore")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("USE bookstore")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  )`)
	if err != nil {
		panic(err.Error())
	}
}	

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:qw@tcp(localhost:3306)/")
	if err != nil {
		panic(err.Error())
	}
	db.Ping()
	checkbookdatabase(db)
	db.Close()

	db, err = sql.Open("mysql", "root:qw@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
	db.Ping()
	return db
}