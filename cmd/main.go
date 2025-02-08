package main

import (
	"context"
	"fmt"

	"github.com/matheusmosca/arch-experiment/domain/usecases"
	"github.com/matheusmosca/arch-experiment/gateways/database/inmemory/tickets"
)

func main() {
	fmt.Println("init")

	ticketRepo := tickets.NewRepository()

	getTicketUC := usecases.NewGetTicketUC(ticketRepo)

	ticket, err := getTicketUC.GetTicket(context.Background(), usecases.GetTicketDTO{
		ID: "85403a74-441e-444f-8853-6a4530ec39dd",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(ticket)
}
