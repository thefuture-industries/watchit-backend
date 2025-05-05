// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"go-user-service/internal/module/admin"
	"go-user-service/internal/module/auth"
	"go-user-service/internal/module/sync"
	"go-user-service/internal/packages"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// -----------------------------------
	// Создания router и префикс /micro/user
	// -----------------------------------

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/micro/user").Subrouter()

	// logger
	logger, _ := zap.NewProduction()

	// ------
	// Scalar
	// ------
	subrouter.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))))
	subrouter.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	subrouter.HandleFunc("/adm/doc", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:8001/micro/user/docs",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "User Microservice",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	// Vision monitoring
	monitoring := vision.NewVision()
	subrouter.HandleFunc("/admin/vision", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/vision/index.html")
	}).Methods("GET")

	// -------------
	// ROUTERS PATHS
	// -------------
	errors := packages.NewErrors(monitoring)

	auth.NewHandler(monitoring, logger, errors).RegisterRoutes(subrouter)
	admin.NewHandler(monitoring, logger, errors).RegisterRoutes(subrouter)
	sync.NewHandler(monitoring, logger, errors).RegisterRoutes(subrouter)

	// ------------------
	// Middlewares
	// ------------------
	router.Use(packages.NewLogger(logger).LoggerMiddleware)

	// ---------------
	// подключаем CORS
	// ---------------
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})

	// --------------------
	// Возращяем http ответ
	// --------------------
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, handlers.CORS(origins, methods, headers)(router))
}
