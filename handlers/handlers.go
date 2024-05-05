package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// Listar los contactos
func ListConstats(db *sql.DB) {
	query := "SELECT * FROM contact"

	//Ejecutamos la consulta
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//Si esto es un exito iterar
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("-----------------------------------------------------------")

	//Recorremos cada registro
	for rows.Next() {
		contact := models.Contact{}
		//Manejamos los valores nulos
		var valueEmail sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electronico"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-----------------------------------------------------------")
	}
}

func GetContactById(db *sql.DB, id int) {
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, id)

	contact := models.Contact{}
	var valueEmail sql.NullString

	//Escaneamos el resultado
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

	if err != nil {
		//Validamos si la persona ingreso un id incorrecto
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontro ningun contacto con el id %d", id)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electronico"
	}

	fmt.Println("\nCONTACTO:")
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("-----------------------------------------------------------")
}

func CreateContac(db *sql.DB, contact models.Contact) {

	query := "INSERT INTO contact (name, email, phone) VALUES (?,?,?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado exitosamente")
}

func UpdateContac(db *sql.DB, contact models.Contact) {

	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("El contacto se ha actualizado correctamente")
}

func DeleteContac(db *sql.DB, id int) {

	query := "DELETE FROM contact WHERE id = ?"

	_, err := db.Exec(query, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Se ha eliminado correctamente")
}
