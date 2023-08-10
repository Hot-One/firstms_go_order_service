package service

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/genproto/user_service"
	"github.com/Hot-One/firstms_go_order_service/grpc/client"
	"github.com/Hot-One/firstms_go_order_service/pkg/logger"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *UserService {
	return &UserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *UserService) Create(ctx context.Context, req *user_service.UserCreate) (*user_service.User, error) {
	u.log.Info("====== User Create ======", logger.Any("req", req))

	resp, err := u.strg.User().Create(ctx, req)
	if err != nil {
		u.log.Error("Error While Create User: u.strg.User().Create", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *UserService) GetById(ctx context.Context, req *user_service.UserPrimaryKey) (*user_service.User, error) {
	u.log.Info("====== User Get By Id ======", logger.Any("req", req))

	resp, err := u.strg.User().GetByID(ctx, req)
	if err != nil {
		u.log.Error("Error While User Get By ID: u.strg.User().GetByID", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *UserService) GetList(ctx context.Context, req *user_service.UserGetListRequest) (*user_service.UserGetListResponse, error) {
	u.log.Info("====== User Get List ======", logger.Any("req", req))

	resp, err := u.strg.User().GetList(ctx, req)
	if err != nil {
		u.log.Error("Error While User Get List: u.strg.User().GetList", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *UserService) Update(ctx context.Context, req *user_service.UserUpdate) (*user_service.User, error) {
	u.log.Info("====== User Update ======", logger.Any("req", req))

	resp, err := u.strg.User().Update(ctx, req)
	if err != nil {
		u.log.Error("Error While User Update: u.strg.User().Update", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *UserService) Delete(ctx context.Context, req *user_service.UserPrimaryKey) (*emptypb.Empty, error) {
	u.log.Info("====== User Delete ======", logger.Any("req", req))

	err := u.strg.User().Delete(ctx, req)
	if err != nil {
		u.log.Error("Error While User Delete: u.strg.User().Delete", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
