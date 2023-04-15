package repository

import (
	"Dzaakk/go-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(c context.Context, tx *sql.Tx, category domain.Category)
	FindById(c context.Context, tx *sql.Tx, categoryId int) (domain.Category, error )
	FindAll(c context.Context, tx *sql.Tx) []domain.Category
}