package di

import (
	"github.com/dropboks/file-service/config/logger"
	"github.com/dropboks/file-service/config/router"
	minioCon "github.com/dropboks/file-service/config/storage"
	"github.com/dropboks/file-service/internal/domain/repository"
	"github.com/dropboks/file-service/internal/domain/service"
	minioStorage "github.com/dropboks/file-service/internal/infrastructure/storage"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	if err := container.Provide(logger.New); err != nil {
		panic("Failed to provide logger: " + err.Error())
	}
	if err := container.Provide(minioCon.NewMinioConnection); err != nil {
		panic("Failed to provide minio Connection: " + err.Error())
	}
	if err := container.Provide(minioStorage.NewMinioStorage); err != nil {
		panic("Failed to provide minio storage interface: " + err.Error())
	}
	if err := container.Provide(repository.NewUserRepository); err != nil {
		panic("Failed to provide user repository: " + err.Error())
	}
	if err := container.Provide(service.NewUserService); err != nil {
		panic("Failed to provide user repository: " + err.Error())
	}
	if err := container.Provide(router.NewGRPC); err != nil {
		panic("Failed to provide gRPC Server: " + err.Error())
	}
	return container
}
