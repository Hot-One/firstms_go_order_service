package service

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/customer_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*customer_service.UnimplementedCustomerServiceServer
}

func NewCustomerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *CustomerService {
	return &CustomerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *CustomerService) Create(ctx context.Context, req *customer_service.CustomerCreate) (*customer_service.Customer, error) {
	u.log.Info("====== Customer Create ======", logger.Any("req", req))

	resp, err := u.strg.Customer().Create(ctx, req)
	if err != nil {
		u.log.Error("Error While Create Customer: u.strg.Customer().Create", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CustomerService) GetById(ctx context.Context, req *customer_service.CustomerPrimaryKey) (*customer_service.Customer, error) {
	u.log.Info("====== Customer Get By Id ======", logger.Any("req", req))

	resp, err := u.strg.Customer().GetByID(ctx, req)
	if err != nil {
		u.log.Error("Error While Customer Get By ID: u.strg.Customer().GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CustomerService) GetList(ctx context.Context, req *customer_service.CustomerGetListRequest) (*customer_service.CustomerGetListResponse, error) {
	u.log.Info("====== Customer Get List ======", logger.Any("req", req))

	resp, err := u.strg.Customer().GetList(ctx, req)
	if err != nil {
		u.log.Error("Error While Customer Get List: u.strg.Customer().GetList", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CustomerService) Update(ctx context.Context, req *customer_service.CustomerUpdate) (*customer_service.Customer, error) {
	u.log.Info("====== Customer Update ======", logger.Any("req", req))

	resp, err := u.strg.Customer().Update(ctx, req)
	if err != nil {
		u.log.Error("Error While Customer Update: u.strg.Customer().Update", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *CustomerService) Delete(ctx context.Context, req *customer_service.CustomerPrimaryKey) (*emptypb.Empty, error) {
	u.log.Info("====== Customer Delete ======", logger.Any("req", req))

	err := u.strg.Customer().Delete(ctx, req)
	if err != nil {
		u.log.Error("Error While Customer Delete: u.strg.Customer().Delete", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
