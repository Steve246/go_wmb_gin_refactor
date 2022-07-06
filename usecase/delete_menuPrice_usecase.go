package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type DeleteMenuPriceUsecase interface {
	DeletePrice(newMenuPrice *model.Menu_Price) error
}

type deleteMenuPriceUsecase struct {
	repo repo.PriceRepository
}

func (c *deleteMenuPriceUsecase) DeletePrice(priceDelete *model.Menu_Price) error {

	return c.repo.Delete(priceDelete)

}

func NewDeleteMenuPriceUsecase(repo repo.PriceRepository) DeleteMenuPriceUsecase {
	return &deleteMenuPriceUsecase{
		repo: repo,
	}
}
