package shop

import (
	"github.com/tozastation/architecture-example/internal/mvc/model/support"
	"github.com/tozastation/architecture-example/pkg/rdb/mysqldriver"
)

type (
	ShoppingCart  struct{
		rdb *mysqldriver.DriverWithGorm
	}
	IShoppingCart interface {
		SearchProduct(option support.SearchOption, productID uint) ([]*Product, error)
		AddProduct(item Product) error
		Checkout() error
	}
)
