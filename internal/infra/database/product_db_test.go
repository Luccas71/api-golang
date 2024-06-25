package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Luccas1/api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	p, _ := entity.NewProduct("Product 1", 10.00)

	productDB := NewProduct(db)

	err = productDB.Create(p)
	assert.NoError(t, err)
	assert.NotEmpty(t, p.ID)
}

func TestListProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	for i := 1; i < 25; i++ {
		p, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)

		//usando metodo nativo
		db.Create(p)
	}
	products, err := productDB.List(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDB.List(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDB.List(3, 10, "asc")
	assert.Len(t, products, 4)
	assert.NoError(t, err)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[3].Name, "Product 24")
}

func TestProductFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    if err!= nil {
        t.Error(err)
    }

    db.AutoMigrate(&entity.Product{})
    productDB := NewProduct(db)

    p, err := entity.NewProduct("Product 1", 10.00)
    assert.NoError(t, err)

    db.Create(p)

    product, err := productDB.FindByID(p.ID.String())
    assert.NoError(t, err)
    assert.NotNil(t, product)
    assert.Equal(t, product.Name, "Product 1")
    assert.Equal(t, product.Price, 10.00)
}
