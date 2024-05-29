package models

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
	Id         int
	Name       string
	Categories []Category
}
