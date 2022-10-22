package dataservice

import (
	"testing"

	"github.com/Maulidito/tugasday5/app/database"
	"github.com/Maulidito/tugasday5/dataservice"
	"github.com/Maulidito/tugasday5/models"
	"github.com/gofiber/fiber/v2"
)

func TestCreateUser(t *testing.T) {

	data := models.User{
		Name:     "dito",
		Password: "dty",
		Email:    "dfasd",
	}

	db, err := database.NewDatabasePostgres()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	service := dataservice.NewDataServiceUser()

	err = service.Add(&fiber.Ctx{}, db, &data)

	if err != nil {
		t.Fail()
		return
	}
	t.Log(data)

}

func TestDeleteUser(t *testing.T) {

	id := 1

	db, err := database.NewDatabasePostgres()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	service := dataservice.NewDataServiceUser()

	err = service.Delete(&fiber.Ctx{}, db, id)

	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

}
