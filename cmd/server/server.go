package server

import (
	"net/http"
	"os"

	chiPrometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/sithumonline/red-timer/api/router"
	"github.com/sithumonline/red-timer/api/websocket"
	"github.com/sithumonline/red-timer/config"
	_ "github.com/sithumonline/red-timer/docs"
)

// @title Go Puso
// @version 0.0.1
// @description Golang Sri Lanka template repo

// @contact.name Golang Sri Lanka
// @contact.email golangsrilanka@mail.com

// @host golangsrilanka.github.io/go-puso
// @BasePath /api/v1

var RunServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start go-puso server",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	port := config.GetEnv("server.PORT")

	r.Use(chiPrometheus.NewMiddleware("go-puso"))
	r.Handle("/metrics", promhttp.Handler())
	r.Mount("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))
	r.Mount("/healthz", router.HealthRoute())
	r.Mount("/api/v1", router.Router())

	r.Get("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(r.RequestURI, http.FileServer(http.Dir("ui"))).ServeHTTP(w, r)
	})
	hub := websocket.NewHub()
	go hub.Run()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	log.Info("red-timer started listening on ", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
