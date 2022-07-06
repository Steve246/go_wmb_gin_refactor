package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount *model.Discount) error
	Delete(discount *model.Discount) error
	// Update(discount *model.Discount, by map[string]interface{}) error
	Update(discount *model.Discount, id string) error
}

type discountRepository struct {
	db *gorm.DB
}

func (c *discountRepository) Update(discount *model.Discount, id string) error {

	var menu model.Menu

	result := c.db.Model(discount).First(&menu, "id = ?", id).Updates(discount)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

// func (c *discountRepository) Update(discount *model.Discount, by map[string]interface{}) error {

// 	result := c.db.Model(discount).Updates(by)

// 	if err := result.Error; err != nil {
// 		return err
// 	}
// 	return nil

// }

func (c *discountRepository) Delete(discount *model.Discount) error {

	result := c.db.Delete(discount).Error
	return result

}

func (c *discountRepository) Create(discount *model.Discount) error {
	result := c.db.Create(discount).Error

	return result
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo
}
