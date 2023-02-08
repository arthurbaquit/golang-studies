package app_test

import (
	"errors"
	"testing"

	"github.com/arthurbaquit/hexagonal-arch/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product1 := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	product2 := app.Product{
		Id:     "2",
		Status: app.ProductStatusDisabled,
		Name:   "Product 2",
		Price:  0,
	}
	product3 := app.Product{
		Id:     "3",
		Status: app.ProductStatusDisabled,
		Name:   "Product 3",
		Price:  -1000,
	}
	product4 := app.Product{
		Id:     "4",
		Status: app.ProductStatusEnabled,
		Name:   "Product 4",
		Price:  10,
	}
	err1 := product1.Enable()
	err2 := product2.Enable()
	err3 := product3.Enable()
	err4 := product4.Enable()

	require.Nil(t, err1)
	require.NotNil(t, err2)
	require.NotNil(t, err3)
	require.Equal(t, errors.New("Product is already enabled"), err4)
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Name = ""
	_, err = product.IsValid()
	require.Equal(t, errors.New("Product name is required"), err)

	product.Name = "Product 1"
	product.Price = 0
	_, err = product.IsValid()
	require.Equal(t, errors.New("Product price must be greater than zero"), err)

	product.Price = 1000
	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, errors.New("Product status should be enable or disabled"), err)
}

func TestProduct_Disable(t *testing.T) {
	product1 := app.Product{
		Id:     "1",
		Status: app.ProductStatusEnabled,
		Name:   "Product 1",
		Price:  1000,
	}
	product2 := app.Product{
		Id:     "2",
		Status: app.ProductStatusDisabled,
		Name:   "Product 2",
		Price:  0,
	}
	err1 := product1.Disable()
	err2 := product2.Disable()

	require.Nil(t, err1)
	require.Equal(t, errors.New("Product is already disabled"), err2)
}

func TestProduct_GetId(t *testing.T) {
	product := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	require.Equal(t, "1", product.GetId())
}

func TestProduct_GetStatus(t *testing.T) {
	product := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	require.Equal(t, app.ProductStatusDisabled, product.GetStatus())
}

func TestProduct_GetName(t *testing.T) {
	product := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	require.Equal(t, "Product 1", product.GetName())
}

func TestProduct_GetPrice(t *testing.T) {
	product := app.Product{
		Id:     "1",
		Status: app.ProductStatusDisabled,
		Name:   "Product 1",
		Price:  1000,
	}
	require.Equal(t, int64(1000), product.GetPrice())
}
