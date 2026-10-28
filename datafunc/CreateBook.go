package datafunc

import (
	"net/http"
	"encoding/json"
)



func (b *Book) CreateBook(w http.ResponseWriter, r *http.Request)  {
	var bookRequest BookRequest
	var DB = b.DB
	err := json.NewDecoder(r.Body).Decode(&bookRequest)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if bookRequest.Title == "" || bookRequest.Author == "" {
		http.Error(w, "Title and Author are required fields", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec("INSERT INTO books (title, author) VALUES (?, ?)", bookRequest.Title, bookRequest.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}	