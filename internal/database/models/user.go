package models

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
	Addres   Addres_User
	Role     string
	IpAddres sql.NullString
}
