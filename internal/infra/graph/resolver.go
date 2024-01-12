package graph

import (
	uc_ordercreate "github.com/GiovaniGitHub/clean-architecture/internal/usecase/order/create"
	uc_orderlist "github.com/GiovaniGitHub/clean-architecture/internal/usecase/order/list"
)

type Resolver struct {
	CreateOrderUseCase uc_ordercreate.CreateOrderUseCase
	ListOrderUseCase   uc_orderlist.ListOrderUseCase
}
