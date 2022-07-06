package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type UpdateTransUsecase interface {
	// UpdateMenu(menuName *model.Menu, by map[string]interface{}) error
	UpdateTrans(trans *model.Trans_Type, id string) error
}

type updateTransUsecase struct {
	repo repo.TransTypeRepository
}

func (c *updateTransUsecase) UpdateTrans(trans *model.Trans_Type, id string) error {

	return c.repo.Update(trans, id)

}

func NewUpdateTransUsecase(repo repo.TransTypeRepository) UpdateTransUsecase {
	return &updateTransUsecase{
		repo: repo,
	}
}
