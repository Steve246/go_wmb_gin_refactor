package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type UpdateMenuPriceUsecase interface {
	// UpdateMenu(menuName *model.Menu, by map[string]interface{}) error
	UpdateMenuPrice(menuName *model.Menu_Price, id string) error
}

type updateMenuPriceUsecase struct {
	repo repo.PriceRepository
}

func (c *updateMenuPriceUsecase) UpdateMenuPrice(menuName *model.Menu_Price, id string) error {

	return c.repo.Update(menuName, id)

}

// func (c *updateMenuUsecase) UpdateMenu(menuName *model.Menu, by map[string]interface{}) error {

// 	return c.repo.Update(menuName, by)

// }

func NewUpdateMenuPriceUsecase(repo repo.PriceRepository) UpdateMenuPriceUsecase {
	return &updateMenuPriceUsecase{
		repo: repo,
	}
}
