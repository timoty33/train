package server

import (
	chi "github.com/go-chi/chi/v5"

	"net/http"
)

func createServer() http.Handler {
	server := chi.NewRouter()

	// SEM middleware: o ResponseWriter chega original ao Upgrade.
	// HandleFunc (não Get) porque o WebTransport usa o método CONNECT.
	server.HandleFunc("/wt", handleWT(wt, hub))

}
