package datafunc

import (
	"net/http"
	"github.com/gorilla/mux"
	"unicode"
)

func (b *Book) DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var DB = b.DB
	//check if id is empty or not number
	if (id == ""|| id == "0" || !unicode.IsNumber(rune(id[0]))) {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	result, err := DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)	
}