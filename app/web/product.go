package web

import (
	"strconv"

	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type ProductRouter struct {
	dataserviceProd dataservice.IDataService[models.Product]
	session         *session.Store
	db              *gorm.DB
}

func NewProductRouter(dataService dataservice.IDataService[models.Product], db *gorm.DB, session *session.Store) *ProductRouter {
	return &ProductRouter{dataserviceProd: dataService, db: db, session: session}
}

func (serv *ProductRouter) MountRouterProduct(app *fiber.App) {

	app.Get("/create-product", serv.createProduct)
	app.Post("/create-product", serv.postCreateProduct)
	app.Get("/edit-product/:id", serv.editProduct)
	app.Post("/edit-product", serv.postEditProduct)
	app.Get("/delete-product/:id", serv.deleteProduct)

}

func (serv *ProductRouter) createProduct(c *fiber.Ctx) error {

	return c.Render("form_create", nil)

}

func (serv *ProductRouter) postCreateProduct(c *fiber.Ctx) error {
	sess, err := serv.session.Get(c)
	if err != nil {
		return err
	}

	id := sess.Get("id")

	prod := models.Product{
		User_Fk: id.(uint),
	}
	if err := c.BodyParser(&prod); err != nil {
		return err
	}

	if err := serv.dataserviceProd.Add(c, serv.db, &prod); err != nil {
		return err
	}

	return c.Redirect("/")

}

func (serv *ProductRouter) editProduct(c *fiber.Ctx) error {
	sess, _ := serv.session.Get(c)
	idUser := sess.Get("id")
	nameUser := sess.Get("name")

	if idUser == nil || nameUser == nil {
		return fiber.ErrUnauthorized
	}
	id := c.Params("id")
	idString, _ := strconv.Atoi(id)
	dataProd, _ := serv.dataserviceProd.GetOne(c, serv.db, idString)
	return c.Render("form_update", struct {
		User models.User
		Prod *models.Product
	}{
		User: models.User{Name: nameUser.(string)},
		Prod: dataProd,
	})

}

func (serv *ProductRouter) postEditProduct(c *fiber.Ctx) error {
	prod := models.Product{}
	if err := c.BodyParser(&prod); err != nil {
		return err
	}

	serv.dataserviceProd.Update(c, serv.db, &prod)

	return c.Redirect("/")
}

func (serv *ProductRouter) deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idint, _ := strconv.Atoi(id)

	serv.dataserviceProd.Delete(c, serv.db, idint)

	return c.Redirect("/")
}
