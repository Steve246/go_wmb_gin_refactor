package repo

import (
	"errors"
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(tBil *model.Bill) error
	FindByIdBill(id string) (model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func (c *billRepository) FindByIdBill(id string) (model.Bill, error) {
	var bill model.Bill
	result := c.db.Unscoped().First(&bill, "id = ?", id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}

	}

	return bill, nil
}

func (c *billRepository) Create(tBil *model.Bill) error {
	result := c.db.Create(tBil).Error

	return result
}

func NewBillRepository(db *gorm.DB) BillRepository {
	repo := new(billRepository)
	repo.db = db
	return repo
}
