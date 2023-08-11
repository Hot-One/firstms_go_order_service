package service

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/category_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*category_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *CategoryService) Create(ctx context.Context, req *category_service.CategoryCreate) (*category_service.Category, error) {
	u.log.Info("====== Category Create ======", logger.Any("req", req))

	resp, err := u.strg.Category().Create(ctx, req)
	if err != nil {
		u.log.Error("Error While Create Category: u.strg.Category().Create", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) GetById(ctx context.Context, req *category_service.CategoryPrimaryKey) (*category_service.Category, error) {
	u.log.Info("====== Category Get By Id ======", logger.Any("req", req))

	resp, err := u.strg.Category().GetByID(ctx, req)
	if err != nil {
		u.log.Error("Error While Category Get By ID: u.strg.Category().GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) GetList(ctx context.Context, req *category_service.CategoryGetListRequest) (*category_service.CategoryGetListResponse, error) {
	u.log.Info("====== Category Get List ======", logger.Any("req", req))

	resp, err := u.strg.Category().GetList(ctx, req)
	if err != nil {
		u.log.Error("Error While Category Get List: u.strg.Category().GetList", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) Update(ctx context.Context, req *category_service.CategoryUpdate) (*category_service.Category, error) {
	u.log.Info("====== Category Update ======", logger.Any("req", req))

	resp, err := u.strg.Category().Update(ctx, req)
	if err != nil {
		u.log.Error("Error While Category Update: u.strg.Category().Update", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CategoryService) Delete(ctx context.Context, req *category_service.CategoryPrimaryKey) (*emptypb.Empty, error) {
	u.log.Info("====== Category Delete ======", logger.Any("req", req))

	err := u.strg.Category().Delete(ctx, req)
	if err != nil {
		u.log.Error("Error While Category Delete: u.strg.Category().Delete", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
