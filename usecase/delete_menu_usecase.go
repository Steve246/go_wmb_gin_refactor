package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type DeleteMenuUsecase interface {
	DeleteMenu(newMenu *model.Menu) error
}

type deleteMenuUsecase struct {
	repo repo.MenuRepository
}

func (c *deleteMenuUsecase) DeleteMenu(menuDelete *model.Menu) error {

	return c.repo.Delete(menuDelete)

}

func NewDeleteMenuUseCase(repo repo.MenuRepository) DeleteMenuUsecase {
	return &deleteMenuUsecase{
		repo: repo,
	}
}
