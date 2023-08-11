package postgres

import (
	"context"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/config"
	"github.com/Hot-One/firstms_go_order_service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	user     *UserRepo
	customer *CustomerRepo
	courier  *CourierRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, nil

}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}

func (s *Store) Courier() storage.CourierRepoI {
	if s.courier == nil {
		s.courier = NewCourierRepo(s.db)
	}

	return s.courier
}
