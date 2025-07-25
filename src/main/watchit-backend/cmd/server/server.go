package main

import (
	"fmt"
	"net/http"
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

	return http.ListenAndServe(fmt.Sprintf(":%d", config.HttpServerPort), routes)
}
