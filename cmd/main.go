package main

import (
	"context"
	"net"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/grpc"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage/postgres"
)

func main() {
	// Load Config
	cfg := config.Load()

	// Setup Logger
	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	// Connect to DataBase
	pgconn, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("Postgres no connection: ", logger.Error(err))
	}
	defer pgconn.CloseDB()

	// Connect To Server
	srvc, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("Service no connection: ", logger.Error(err))
	}

	grpcServer := grpc.SetUpServver(cfg, log, pgconn, srvc)

	// Listen Port
	lis, err := net.Listen("tcp", cfg.OrderGRPCPort)
	if err != nil {
		log.Panic("net.listen", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.OrderGRPCPort))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Panic("grpcServer.Server", logger.Error(err))
	}
}
