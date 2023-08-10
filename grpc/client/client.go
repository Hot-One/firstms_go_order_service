package client

import "github.com/Hot-One/firstms_go_order_service/config"

type ServiceManagerI interface {
}

type grpcClients struct {
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	return grpcClients{}, nil
}
