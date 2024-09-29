package models

type Product struct {
	ID    int64   `pg:"id,pk"`
	Name  string  `pg:"name"`
	Price float64 `pg:"price"`
}
