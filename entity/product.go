package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 5 {
		status = "Stock Hampir Habis"
	} else {
		status = ""
	}
	return status
}