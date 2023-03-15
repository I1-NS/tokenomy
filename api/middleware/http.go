// Copyright 2023 © Tokenomy. All rights reserved.
package middleware

import "net/http"

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case http.MethodGet:
			next.ServeHTTP(w, r)
		case http.MethodOptions:
			return // Preflighted OPTIONS request method
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
}