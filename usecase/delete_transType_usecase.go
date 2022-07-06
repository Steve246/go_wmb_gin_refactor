package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type DeleteTransUsecase interface {
	DeleteTrans(newMenuPrice *model.Trans_Type) error
}

type deleteTransUsecase struct {
	repo repo.TransTypeRepository
}

func (c *deleteTransUsecase) DeleteTrans(transData *model.Trans_Type) error {

	return c.repo.Delete(transData)

}

func NewDeleteTransUsecase(repo repo.TransTypeRepository) DeleteTransUsecase {
	return &deleteTransUsecase{
		repo: repo,
	}
}
