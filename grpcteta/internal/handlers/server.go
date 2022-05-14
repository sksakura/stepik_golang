package handlers

import (
	"context"
	"stepik_golang/grpcteta/pkg/orderservice"
)

type Server struct {
	Orders []*orderservice.Order
	orderservice.UnimplementedOrdersServer
}

func (s *Server) CreateOrder(ctx context.Context, order *orderservice.Order) (*orderservice.OrderResponse, error) {
	s.Orders = append(s.Orders, order)
	return &orderservice.OrderResponse{Success: true}, nil
}

func (s *Server) GetOrders(ctx context.Context, filter *orderservice.Filter) (*orderservice.OrdersList, error) {
	return &orderservice.OrdersList{List: s.Orders}, nil
}
