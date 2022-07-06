package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type PriceRepository interface {
	Create(menuPrice *model.Menu_Price) error
	Delete(menuPrice *model.Menu_Price) error
	Update(menuPrice *model.Menu_Price, id string) error
}

type priceRepository struct {
	db *gorm.DB
}

func (c *priceRepository) Update(menuPrice *model.Menu_Price, id string) error {

	var modelPrice model.Menu_Price

	result := c.db.Model(menuPrice).First(&modelPrice, "id = ?", id).Updates(menuPrice)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *priceRepository) Delete(menuPrice *model.Menu_Price) error {
	result := c.db.Delete(menuPrice).Error
	return result
}

func (c *priceRepository) Create(menuPrice *model.Menu_Price) error {
	result := c.db.Create(menuPrice).Error

	return result
}

func NewPriceRepository(db *gorm.DB) PriceRepository {
	repo := new(priceRepository)
	repo.db = db
	return repo
}
