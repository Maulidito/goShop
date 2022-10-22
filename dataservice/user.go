package dataservice

import (
	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceUser struct {
}

func NewDataServiceUser() IDataService[models.User] {
	return &DataServiceUser{}
}

func (ds *DataServiceUser) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.User, err error) {
	db.Find(&dataReturn)
	return
}

func (ds *DataServiceUser) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.User, err error) {
	db.Where("id = ?", id).Find(&data)
	return
}

func (ds *DataServiceUser) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.User) error {

	db.Create(&data)
	return nil
}

func (ds *DataServiceUser) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.User) error {
	db.Updates(&data)
	return nil
}

func (ds *DataServiceUser) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.User{})
	return nil
}

func (ds *DataServiceUser) GetOneByName(ctx *fiber.Ctx, db *gorm.DB, name string) (data *models.User, err error) {
	db.Where("name = ?", name).First(&data)
	return
}
