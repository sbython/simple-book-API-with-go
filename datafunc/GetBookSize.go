package datafunc

import (
	"net/http"
	"encoding/json"
)

func (b *Book) GetBookSize(w http.ResponseWriter, r *http.Request) {
	var DB = b.DB
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}