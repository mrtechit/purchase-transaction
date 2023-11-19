package httpserver

import (
	"net/http"
)

func Handler() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
