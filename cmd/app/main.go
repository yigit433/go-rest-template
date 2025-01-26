package main

import (
	"fmt"
	"net/http"

	"os"

	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
	"github.com/yigit433/go-rest-template/internal/app/routes/tasks"
	"github.com/yigit433/go-rest-template/internal/configs"
	"github.com/yigit433/go-rest-template/internal/database"
	"github.com/yigit433/go-rest-template/internal/logging"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    logging.InitializeLogger("info", "app")
	configs.LoadEnv()
	database.NewMongoAdapter()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	tasksrouter.SetupRouter(r)

	logging.CustomLogger.Info("🚀 Server is running", map[string]interface{}{"port": os.Getenv("PORT"), "eventId": "startup"})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
