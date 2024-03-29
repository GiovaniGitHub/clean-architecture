package service

import (
	"context"

	"github.com/GiovaniGitHub/clean-architecture/internal/infra/grpc/pb"
	uc_ordercreate "github.com/GiovaniGitHub/clean-architecture/internal/usecase/order/create"
	uc_orderlist "github.com/GiovaniGitHub/clean-architecture/internal/usecase/order/list"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase uc_ordercreate.CreateOrderUseCase
	ListOrderUseCase   uc_orderlist.ListOrderUseCase
}

func NewOrderService(createOrderUC uc_ordercreate.CreateOrderUseCase, listOrderUC uc_orderlist.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUC,
		ListOrderUseCase:   listOrderUC,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	output, err := s.CreateOrderUseCase.Execute(uc_ordercreate.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, blank *pb.Blank) (*pb.OrderResponseList, error) {
	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	ordersResponse := make([]*pb.CreateOrderResponse, len(orders.Orders))
	for i, order := range orders.Orders {
		ordersResponse[i] = &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}

	return &pb.OrderResponseList{Orders: ordersResponse}, nil
}
