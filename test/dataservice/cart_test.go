package dataservice

import (
	"fmt"
	"testing"

	"github.com/Maulidito/tugasday5/app/database"
	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/gofiber/fiber/v2"
)

func TestGetCartAndProduct(t *testing.T) {
	dataService := dataservice.DataServiceCart{}
	db, _ := database.NewDatabasePostgres()

	retur, _ := dataService.GetAllByUser(&fiber.Ctx{}, db, 1)

	fmt.Println(retur)

}
