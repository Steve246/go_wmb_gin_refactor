package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CrudDiscount interface {
	CreateDiscount(discount *model.Discount) error
	DeleteDiscount(discount *model.Discount) error
	UpdateDiscount(discount *model.Discount, id string) error
}

type crudDiscount struct {
	repo repo.DiscountRepository
}

func (c *crudDiscount) UpdateDiscount(discount *model.Discount, id string) error {

	return c.repo.Update(discount, id)

}

func (c *crudDiscount) DeleteDiscount(discount *model.Discount) error {
	return c.repo.Delete(discount)
}

func (c *crudDiscount) CreateDiscount(discount *model.Discount) error {
	return c.repo.Create(discount)
}

func NewCrudDiscount(repo repo.DiscountRepository) CrudDiscount {
	return &crudDiscount{
		repo: repo,
	}
}
