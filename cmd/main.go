package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/matheusmosca/arch-experiment/domain/usecases"
	"github.com/matheusmosca/arch-experiment/gateways/database/inmemory/tickets"
	"github.com/matheusmosca/arch-experiment/gateways/http/api"
)

func main() {
	fmt.Println("init")

	ticketRepo := tickets.NewRepository()

	getTicketUC := usecases.NewGetTicketUC(ticketRepo)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	getTicketHandler := api.NewGetTicket(getTicketUC)

	r.Route("/api/v1", func(api chi.Router) {
		api.Get("/tickets/{id}", getTicketHandler.GetTicket)
	})

	slog.Info("starting api in port 8080...")

	// Start the HTTP server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
