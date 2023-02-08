package app

import (
	"errors"

	"github.com/gofrs/uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetStatus() string
	GetName() string
	GetPrice() int64 // in decimals => 1000 = 10.00
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	CreateProduct(name string, price int64) (ProductInterface, error)
	EnableProduct(Product ProductInterface) (ProductInterface, error)
	DisableProduct(Product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

const (
	ProductStatusEnabled  = "enabled"
	ProductStatusDisabled = "disabled"
)

type Product struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
}

func newProduct() *Product {
	return &Product{Id: uuid.Must(uuid.NewV7()).String(), Status: ProductStatusDisabled}
}

func (p *Product) IsValid() (bool, error) {
	if p.Name == "" {
		return false, errors.New("Product name is required")
	}
	if p.Price <= 0 {
		return false, errors.New("Product price must be greater than zero")
	}
	if p.Status != ProductStatusEnabled && p.Status != ProductStatusDisabled {
		return false, errors.New("Product status should be enable or disabled")
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Status == ProductStatusEnabled {
		return errors.New("Product is already enabled")
	}
	if p.Price <= 0 {
		return errors.New("Product price must be greater than zero to be enabled")
	}
	p.Status = ProductStatusEnabled
	return nil
}

func (p *Product) Disable() error {
	if p.Status == ProductStatusDisabled {
		return errors.New("Product is already disabled")
	}
	p.Status = ProductStatusDisabled
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() int64 {
	return p.Price
}
