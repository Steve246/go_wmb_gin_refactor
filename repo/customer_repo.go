package repo

import (
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type CustRepository interface {
	Create(menuPrice *model.Customer) error
	Delete(menuPrice *model.Customer) error
	Update(menuPrice *model.Customer, id string) error

	//bikin get first by dan update
	// Update(menuPrice *model.Customer, by map[string]interface{}) error

	// FindFirstBy(by map[string]interface{}) (model.Customer, error)
}

type custRepository struct {
	db *gorm.DB
}

func (c *custRepository) Delete(menuPrice *model.Customer) error {

	result := c.db.Delete(menuPrice).Error
	return result

}

// func (c *custRepository) FindFirstBy(by map[string]interface{}) (model.Customer, error) {
// 	var customer model.Customer
// 	result := c.db.Unscoped().Where(by).First(&customer)

// 	if err := result.Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return customer, nil
// 		} else {
// 			return customer, err
// 		}
// 	}
// 	return customer, nil
// }

func (c *custRepository) Update(menuPrice *model.Customer, id string) error {

	var customer model.Customer

	result := c.db.Model(menuPrice).First(&customer, "id = ?", id).Updates(menuPrice)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *custRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer).Error

	return result
}

func NewCustRepository(db *gorm.DB) CustRepository {
	repo := new(custRepository)
	repo.db = db
	return repo
}
