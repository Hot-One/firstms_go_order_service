package storage

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/genproto/courier_service"
	"github.com/Hot-One/firstms_go_order_service/genproto/customer_service"
	"github.com/Hot-One/firstms_go_order_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Customer() CustomerRepoI
	Courier() CourierRepoI
}

type UserRepoI interface {
	Create(context.Context, *user_service.UserCreate) (*user_service.User, error)
	GetByID(context.Context, *user_service.UserPrimaryKey) (*user_service.User, error)
	GetList(context.Context, *user_service.UserGetListRequest) (*user_service.UserGetListResponse, error)
	Update(context.Context, *user_service.UserUpdate) (*user_service.User, error)
	Delete(context.Context, *user_service.UserPrimaryKey) error
}

type CustomerRepoI interface {
	Create(context.Context, *customer_service.CustomerCreate) (*customer_service.Customer, error)
	GetByID(context.Context, *customer_service.CustomerPrimaryKey) (*customer_service.Customer, error)
	GetList(context.Context, *customer_service.CustomerGetListRequest) (*customer_service.CustomerGetListResponse, error)
	Update(context.Context, *customer_service.CustomerUpdate) (*customer_service.Customer, error)
	Delete(context.Context, *customer_service.CustomerPrimaryKey) error
}

type CourierRepoI interface {
	Create(context.Context, *courier_service.CourierCreate) (*courier_service.Courier, error)
	GetByID(context.Context, *courier_service.CourierPrimaryKey) (*courier_service.Courier, error)
	GetList(context.Context, *courier_service.CourierGetListRequest) (*courier_service.CourierGetListResponse, error)
	Update(context.Context, *courier_service.CourierUpdate) (*courier_service.Courier, error)
	Delete(context.Context, *courier_service.CourierPrimaryKey) error
}
