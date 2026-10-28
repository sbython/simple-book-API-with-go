package datafunc

import (
	"net/http"
	"encoding/json"
)



func (b *Book) GetBooks(w http.ResponseWriter, r *http.Request) {
	var DB = b.DB
	var books []BookContent

	rows, err := DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {				
		var book BookContent
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}