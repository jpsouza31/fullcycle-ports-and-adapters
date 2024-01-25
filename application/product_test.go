package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"ports-and-adapters/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "any_name"
	product.Status = application.DISABLED
	product.Price = 10

	sut := product.Enable()
	require.Nil(t, sut)

	product.Price = 0
	sut = product.Enable()
	require.Equal(t, "the price must be grater than zero", sut.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "any_name"
	product.Status = application.ENABLED
	product.Price = 10

	sut := product.Disable()
	require.Equal(t, "the price must be zero", sut.Error())

	product.Price = 0
	sut = product.Disable()
	require.Nil(t, sut)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
