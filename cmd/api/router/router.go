package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/api/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/repository"
	"github.com/bootcamp-go/desafio-go-web/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ginEngine *gin.Engine
	tickets   []domain.Ticket
}

func NewRouter(ginEngine *gin.Engine, tickets []domain.Ticket) *Router {
	return &Router{
		ginEngine,
		tickets,
	}
}

func (r *Router) MapRoutes() {
	repository := repository.NewRepository(r.tickets)
	service := service.NewService(repository)
	ticketHandler := handler.NewService(service)
	ticketGroup := r.ginEngine.Group("/ticket")
	{
		ticketGroup.GET("getByCountry/:dest", ticketHandler.GetTicketsByCountry())
		ticketGroup.GET("getAverage/:dest", ticketHandler.AverageDestination())
	}
}
