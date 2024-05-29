package models

type Product struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	Description string `db:"description"`
	Category    int    `db:"category"`
}

type Parameters struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Value      string `db:"value"`
	Id_Product int    `db:"id_product"`
}

type Product_Parameters struct {
	Id          int
	Name        string
	Price       int
	Description string
	Category    int
	Parameters  []Parameters
}
