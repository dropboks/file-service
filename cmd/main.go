package main

import (
	"github.com/dropboks/file-service/cmd/bootstrap"
	"github.com/dropboks/file-service/cmd/server"
	"github.com/spf13/viper"
)

func main() {
	container := bootstrap.Run()
	serverReady := make(chan bool)
	grpcServer := server.GRPCServer{
		Container:   container,
		ServerReady: serverReady,
		Address:     ":" + viper.GetString("app.port"),
	}
	grpcServer.Run()
	<-serverReady
}
