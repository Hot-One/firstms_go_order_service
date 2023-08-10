package storage

import (
	"context"

	"github.com/Hot-One/firstms_go_order_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(context.Context, *user_service.UserCreate) (*user_service.User, error)
	GetByID(context.Context, *user_service.UserPrimaryKey) (*user_service.User, error)
	GetList(context.Context, *user_service.UserGetListRequest) (*user_service.UserGetListResponse, error)
	Update(context.Context, *user_service.UserUpdate) (*user_service.User, error)
	Delete(context.Context, *user_service.UserPrimaryKey) error
}
