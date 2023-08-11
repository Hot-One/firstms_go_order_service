package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/genproto/courier_service"
	"github.com/Hot-One/firstms_go_order_service/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CourierRepo struct {
	db *pgxpool.Pool
}

func NewCourierRepo(db *pgxpool.Pool) *CourierRepo {
	return &CourierRepo{
		db: db,
	}
}

func (r *CourierRepo) Create(ctx context.Context, req *courier_service.CourierCreate) (*courier_service.Courier, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO couriers(id, name, phone_number)
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

	return &courier_service.Courier{
		Id:          id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}, nil
}

func (r *CourierRepo) GetByID(ctx context.Context, req *courier_service.CourierPrimaryKey) (*courier_service.Courier, error) {

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
	)

	query = `
		SELECT
			id,
			name,
			phone_number
		FROM couriers
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&phone_number,
	)

	if err != nil {
		return nil, err
	}

	return &courier_service.Courier{
		Id:          id.String,
		Name:        name.String,
		PhoneNumber: phone_number.String,
	}, nil
}

func (r *CourierRepo) GetList(ctx context.Context, req *courier_service.CourierGetListRequest) (*courier_service.CourierGetListResponse, error) {

	var (
		resp   = &courier_service.CourierGetListResponse{}
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
		FROM couriers
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

		resp.Couriers = append(resp.Couriers, &courier_service.Courier{
			Id:          id.String,
			Name:        name.String,
			PhoneNumber: phone_number.String,
		})
	}
	rows.Close()

	return resp, nil
}

func (r *CourierRepo) Update(ctx context.Context, req *courier_service.CourierUpdate) (*courier_service.Courier, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			couriers
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

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, nil
	}

	return &courier_service.Courier{
		Id:          req.Id,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}, nil
}

func (r *CourierRepo) Delete(ctx context.Context, req *courier_service.CourierPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM couriers WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
