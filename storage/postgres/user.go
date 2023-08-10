package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/genproto/user_service"
	"github.com/Hot-One/firstms_go_order_service/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, req *user_service.UserCreate) (*user_service.User, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(id, name, phone_number)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &user_service.User{
		Id:          id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}, nil
}

func (r *UserRepo) GetByID(ctx context.Context, req *user_service.UserPrimaryKey) (*user_service.User, error) {

	// var whereField = "id"
	// if len(req.Username) > 0 {
	// 	whereField = "username"
	// 	req.Id = req.Username
	// }

	var (
		query string

		id           sql.NullString
		name         sql.NullString
		phone_number sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			phone_number
		FROM users
		WHERE = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&phone_number,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user_service.User{
		Id:          id.String,
		Name:        name.String,
		PhoneNumber: phone_number.String,
	}, nil
}

func (r *UserRepo) GetList(ctx context.Context, req *user_service.UserGetListRequest) (*user_service.UserGetListResponse, error) {

	var (
		resp   = &user_service.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			phone_number
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id           sql.NullString
			name         sql.NullString
			phone_number sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&phone_number,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user_service.User{
			Id:          id.String,
			Name:        name.String,
			PhoneNumber: phone_number.String,
		})
	}
	rows.Close()

	return resp, nil
}

func (r *UserRepo) Update(ctx context.Context, req *user_service.UserUpdate) (*user_service.User, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
			name = :name,
			phone_number = :phone_number
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &user_service.User{
		Id:          req.Id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}, nil
}

func (r *UserRepo) Delete(ctx context.Context, req *user_service.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
