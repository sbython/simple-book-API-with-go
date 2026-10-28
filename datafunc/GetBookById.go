package datafunc

import (
	"net/http"
	"encoding/json"	
	"github.com/gorilla/mux"
	"database/sql"
	"unicode"
)

func (b *Book) GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var DB = b.DB
	var book BookContent
	//check if id is empty or not number
	if (id == ""|| id == "0" || !unicode.IsNumber(rune(id[0]))) {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	err := DB.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)	

}