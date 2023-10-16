package db

import (
	"database/sql"
	"fmt"

	app "github.com/arthurbaquit/hexagonal-arch/app"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	fmt.Println(db)
	return &ProductDB{db}
}

func (pdb *ProductDB) Get(id string) (app.ProductInterface, error) {
	var product app.Product
	stmt, err := pdb.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	return &product, err
}

func (pdb *ProductDB) create(product app.ProductInterface) error {
	stmt, err := pdb.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	return err
}

func (pdb *ProductDB) update(product app.ProductInterface) error {
	stmt, err := pdb.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	return err
}

func (pdb *ProductDB) Save(product app.ProductInterface) (app.ProductInterface, error) {
	_, err := pdb.Get(product.GetId())
	if err != nil {
		if err == sql.ErrNoRows {
			err = pdb.create(product)
		}
	} else {
		err = pdb.update(product)
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}

// type ProductReaderInterface interface {
// 	Get(id string) (ProductInterface, error)
// }

// type ProductWriterInterface interface {
// 	Save(product ProductInterface) (ProductInterface, error)
// }
