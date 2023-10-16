package db

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	app "github.com/arthurbaquit/hexagonal-arch/app"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setup() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}
	createTableProduct(db)
	createMockProducts(db)
	fmt.Println("setup")
	return db
}

func createTableProduct(Db *sql.DB) {
	table := `CREATE TABLE IF NOT EXISTS products (
	id TEXT PRIMARY KEY,
	name TEXT,
	price INTEGER,
	status TEXT
)`
	stmt, err := Db.Prepare(table)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	stmt.Exec()
}

func createMockProducts(Db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("123", "test", 0, "disabled")`
	stmt, err := Db.Prepare(insert)
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	stmt.Exec()

}

func TestProductDb_Get(t *testing.T) {
	DB := setup()
	defer DB.Close()
	productDb := NewProductDB(DB)
	product, err := productDb.Get("123")
	require.Equal(t, nil, err)
	require.NotNil(t, product)
	require.Equal(t, "123", product.GetId())
	require.Equal(t, "test", product.GetName())
	require.Equal(t, int64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	DB := setup()
	defer DB.Close()
	productDb := NewProductDB(DB)
	fmt.Println(productDb)
	product, err := productDb.Save(&app.Product{
		Id:     "1234",
		Name:   "testing create",
		Price:  10,
		Status: "enabled",
	})
	require.Equal(t, nil, err)
	require.NotNil(t, product)
	require.Equal(t, "1234", product.GetId())
	require.Equal(t, "testing create", product.GetName())
	require.Equal(t, int64(10), product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())

	product, err = productDb.Save(&app.Product{
		Id:     "1234",
		Name:   "testing update",
		Price:  100,
		Status: "disabled",
	})
	require.Equal(t, nil, err)
	require.NotNil(t, product)
	require.Equal(t, "1234", product.GetId())
	require.Equal(t, "testing update", product.GetName())
	require.Equal(t, int64(100), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
