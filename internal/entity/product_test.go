package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Name, "Product 1")
	assert.Equal(t, product.Price, 10.99)
	assert.NotEmpty(t, product.ID)
	assert.NotZero(t, product.CreatedAt)
}

func TestProduct_Validate(t *testing.T) {
	product, err := NewProduct("Product 1", 10.99)
    assert.Nil(t, err)
    assert.Nil(t, product.Validate())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.99)
    assert.Nil(t, product)
    assert.Error(t, err)
    assert.Equal(t, err, ErrorNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
    assert.Nil(t, product)
    assert.Error(t, err)
    assert.Equal(t, err, ErrorPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -10.99)
    assert.Nil(t, product)
    assert.Error(t, err)
    assert.Equal(t, err, ErrorInvalidPrice)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product 1", 10.99)
	assert.Nil(t, product.Validate())
    assert.Nil(t, err)
	assert.NotNil(t, product)
}
