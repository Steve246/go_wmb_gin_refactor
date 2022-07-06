package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CreateMenuPriceUsecase interface {
	CreateMenuPrice(newMenuPrice *model.Menu_Price) error
}

type createMenuPriceUsecase struct {
	repo repo.PriceRepository
}

func (c *createMenuPriceUsecase) CreateMenuPrice(newMenuPrice *model.Menu_Price) error {

	return c.repo.Create(newMenuPrice)

}

func NewCreateMenuPriceUseCase(repo repo.PriceRepository) CreateMenuPriceUsecase {
	return &createMenuPriceUsecase{
		repo: repo,
	}
}
