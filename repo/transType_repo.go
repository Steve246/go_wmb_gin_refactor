package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type TransTypeRepository interface {
	Create(tableType *model.Trans_Type) error
	Delete(tableType *model.Trans_Type) error
	Update(tableType *model.Trans_Type, id string) error
}

type transtypeRepository struct {
	db *gorm.DB
}

func (c *transtypeRepository) Update(tableType *model.Trans_Type, id string) error {

	var trans model.Trans_Type

	result := c.db.Model(tableType).First(&trans, "id = ?", id).Updates(trans)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *transtypeRepository) Delete(tableType *model.Trans_Type) error {
	result := c.db.Delete(tableType).Error
	return result
}

func (c *transtypeRepository) Create(tableType *model.Trans_Type) error {
	result := c.db.Create(tableType).Error

	return result
}

func NewTransTypeRepository(db *gorm.DB) TransTypeRepository {
	repo := new(transtypeRepository)
	repo.db = db
	return repo
}
