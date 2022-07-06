package manager

import (
	"go_wmb_gin_refactor/config"
	"go_wmb_gin_refactor/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	env := os.Getenv("ENV")
	if env == "migration" {
		db.Debug()
		db.AutoMigrate(&model.Bill{}, &model.Discount{}, &model.Discount{}, &model.Menu{}, &model.Menu_Price{}, &model.Bill_Detail{}, &model.Customer{})

		//masukin table untuk dimigrate
	} else if env == "dev" {
		db.Debug()
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
