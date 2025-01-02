package main

import (
	"log"
	nethttp "net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/cors"

	"backend/internal/config"
	"backend/internal/delivery/http"
	"backend/internal/domain" // Импортируем пакет с моделями
	"backend/internal/repository"
	"backend/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// Автоматическое создание таблиц, если они не существуют
	db.AutoMigrate(&domain.Tire{}, &domain.Admin{})

	// Initialize repositories
	tireRepo := repository.NewTireRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	// Initialize use cases
	tireUsecase := usecase.NewTireUsecase(tireRepo)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)

	// Initialize HTTP handlers
	tireHandler := http.NewTireHandler(tireUsecase)
	adminHandler := http.NewAdminHandler(adminUsecase)

	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Настройка CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Разрешить все источники
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Максимальное время жизни предварительных запросов
	})
	r.Use(corsMiddleware.Handler)

	// Register routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/tires", func(r chi.Router) {
			r.Get("/", tireHandler.GetAllTires)
			r.Get("/{id}", tireHandler.GetTireByID)
			r.Post("/", tireHandler.CreateTire)
			r.Put("/{id}", tireHandler.UpdateTire)
			r.Delete("/{id}", tireHandler.DeleteTire)
		})

		r.Route("/admins", func(r chi.Router) {
			r.Get("/", adminHandler.GetAllAdmins)
			r.Post("/login", adminHandler.GetAdminByName)
			r.Post("/", adminHandler.CreateAdmin)
			r.Put("/{id}", adminHandler.UpdateAdmin)
			r.Delete("/{id}", adminHandler.DeleteAdmin)
		})

		r.Route("/health", func(r chi.Router) {
			r.Get("/", http.GetHealth)
		})
	})

	// Start server
	log.Printf("Server started on %s", cfg.ServerAddress)
	log.Fatal(nethttp.ListenAndServe(cfg.ServerAddress, r))
}
