package web

import (
	"fmt"

	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/Maulidito/tugasday5/models"
	"github.com/Maulidito/tugasday5/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type HomeRouter struct {
	serviceAuth        service.IServiceAuth
	dataserviceUser    *dataservice.DataServiceUser
	dataserviceProduct dataservice.IDataService[models.Product]
	session            *session.Store
	db                 *gorm.DB
}

func NewHomeRouter(serviceAuth service.IServiceAuth, sess *session.Store, dataService dataservice.IDataService[models.User], dataServiceProd dataservice.IDataService[models.Product], db *gorm.DB) *HomeRouter {
	return &HomeRouter{serviceAuth: serviceAuth, session: sess, dataserviceUser: dataService.(*dataservice.DataServiceUser), db: db, dataserviceProduct: dataServiceProd}
}

func (serv *HomeRouter) MountRouterHome(app *fiber.App) {

	app.Get("/", serv.home)
	app.Get("/login", serv.login)
	app.Post("/login", serv.postLogin)
	app.Get("/registration", serv.registration)
	app.Post("/registration", serv.postRegistration)
	app.Get("/logout", serv.logout)

}

func (serv *HomeRouter) home(c *fiber.Ctx) error {
	sess, err := serv.session.Get(c)
	dataUser := &models.User{}

	if err != nil {
		return err
	}

	if sess.Get("name") != nil {

		dataUser, err = serv.dataserviceUser.GetOneByName(c, serv.db, sess.Get("name").(string))
		if err != nil {
			return err
		}
		sess.Set("id", dataUser.ID)
		sess.Set("name", dataUser.Name)

		sess.Save()

	}

	listProduct, err := serv.dataserviceProduct.GetAll(c, serv.db)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	fmt.Println(listProduct)

	return c.Render("home", struct {
		User    interface{}
		Product interface{}
	}{User: dataUser, Product: listProduct})

}

func (serv *HomeRouter) login(c *fiber.Ctx) error {

	return c.Render("login", nil)

}

func (serv *HomeRouter) postLogin(c *fiber.Ctx) error {

	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := serv.serviceAuth.Login(c, user.Name, user.Password); err != nil {
		return err
	}

	sess, err := serv.session.Get(c)

	if err != nil {
		return err
	}

	sess.Set("name", user.Name)

	if err := sess.Save(); err != nil {
		return err
	}

	return c.Redirect("/")

}

func (serv *HomeRouter) registration(c *fiber.Ctx) error {

	return c.Render("regis", nil)
}

func (serv *HomeRouter) postRegistration(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	serv.serviceAuth.Register(c, &user)

	return c.Redirect("/")
}

func (serv *HomeRouter) logout(c *fiber.Ctx) error {

	sess, err := serv.session.Get(c)

	if err != nil {
		return err
	}

	if err := sess.Destroy(); err != nil {
		return err
	}
	return c.Redirect("/")

}
