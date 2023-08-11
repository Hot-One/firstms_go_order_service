package grpc

import (
	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/category_service"
	"github.com/Hot-One/firstms_go_order_service/genproto/courier_service"
	"github.com/Hot-One/firstms_go_order_service/genproto/customer_service"
	"github.com/Hot-One/firstms_go_order_service/genproto/user_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/grpc/service"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServver(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, srvc))
	customer_service.RegisterCustomerServiceServer(grpcServer, service.NewCustomerService(cfg, log, strg, srvc))
	courier_service.RegisterCourierServiceServer(grpcServer, service.NewCourierService(cfg, log, strg, srvc))
	category_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
