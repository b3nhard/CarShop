package handlers

import "net/http"

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
		w.WriteHeader(200)
	}
}
