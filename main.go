package main

import (
	"log"

	"github.com/Maulidito/tugasday5/app/database"
	"github.com/Maulidito/tugasday5/app/web"
	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/Maulidito/tugasday5/models"
	"github.com/Maulidito/tugasday5/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

func main() {

	db, err := database.NewDatabasePostgres()

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(models.User{}, models.Product{}, models.Transaction{}, models.Cart{})

	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	sess := session.New()

	cartDataService := dataservice.NewDataServiceCart()
	productDataService := dataservice.NewDataServiceProduct()
	// transactionDataService := dataservice.NewDataServiceTransaction()
	userDataService := dataservice.NewDataServiceUser()

	serviceAuth := service.NewServiceAuth(userDataService, db)

	homeRouter := web.NewHomeRouter(serviceAuth, sess, userDataService, productDataService, db)
	productRouter := web.NewProductRouter(productDataService, db, sess)
	cartRouter := web.NewCartRouter(sess, cartDataService, db)

	homeRouter.MountRouterHome(app)
	productRouter.MountRouterProduct(app)
	cartRouter.MountCartRouter(app)

	app.Listen(":8080")

}
