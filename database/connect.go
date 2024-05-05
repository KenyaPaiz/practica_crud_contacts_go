package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	//conexion de la base de datos
	//Abrimos la conexion de la base de datos ofreciendo su info

	//cargando variables de entorno
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	//Formateamos la cadena y llamamos las variables mediante os.Getenv
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	//Verificamos si mantenemos la conexion
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Conexion a la base de datos mysql exitosa")
	return db, nil
}
