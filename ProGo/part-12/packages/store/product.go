// Пакет store описывает методы и типы
// обычно требуемые для интернет магазина
package store

var standardTax = newTaxRate(0.25, 20)

// Product описывает товары для продажи
type Product struct {
  Name, Category string // Имя и тип продукта
  price float64
}

func NewProduct(name, category string, price float64) *Product {
  return &Product{ name, category, price }
}

func (p *Product) Price() float64 {
  return standardTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
  p.price = newPrice
}
