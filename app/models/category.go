package models

import (
)

type Category struct {
	Id 		int64  		`db: "primarykey, autoincrement" `
	Name  string 	`db:",size:255"`
	UpdatedAt  int64
	CreatedAt  int64
}

