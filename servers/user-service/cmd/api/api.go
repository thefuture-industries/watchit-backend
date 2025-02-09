package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"

	// "go-user-service/internal/common/packages/logging"
	"go-user-service/internal/common/packages"
)

type APIServer struct {
	addr   string
	db     *gorm.DB
	server *http.Server
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
	// subrouter := router.PathPrefix("/micro/user").Subrouter()

	// ------------------
	// Логирование server
	// ------------------
	logger, _ := zap.NewProduction()
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
