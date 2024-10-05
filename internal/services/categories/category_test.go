package categories

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq" // Importar el controlador de PostgreSQL

	database "github.com/amorindev/backend-inventory/internal/db"
)


func TestGetCategoriesDto(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
	  t.Fatalf("An error %v was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	database.DB = db

	// Simulamos las filas que devuelve las columas
	rows := sqlmock.NewRows([]string{"cat_id", "cat_name"}).AddRow(1,"TECNOLOGÍA").AddRow(7,"ELECTRODOMÉSTICOS")

	//simular la consulta
	mock.ExpectQuery("SELECT cat_id, cat_name FROM category").WillReturnRows(rows)

	categories, err := GetCategoriesDto()
	if err != nil {
		t.Errorf("error was not expected while getting categories: %s", err)
	}

	// Definimos lo que esperamos obtener
	expected := []Category{
		{CatID: 1, CatName: "TECNOLOGÍA"},
		{CatID: 7, CatName: "ELECTRODOMÉSTICOS"},
	}

	// VERIFICAMOS QUE EL RESULTADO SEA ESPERADO
	if !reflect.DeepEqual(categories, expected) {
		t.Errorf("expected %v, got %v", expected, categories)
	}

	// VERFICAMOS QUE SE HAYAN CUMPLIDO TODAS LAS EXPECTATIVAS
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}



