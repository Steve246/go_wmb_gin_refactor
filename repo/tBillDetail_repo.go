package repo

import (
	"errors"
	"go_wmb_gin_refactor/model"

	"gorm.io/gorm"
)

type BillDetailRepository interface {
	// Create(tBil *model.Bill_Detail) error
	FindById(id string) (model.Bill_Detail, error)
}

type billDetailRepository struct {
	db *gorm.DB
}

func (c *billDetailRepository) FindById(id string) (model.Bill_Detail, error) {
	var billDetail model.Bill_Detail
	result := c.db.Unscoped().First(&billDetail, "id = ?", id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return billDetail, nil
		} else {
			return billDetail, err
		}

	}

	return billDetail, nil
}

func NewBillDetailRepository(db *gorm.DB) BillDetailRepository {
	repo := new(billDetailRepository)
	repo.db = db
	return repo
}
