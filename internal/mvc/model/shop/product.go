package shop

import (
	"github.com/jinzhu/gorm"
	"github.com/tozastation/architecture-example/internal/mvc/model/support"
)

type (
	ProductOnJson struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}
	Product struct {
		ID    uint
		Name  string
		Price support.Price
	}
	ProductOnDB struct {
		gorm.Model
		Name  string
		Price uint
	}
)
