package main

import (
	"fmt"
	"net/http"
	"time"
	"watchit/httpx/config"
)

func (s *httpServer) httpStart() error {
	routes := s.routes()

	fmt.Println("\n" + `  _    _ _______ _______ _____
 | |  | |__   __|__   __|  __ \
 | |__| |  | |     | |  | |__) |
 |  __  |  | |     | |  |  ___/
 | |  | |  | |     | |  | |
 |_|  |_|  |_|     |_|  |_|

                                `)

	fmt.Printf("[INFO] Listening on :%d\n", config.HttpServerPort)

	httpServe := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.HttpServerPort),
		Handler:           routes,
		ReadTimeout:       7 * time.Second,
		WriteTimeout:      12 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 4 * time.Second,
	}

	return httpServe.ListenAndServe()
}
