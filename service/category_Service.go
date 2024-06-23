package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	GetById(ctx context.Context, categoryId int) web.CategoryResponse
	GetAll(ctx context.Context) []web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
