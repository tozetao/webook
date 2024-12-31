// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"github.com/google/wire"
	"learn_go/webook/interaction/grpc"
	"learn_go/webook/interaction/repository"
	"learn_go/webook/interaction/repository/cache"
	"learn_go/webook/interaction/repository/dao"
	"learn_go/webook/interaction/service"
	"learn_go/webook/ioc"
)

// Injectors from wire.go:

func InitInteractionService() service.InteractionService {
	db := NewDB()
	interactionDao := dao.NewInteractionDao(db)
	cmdable := NewRedis()
	interactionCache := cache.NewInteractionCache(cmdable)
	interactionRepository := repository.NewInteractionRepository(interactionDao, interactionCache)
	interactionService := service.NewInteractionService(interactionRepository)
	return interactionService
}

func InitInteractionServiceServer() *grpc.InteractionServiceServer {
	db := NewDB()
	interactionDao := dao.NewInteractionDao(db)
	cmdable := NewRedis()
	interactionCache := cache.NewInteractionCache(cmdable)
	interactionRepository := repository.NewInteractionRepository(interactionDao, interactionCache)
	interactionService := service.NewInteractionService(interactionRepository)
	interactionServiceServer := grpc.NewInteractionServiceServer(interactionService)
	return interactionServiceServer
}

// wire.go:

var thirdPartySet = wire.NewSet(
	NewDB,
	NewRedis, ioc.NewLogger,
)

var interactionSet = wire.NewSet(service.NewInteractionService, repository.NewInteractionRepository, dao.NewInteractionDao, cache.NewInteractionCache)
