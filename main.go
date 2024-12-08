package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/enricomilli/neat-server/api"
	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	"github.com/enricomilli/neat-server/msg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		msg.MsgTelegram("Error from Neatblock API: \n" + err.Error())
		log.Fatal(err)
	}
}

func run() error {

	router := chi.NewRouter()

	godotenv.Load()

	setupCORS(router)
	setupMiddlewares(router)
	setupRoutes(router)

	port := getPort()
	slog.Info("Server running", "port", port)
	return http.ListenAndServe(":"+port, router)
}

func setupMiddlewares(router *chi.Mux) {

	// recommended stack
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	// rate limiter
	router.Use(httprate.Limit(
		200,
		40*time.Second,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			apiutil.ResponseWithError(w, http.StatusTooManyRequests, "Rate limited. Slow down.")
			return
		}),
	))
}

func setupCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"https://neatblock.org",
			"https://www.neatblock.org",
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"PUT",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
			"Accept",
			"Origin",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func setupRoutes(router *chi.Mux) {
	api.CreateRoutes(router)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		slog.Warn("Could not get port from env, setting default to 8080")
		port = "8080"
	}
	return port
}
