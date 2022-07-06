package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type DeleteTableUsecase interface {
	DeleteTable(newTable *model.Table) error
}

type deleteTableUsecase struct {
	repo repo.TableRepository
}

func (c *deleteTableUsecase) DeleteTable(newTable *model.Table) error {

	return c.repo.Delete(newTable)

}

func NewDeleteTableUsecase(repo repo.TableRepository) DeleteTableUsecase {
	return &deleteTableUsecase{
		repo: repo,
	}
}
