package service

import (
	"Dzaakk/go-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(c context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(c context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(c context.Context, categoryId int)
	FindById(c context.Context, categoryId int) web.CategoryResponse
	FindAll(c context.Context) []web.CategoryResponse
}
