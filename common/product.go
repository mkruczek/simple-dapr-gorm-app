package common

import (
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) String() string {
	return fmt.Sprintf("Product{ID: %d, Code: %s, Price: %d}", p.ID, p.Code, p.Price)
}
