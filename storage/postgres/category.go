package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hot-One/firstms_go_order_service/genproto/category_service"
	"github.com/Hot-One/firstms_go_order_service/pkg/helper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Create(ctx context.Context, req *category_service.CategoryCreate) (*category_service.Category, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO categories(id, name)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)

	if err != nil {
		return nil, err
	}

	return &category_service.Category{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (r *CategoryRepo) GetByID(ctx context.Context, req *category_service.CategoryPrimaryKey) (*category_service.Category, error) {

	// var whereField = "id"
	// if len(req.Username) > 0 {
	// 	whereField = "username"
	// 	req.Id = req.Username
	// }

	var (
		query string

		id   sql.NullString
		name sql.NullString
	)

	query = `
		SELECT
			id,
			name
		FROM categories
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
	)

	if err != nil {
		return nil, err
	}

	return &category_service.Category{
		Id:   id.String,
		Name: name.String,
	}, nil
}

func (r *CategoryRepo) GetList(ctx context.Context, req *category_service.CategoryGetListRequest) (*category_service.CategoryGetListResponse, error) {

	var (
		resp   = &category_service.CategoryGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name
		FROM categories
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
			id   sql.NullString
			name sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
		)

		if err != nil {
			return nil, err
		}

		resp.Categorys = append(resp.Categorys, &category_service.Category{
			Id:   id.String,
			Name: name.String,
		})
	}
	rows.Close()

	return resp, nil
}

func (r *CategoryRepo) Update(ctx context.Context, req *category_service.CategoryUpdate) (*category_service.Category, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			categories
		SET
			name = :name
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, nil
	}

	return &category_service.Category{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}

func (r *CategoryRepo) Delete(ctx context.Context, req *category_service.CategoryPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM categories WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
