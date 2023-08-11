package service

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/product_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*product_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *ProductService) Create(ctx context.Context, req *product_service.ProductCreate) (*product_service.Product, error) {
	u.log.Info("====== Product Create ======", logger.Any("req", req))

	resp, err := u.strg.Product().Create(ctx, req)
	if err != nil {
		u.log.Error("Error While Create Product: u.strg.Product().Create", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) GetById(ctx context.Context, req *product_service.ProductPrimaryKey) (*product_service.Product, error) {
	u.log.Info("====== Product Get By Id ======", logger.Any("req", req))

	resp, err := u.strg.Product().GetByID(ctx, req)
	if err != nil {
		u.log.Error("Error While Product Get By ID: u.strg.Product().GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) GetList(ctx context.Context, req *product_service.ProductGetListRequest) (*product_service.ProductGetListResponse, error) {
	u.log.Info("====== Product Get List ======", logger.Any("req", req))

	resp, err := u.strg.Product().GetList(ctx, req)
	if err != nil {
		u.log.Error("Error While Product Get List: u.strg.Product().GetList", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) Update(ctx context.Context, req *product_service.ProductUpdate) (*product_service.Product, error) {
	u.log.Info("====== Product Update ======", logger.Any("req", req))

	resp, err := u.strg.Product().Update(ctx, req)
	if err != nil {
		u.log.Error("Error While Product Update: u.strg.Product().Update", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *ProductService) Delete(ctx context.Context, req *product_service.ProductPrimaryKey) (*emptypb.Empty, error) {
	u.log.Info("====== Product Delete ======", logger.Any("req", req))

	err := u.strg.Product().Delete(ctx, req)
	if err != nil {
		u.log.Error("Error While Product Delete: u.strg.Product().Delete", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
