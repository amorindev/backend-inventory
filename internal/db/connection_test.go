package db

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importar el controlador de PostgreSQL
)

// test DB conection
func TestDBConnection(t *testing.T){
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatal(err)
	}
	DBConnection()
	err = DB.Ping()
	if err != nil {
	  t.Fatalf("Ping to DB: %v", err)
	}
}
