package app_test

import (
	"testing"

	"github.com/arthurbaquit/hexagonal-arch/app"
	mock_app "github.com/arthurbaquit/hexagonal-arch/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{ProductPersistence: persistence}
	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{ProductPersistence: persistence}
	result, err := service.CreateProduct("abc", 1000)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisableProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{ProductPersistence: persistence}
	result, err := service.EnableProduct(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
	result, err = service.DisableProduct(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
