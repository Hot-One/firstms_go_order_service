package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/genproto/customer_service"
	"github.com/Hot-One/firstms_go_order_service/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CustomerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) Create(ctx context.Context, req *customer_service.CustomerCreate) (*customer_service.Customer, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO customers(id, name, phone)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Phone,
	)

	if err != nil {
		return nil, err
	}

	return &customer_service.Customer{
		Id:    id,
		Name:  req.Name,
		Phone: req.Phone,
	}, nil
}

func (r *CustomerRepo) GetByID(ctx context.Context, req *customer_service.CustomerPrimaryKey) (*customer_service.Customer, error) {

	// var whereField = "id"
	// if len(req.Username) > 0 {
	// 	whereField = "username"
	// 	req.Id = req.Username
	// }

	var (
		query string

		id    sql.NullString
		name  sql.NullString
		phone sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			phone
		FROM customers
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&phone,
	)

	if err != nil {
		return nil, err
	}

	return &customer_service.Customer{
		Id:    id.String,
		Name:  name.String,
		Phone: phone.String,
	}, nil
}

func (r *CustomerRepo) GetList(ctx context.Context, req *customer_service.CustomerGetListRequest) (*customer_service.CustomerGetListResponse, error) {

	var (
		resp   = &customer_service.CustomerGetListResponse{}
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
			phone
		FROM customers
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
			id    sql.NullString
			name  sql.NullString
			phone sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&phone,
		)

		if err != nil {
			return nil, err
		}

		resp.Customers = append(resp.Customers, &customer_service.Customer{
			Id:    id.String,
			Name:  name.String,
			Phone: phone.String,
		})
	}
	rows.Close()

	return resp, nil
}

func (r *CustomerRepo) Update(ctx context.Context, req *customer_service.CustomerUpdate) (*customer_service.Customer, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			customers
		SET
			name = :name,
			phone = :phone
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
		"phone": req.Phone,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, nil
	}

	return &customer_service.Customer{
		Id:    req.Id,
		Name:  req.Name,
		Phone: req.Phone,
	}, nil
}

func (r *CustomerRepo) Delete(ctx context.Context, req *customer_service.CustomerPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM customers WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
