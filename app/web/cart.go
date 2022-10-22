package web

import (
	"strconv"

	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type CartRouter struct {
	dataservice *dataservice.DataServiceCart
	session     *session.Store
	db          *gorm.DB
}

func NewCartRouter(sess *session.Store, dataService dataservice.IDataService[models.Cart], db *gorm.DB) *CartRouter {
	return &CartRouter{session: sess, db: db, dataservice: dataService.(*dataservice.DataServiceCart)}
}

func (serv *CartRouter) MountCartRouter(app *fiber.App) {

	app.Get("/cart", serv.cart)
	app.Get("/add-cart/:id", serv.addCart)

}

func (serv *CartRouter) cart(c *fiber.Ctx) error {
	sess, _ := serv.session.Get(c)
	idUser := sess.Get("id")
	nameUser := sess.Get("name")

	if idUser == nil || nameUser == nil {
		return fiber.ErrUnauthorized
	}

	listCart, _ := serv.dataservice.GetAllByUser(c, serv.db, int(idUser.(uint)))
	return c.Render("mycart", struct {
		User models.User
		Cart []models.Product
	}{User: models.User{Name: nameUser.(string)}, Cart: listCart})
}

func (serv *CartRouter) addCart(c *fiber.Ctx) error {

	sess, _ := serv.session.Get(c)

	idUser := sess.Get("id")

	if idUser == nil {
		return fiber.ErrUnauthorized
	}

	id := c.Params("id")
	idproduct, _ := strconv.Atoi(id)
	serv.dataservice.Add(c, serv.db, &models.Cart{Product_Fk: uint(idproduct), User_Fk: idUser.(uint)})
	return c.Redirect("/")
}
