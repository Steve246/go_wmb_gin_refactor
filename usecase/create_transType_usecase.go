package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CreateTransUsecase interface {
	CreateTrans(newMenu *model.Trans_Type) error
}

type createTransUsecase struct {
	repo repo.TransTypeRepository
}

func (c *createTransUsecase) CreateTrans(newMenu *model.Trans_Type) error {

	return c.repo.Create(newMenu)

}

func NewCreateTransUsecase(repo repo.TransTypeRepository) CreateTransUsecase {
	return &createTransUsecase{
		repo: repo,
	}
}
