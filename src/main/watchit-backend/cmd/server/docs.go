package main

import (
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func (s *httpServer) docs(r *mux.Router) {
	// scalar: not show in production
	go_env := os.Getenv("GO_ENV") == "DEV"
	if !go_env {
		return
	}

	// scalar
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))))
	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	r.HandleFunc("/adm/doc", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:8080/api/v1/docs",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Neomatica Server",
			},
			DarkMode: true,
		})

		if err != nil {
			s.logger.Error("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})
}
