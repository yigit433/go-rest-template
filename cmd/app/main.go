package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yigit433/go-rest-template/internal/configs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	configs.LoadEnv()
	
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	r := chi.NewRouter()
    r.Use(middleware.Logger)

    log.Info().Msg(fmt.Sprintf("🚀 Server is running on port %s", os.Getenv("PORT")))

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}