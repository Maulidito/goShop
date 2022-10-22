package dataservice

import (
	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceTransaction struct {
}

func NewDataServiceTransaction() IDataService[models.Transaction] {
	return &DataServiceTransaction{}
}

func (dataService *DataServiceTransaction) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.Transaction, err error) {
	db.Find(&dataReturn)
	return
}

func (dataService *DataServiceTransaction) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.Transaction, err error) {
	db.Where("id = ?", id).Find(&data)
	return
}

func (dataService *DataServiceTransaction) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.Transaction) error {
	db.Create(&data)
	return nil
}

func (dataService *DataServiceTransaction) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.Transaction) error {
	db.Updates(&data)
	return nil
}

func (dataService *DataServiceTransaction) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.Transaction{})
	return nil
}
