package models

import (

	"github.com/go-gorp/gorp"
	"time"
)

type Preperations struct{
	 Id 		int64	 `db:"Id, primarykey, autoincrement"`
	 Name	string
	 UpdatedAt  int64
	 CreatedAt  int64	  	  
}

func(p *Preperations) PreInsert(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}

func(p *Preperations) PreUpdate(s gorp.SqlExecutor) error {
	p.CreatedAt = time.Now().UnixNano()
	p.UpdatedAt = p.CreatedAt
	return nil
}