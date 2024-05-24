package database

import "database/sql"

type User struct {
	Id       int            `db:"id"`
	Phone    string         `db:"phone"`
	Password string         `db:"password"`
	Name     string         `db:"name"`
	Surname  string         `db:"surname"`
	Id_Addr  sql.NullInt32  `db:"id_addr"`
	Role     string         `db:"role"`
	IpAddres sql.NullString `db:"ip_addres"`
}

type Addres_User struct {
	Id       int    `db:"id"`
	City     string `db:"city"`
	Street   string `db:"street"`
	House    string `db:"house"`
	Flat     string `db:"flat"`
	Entrance string `db:"entrance"`
}

type User_Addr struct {
	Id       int
	Phone    string
	Password string
	Name     string
	Surname  string
	Addres  Addres_User
	Role     string
	IpAddres sql.NullString
}

type Main_Category struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Category struct {
	Id            int    `db:"id"`
	Name          string `db:"name"`
	Main_Category int    `db:"main_category"`
}

type Main_CategoryAndCategory struct {
	Id   int
	Name string
	Categories []Category
}

type Product struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	Description string `db:"description"`
	Category    int    `db:"category"`
}

type Parameters struct {
	Id int  	`db:"id"`
	Name string  	`db:"name"`
	Value string  	`db:"value"`
	Id_Product int   	`db:"id_product"`
}

type Product_Parameters struct {
	Id          int
	Name        string
	Price       int
	Description string
	Category    int
	Parameters []Parameters
}

type IPResponse struct {
	IP string `json:"ip"`
}
