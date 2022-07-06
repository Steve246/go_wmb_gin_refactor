package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type UpdateMenuUsecase interface {
	// UpdateMenu(menuName *model.Menu, by map[string]interface{}) error
	UpdateMenu(menuName *model.Menu, id string) error
}

type updateMenuUsecase struct {
	repo repo.MenuRepository
}

func (c *updateMenuUsecase) UpdateMenu(menuName *model.Menu, id string) error {

	return c.repo.Update(menuName, id)

}

// func (c *updateMenuUsecase) UpdateMenu(menuName *model.Menu, by map[string]interface{}) error {

// 	return c.repo.Update(menuName, by)

// }

func NewUpdateMenuUseCase(repo repo.MenuRepository) UpdateMenuUsecase {
	return &updateMenuUsecase{
		repo: repo,
	}
}
