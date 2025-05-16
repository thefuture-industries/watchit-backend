package api

import (
	"fmt"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/noneandundefined/vision-go"
	"gorm.io/gorm"

	"go-movie-service/internal/lib"
	"go-movie-service/internal/module/movie"
	"go-movie-service/internal/module/recommendation"
	"go-movie-service/internal/packages"
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
	// Создания router и префикс /microservice/movie-service
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/microservice/movie-service").Subrouter()

	// logger
	// logger, _ := zap.NewProduction()

	// Vision monitoring
	monitoring := vision.NewVision()
	subrouter.HandleFunc("/admin/vision", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/vision/index.html")
	}).Methods("GET")

	// Scalar
	subrouter.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	subrouter.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	subrouter.HandleFunc("/adm/doc", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:8011/microservice/movie-service/docs",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Movie Microservice",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	// ROUTERS PATHS
	errors := packages.NewErrors(monitoring)

	movie.NewHandler(monitoring, errors).RegisterRoutes(subrouter)
	recommendation.NewHandler(monitoring, errors).RegisterRoutes(subrouter)
	movie.NewHandler(monitoring, errors).MovieRoutes(subrouter)

	// Логирование server
	router.Use(packages.NewLogger().LoggerMiddleware)

	// подключаем CORS
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})
	allowCredentials := handlers.AllowCredentials()

	// Возращяем http ответ
	s.logger.Info(fmt.Sprintf("Listening on %s", s.addr))
	return http.ListenAndServe(s.addr, handlers.CORS(origins, methods, headers, allowCredentials)(router))
}
