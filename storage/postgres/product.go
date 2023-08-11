package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/genproto/product_service"
	"github.com/Hot-One/firstms_go_order_service/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(ctx context.Context, req *product_service.ProductCreate) (*product_service.Product, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO products(id, name, price, category_id)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		helper.NewNullString(req.CategoryId),
	)

	if err != nil {
		return nil, err
	}

	return &product_service.Product{
		Id:         id,
		Name:       req.Name,
		Price:      req.Price,
		CategoryId: req.CategoryId,
	}, nil
}

func (r *ProductRepo) GetByID(ctx context.Context, req *product_service.ProductPrimaryKey) (*product_service.Product, error) {

	// var whereField = "id"
	// if len(req.Username) > 0 {
	// 	whereField = "username"
	// 	req.Id = req.Username
	// }

	var (
		query string

		id          sql.NullString
		name        sql.NullString
		price       sql.NullFloat64
		category_id sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			price,
			category_id
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&price,
		&category_id,
	)

	if err != nil {
		return nil, err
	}

	return &product_service.Product{
		Id:         id.String,
		Name:       name.String,
		Price:      price.Float64,
		CategoryId: category_id.String,
	}, nil
}

func (r *ProductRepo) GetList(ctx context.Context, req *product_service.ProductGetListRequest) (*product_service.ProductGetListResponse, error) {

	var (
		resp   = &product_service.ProductGetListResponse{}
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
			price,
			category_id
		FROM products
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
			id          sql.NullString
			name        sql.NullString
			price       sql.NullFloat64
			category_id sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&category_id,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product_service.Product{
			Id:         id.String,
			Name:       name.String,
			Price:      price.Float64,
			CategoryId: category_id.String,
		})
	}
	rows.Close()

	return resp, nil
}

func (r *ProductRepo) Update(ctx context.Context, req *product_service.ProductUpdate) (*product_service.Product, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			products
		SET
			name = :name,
			price = :price,
			category_id = :category_id
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"price":       req.Price,
		"category_id": helper.NewNullString(req.CategoryId),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, nil
	}

	return &product_service.Product{
		Id:         req.Id,
		Name:       req.Name,
		Price:      req.Price,
		CategoryId: req.CategoryId,
	}, nil
}

func (r *ProductRepo) Delete(ctx context.Context, req *product_service.ProductPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
