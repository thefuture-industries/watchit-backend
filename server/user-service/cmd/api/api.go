package api

import (
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/noneandundefined/vision-go"
	"gorm.io/gorm"

	"go-user-service/internal/lib"
	"go-user-service/internal/module/admin"
	"go-user-service/internal/module/auth"
	"go-user-service/internal/packages"
)

type APIServer struct {
	logger *lib.Logger
	addr   string
	db     *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		logger: lib.NewLogger(),
		addr:   addr,
		db:     db,
	}
}

func (s *APIServer) Run() error {
	// -------------------------------------
	// Создания router и префикс /micro/user
	// -------------------------------------
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/micro/user").Subrouter()

	// logger
	// logger, _ := zap.NewProduction()

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

	auth.NewHandler(monitoring, errors).RegisterRoutes(subrouter)
	admin.NewHandler(monitoring, errors).RegisterRoutes(subrouter)

	// ------------------
	// Middlewares
	// ------------------
	router.Use(packages.NewLogger().LoggerMiddleware)

	// ---------------
	// подключаем CORS
	// ---------------
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})

	// --------------------
	// Возращяем http ответ
	// --------------------
	s.logger.Info(fmt.Sprintf("Listening on %s", s.addr))
	return http.ListenAndServe(s.addr, handlers.CORS(origins, methods, headers)(router))
}
