package controllers

import (
	"github.com/revel/revel"
	"github.com/gunendu/calorieCounter/app/models"
)

type Category struct {
	Application
}

func (c Category) Index() revel.Result {
	revel.INFO.Printf("get all categories")
	results, err := c.Txn.Select(models.Category{},
		`select  *  from Category`)
	if err != nil {
		panic("query fail")
	}
	var categories []*models.Category
	for _,r := range results {
		c := r.(*models.Category)
		categories = append(categories,c)
	}
	revel.INFO.Printf("array of categories %v",categories[0])
	return c.Render(categories)
}




