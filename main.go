package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Establecer la conexion a la bd
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Llamamos al controlador
	handlers.ListConstats(db)

	handlers.GetContactById(db, 2)

	//Creamos una nueva instancia de Contact
	newContact := models.Contact{
		Name:  "Kenia Paiz",
		Email: "kenia@gmail.com",
		Phone: "5467-9876",
	}

	handlers.CreateContac(db, newContact)

	//Actualizando un contacto
	updateContact := models.Contact{
		Id:    3,
		Name:  "Sebastian Carlos Lopez",
		Email: "sebas123@gmail.com",
		Phone: "2345-7865",
	}

	handlers.UpdateContac(db, updateContact)

	handlers.DeleteContac(db, 5)
}
