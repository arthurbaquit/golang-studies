package db

import (
	"database/sql"

	app "github.com/arthurbaquit/hexagonal-arch/app"
)

type ProductDB struct {
	db *sql.DB
}

func (pdb *ProductDB) Get (id string) (app.ProductInterface, error){
	var product app.ProductInterface
	stmt, err := pdb.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	return nil, nil
}

func (pdb *ProductDB) Save (product app.ProductInterface) (app.ProductInterface, error){
	return nil, nil
}

// type ProductReaderInterface interface {
// 	Get(id string) (ProductInterface, error)
// }

// type ProductWriterInterface interface {
// 	Save(product ProductInterface) (ProductInterface, error)
// }