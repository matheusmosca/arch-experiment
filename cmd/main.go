package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/matheusmosca/arch-experiment/config"
	"github.com/matheusmosca/arch-experiment/domain/usecases"
	"github.com/matheusmosca/arch-experiment/extensions/xlog"
	secretsrepo "github.com/matheusmosca/arch-experiment/gateways/database/redis"
	"github.com/matheusmosca/arch-experiment/gateways/http/api"
	"github.com/matheusmosca/arch-experiment/gateways/http/middlewares"
	"github.com/redis/go-redis/v9"
)

func main() {
	logger := xlog.New(os.Stdout)

	var cfg config.Config

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 5,
	})
	defer redisClient.Close()

	secretsRepo := secretsrepo.NewSecretsRepo(redisClient)

	saveSecretUC := usecases.NewSaveSecretUC(secretsRepo)
	getSecretUC := usecases.NewGetSecretUC(secretsRepo)

	err := envconfig.Process("myapp", &cfg)
	if err != nil {
		logger.Fatal(err.Error())
	}

	r := chi.NewRouter()

	r.Use(
		middlewares.LoggerToContext(logger),
		middlewares.Logger,
		middleware.Recoverer,
	)

	saveSecretHandler := api.NewSaveSecret(saveSecretUC)
	getSecretHandler := api.NewGetSecret(getSecretUC)

	r.Route("/api/v1/secrets", func(api chi.Router) {
		api.Post("/", saveSecretHandler.SaveSecret)
		api.Get("/{id}", getSecretHandler.GetSecret)
	})

	httpSRV := http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf(":%s", cfg.API.Address),
		WriteTimeout:      cfg.API.Timeout,
		ReadTimeout:       cfg.API.Timeout,
		ReadHeaderTimeout: cfg.API.Timeout,
	}

	logger.Info("starting web server")

	// Start the HTTP server
	err = httpSRV.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
