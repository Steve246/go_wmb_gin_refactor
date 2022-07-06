package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CreateTableUsecase interface {
	CreateTable(newTable *model.Table) error
}

type createTableUsecase struct {
	repo repo.TableRepository
}

func (c *createTableUsecase) CreateTable(newTable *model.Table) error {

	return c.repo.Create(newTable)

}

func NewCreateTableUseCase(repo repo.TableRepository) CreateTableUsecase {
	return &createTableUsecase{
		repo: repo,
	}
}
