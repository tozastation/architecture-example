package customer

import (
	"github.com/tozastation/architecture-example/internal/mvc/model/shop"
	"github.com/tozastation/architecture-example/internal/mvc/model/support"
)

type (
	Customer struct {}
	ICustomer interface {
		SearchProduct(option support.SearchOption, productID uint) ([]*shop.Product, error)
		Checkout() error
		UpdateCustomerInfo() error
	}
)

func NewCustomer() (ICustomer, error) {
	n := &Customer{}
	return n, nil
}