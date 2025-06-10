package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dropboks/file-service/internal/domain/handler"
	"github.com/dropboks/file-service/internal/domain/service"
	"github.com/dropboks/file-service/internal/infrastructure/storage"
	"github.com/dropboks/file-service/pkg/constant"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Container   *dig.Container
	ServerReady chan bool
	Address     string
}

func (s *GRPCServer) Run() {
	err := s.Container.Invoke(func(
		grpcServer *grpc.Server,
		logger zerolog.Logger,
		db *pgxpool.Pool,
		svc service.UserService,
		miniio *storage.MinioStorage,
	) {
		defer db.Close()
		ctx := context.Background()
		err := miniio.InitBucket(ctx, constant.APP_BUCKET)
		if err != nil {
			log.Fatalf("Failed to init bucket: %v", err)
		}

		listen, err := net.Listen("tcp", s.Address)
		if err != nil {
			logger.Fatal().Msgf("failed to listen:%v", err)
		}
		handler.RegisterUserService(grpcServer, svc)

		go func() {
			if serveErr := grpcServer.Serve(listen); serveErr != nil {
				logger.Fatal().Msgf("gRPC server error: %v", serveErr)
			}
		}()
		logger.Info().Msg("gRPC server running in port :" + viper.GetString("app.port"))
		if s.ServerReady != nil {
			s.ServerReady <- true
		}

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Info().Msg("Shutting down gRPC server...")
		grpcServer.GracefulStop()
		logger.Info().Msg("gRPC server stopped gracefully.")
	})
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}
}
