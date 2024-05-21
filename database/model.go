package database

import "database/sql"

type User struct {
	Id       int    `db:"id"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Id_Addr  sql.NullInt32    `db:"id_addr"`
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
	Id_Addr  Addres_User
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

type Product struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	Discription string `db:"discription"`
	Category    int    `db:"category"`
}