package bootstrap

import (
	"github.com/dropboks/file-service/cmd/di"
	"github.com/dropboks/file-service/config/env"
	"go.uber.org/dig"
)

func Run() *dig.Container {
	env.Load()
	container := di.BuildContainer()
	return container
}
