package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(menuPrice *model.Table) error
	Delete(menuPrice *model.Table) error
	Update(menuPrice *model.Table, id string) error
}

type tableRepository struct {
	db *gorm.DB
}

func (c *tableRepository) Update(menuPrice *model.Table, id string) error {

	var modelPrice model.Menu_Price

	result := c.db.Model(menuPrice).First(&modelPrice, "id = ?", id).Updates(menuPrice)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *tableRepository) Delete(tableMeja *model.Table) error {
	result := c.db.Delete(tableMeja).Error
	return result
}

func (c *tableRepository) Create(tableMeja *model.Table) error {
	result := c.db.Create(tableMeja).Error

	return result
}

func NewTableRepository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo
}