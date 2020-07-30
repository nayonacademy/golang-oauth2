package handlers

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	return mux
}
