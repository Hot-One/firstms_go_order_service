package service

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/courier_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CourierService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*courier_service.UnimplementedCourierServiceServer
}

func NewCourierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *CourierService {
	return &CourierService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *CourierService) Create(ctx context.Context, req *courier_service.CourierCreate) (*courier_service.Courier, error) {
	u.log.Info("====== Courier Create ======", logger.Any("req", req))

	resp, err := u.strg.Courier().Create(ctx, req)
	if err != nil {
		u.log.Error("Error while Create Courier: u.strg.Courier().Create", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CourierService) GetById(ctx context.Context, req *courier_service.CourierPrimaryKey) (*courier_service.Courier, error) {
	u.log.Info("====== Courier Get By Id ======", logger.Any("req", req))

	resp, err := u.strg.Courier().GetByID(ctx, req)
	if err != nil {
		u.log.Error("Error while Courier Get By ID: u.strg.Courier().GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CourierService) GetList(ctx context.Context, req *courier_service.CourierGetListRequest) (*courier_service.CourierGetListResponse, error) {
	u.log.Info("====== Courier Get List ======", logger.Any("req", req))

	resp, err := u.strg.Courier().GetList(ctx, req)
	if err != nil {
		u.log.Error("Error while Courier Get List: u.strg.Courier().GetList", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CourierService) Update(ctx context.Context, req *courier_service.CourierUpdate) (*courier_service.Courier, error) {
	u.log.Info("====== Courier Update ======", logger.Any("req", req))

	resp, err := u.strg.Courier().Update(ctx, req)
	if err != nil {
		u.log.Error("Error while Courier Update: u.strg.Courier().Update", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CourierService) Delete(ctx context.Context, req *courier_service.CourierPrimaryKey) (*emptypb.Empty, error) {
	u.log.Info("====== Courier Delete ======", logger.Any("req", req))

	err := u.strg.Courier().Delete(ctx, req)
	if err != nil {
		u.log.Error("Error while Courier Delete: u.strg.Courier().Delete", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
