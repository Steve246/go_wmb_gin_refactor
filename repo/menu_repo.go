package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menuFood *model.Menu) error
	Delete(menuFood *model.Menu) error
	// Update(menuName *model.Menu, by map[string]interface{}) error
	Update(menuName *model.Menu, id string) error
}

type menuRepository struct {
	db *gorm.DB
}

func (c *menuRepository) Delete(menuFood *model.Menu) error {
	result := c.db.Delete(menuFood).Error
	return result
}

func (c *menuRepository) Update(menuName *model.Menu, id string) error {

	var menu model.Menu

	result := c.db.Model(menuName).First(&menu, "id = ?", id).Updates(menuName)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

// func (c *menuRepository) Update(menuName *model.Menu, by map[string]interface{}) error {

// 	result := c.db.Model(menuName).Updates(by)

// 	if err := result.Error; err != nil {
// 		return err
// 	}
// 	return nil

// }

func (c *menuRepository) Create(menuFood *model.Menu) error {
	err := c.db.Create(menuFood).Error

	if err != nil {
		return err
	}
	return nil
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
