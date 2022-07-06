package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CreateMenuUsecase interface {
	CreateMenu(newMenu *model.Menu) error
}

type createMenuUsecase struct {
	repo repo.MenuRepository
}

func (c *createMenuUsecase) CreateMenu(newMenu *model.Menu) error {

	return c.repo.Create(newMenu)

}

func NewCreateMenuUseCase(repo repo.MenuRepository) CreateMenuUsecase {
	return &createMenuUsecase{
		repo: repo,
	}
}
