package dataservice

import (
	"fmt"

	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DataServiceCart struct {
}

func NewDataServiceCart() IDataService[models.Cart] {
	return &DataServiceCart{}
}

func (dataService *DataServiceCart) GetAll(ctx *fiber.Ctx, db *gorm.DB) (dataReturn *[]models.Cart, err error) {
	db.Find(&dataReturn)
	return

}

func (dataService *DataServiceCart) GetOne(ctx *fiber.Ctx, db *gorm.DB, id int) (data *models.Cart, err error) {
	db.Where("id = ?", id).Find(&data)
	return
}

func (dataService *DataServiceCart) Add(ctx *fiber.Ctx, db *gorm.DB, data *models.Cart) error {
	db.Create(&data)
	return nil
}

func (dataService *DataServiceCart) Update(ctx *fiber.Ctx, db *gorm.DB, data *models.Cart) error {
	db.Updates(&data)
	return nil
}

func (dataService *DataServiceCart) Delete(ctx *fiber.Ctx, db *gorm.DB, id int) error {
	db.Where("id = ?", id).Delete(&models.Cart{})
	return nil
}

func (dataService *DataServiceCart) GetAllByUser(ctx *fiber.Ctx, db *gorm.DB, id int) (dataReturn []models.Product, err error) {

	rows, err := db.Raw("SELECT p.id,p.name,p.description FROM carts AS c INNER JOIN products AS p ON c.product_fk = p.id WHERE c.user_fk = 1").Rows()
	fmt.Println(err)
	defer rows.Close()
	var tempProd models.Product

	for rows.Next() {
		rows.Scan(&tempProd.ID, &tempProd.Name, &tempProd.Description)
		fmt.Println(tempProd)
		dataReturn = append(dataReturn, tempProd)
	}

	return
}
