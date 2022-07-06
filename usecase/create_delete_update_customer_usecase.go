package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CrudCustomer interface {
	CreateCustomer(customer *model.Customer) error
	DeleteCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer, id string) error
}

type crudCustomer struct {
	repo repo.CustRepository
}

func (c *crudCustomer) UpdateCustomer(customer *model.Customer, id string) error {

	return c.repo.Update(customer, id)

}

func (c *crudCustomer) DeleteCustomer(customer *model.Customer) error {
	return c.repo.Delete(customer)
}

func (c *crudCustomer) CreateCustomer(customer *model.Customer) error {
	return c.repo.Create(customer)
}

func NewCrudCustomer(repo repo.CustRepository) CrudCustomer {
	return &crudCustomer{
		repo: repo,
	}
}
