package service

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/repository"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	ticketRepository repository.Repository
}

func NewService(ticketRepository repository.Repository) *service {
	return &service{
		ticketRepository: ticketRepository,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ticketsByDestination, err := s.ticketRepository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return ticketsByDestination, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	totalTickets, errGetAll := s.ticketRepository.GetAll(ctx)
	if errGetAll != nil {
		return 0, errGetAll
	}
	ticketsByDestination, errGetByDestination := s.ticketRepository.GetTicketByDestination(ctx, destination)
	if errGetByDestination != nil {
		return 0, errGetByDestination
	}

	totalPassengers := len(totalTickets)
	totalPassengersByDestination := len(ticketsByDestination)
	avg := float64(totalPassengersByDestination) / float64(totalPassengers)
	return avg, nil
}
