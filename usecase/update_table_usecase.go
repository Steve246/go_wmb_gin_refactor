package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type UpdateTableUsecase interface {
	// UpdateMenu(menuName *model.Menu, by map[string]interface{}) error
	UpdateTable(tableName *model.Table, id string) error
}

type updateTableUsecase struct {
	repo repo.TableRepository
}

func (c *updateTableUsecase) UpdateTable(tableName *model.Table, id string) error {

	return c.repo.Update(tableName, id)

}

func NewUpdateTableUsecase(repo repo.TableRepository) UpdateTableUsecase {
	return &updateTableUsecase{
		repo: repo,
	}
}
