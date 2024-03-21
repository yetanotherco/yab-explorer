//go:build wireinject
// +build wireinject

package configs

import (
	"yab-explorer/controllers"
	"yab-explorer/repository"
	"yab-explorer/services"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var orderServiceSet = wire.NewSet(services.OrderServiceInit,
	wire.Bind(new(services.OrderService), new(*services.OrderServiceImpl)),
)

var orderRepositorySet = wire.NewSet(repository.OrderRepositoryInit,
	wire.Bind(new(repository.OrderRepository), new(*repository.OrderRepositoryImpl)),
)

var orderControllerSet = wire.NewSet(controllers.OrderControllerInit,
	wire.Bind(new(controllers.OrderController), new(*controllers.OrderControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, orderControllerSet, orderServiceSet, orderRepositorySet)
	return nil
}
