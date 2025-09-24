package usecase

import (
	"github.com/dabraga/Desafio03_Clean-Architecture/internal/entity"
	"github.com/dabraga/Desafio03_Clean-Architecture/pkg/events"
)

type ListedOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListedOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []ListedOrdersOutputDTO{}, err
	}

	var listedOrders []ListedOrdersOutputDTO
	for _, order := range orders {
		listedOrders = append(listedOrders, ListedOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.OrdersListed.SetPayload(listedOrders)
	c.EventDispatcher.Dispatch(c.OrdersListed)
	return listedOrders, nil
}
