package controllers

import (
	"github.com/revel/revel"
	"github.com/gunendu/calorieCounter/app/models"
	_"fmt"
)

type Category struct {
	Application
}

func (c Category) Index() revel.Result {
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
	return c.RenderJSON(categories)
}

func(c Category) Foods() revel.Result {
	res, err  :=  c.Txn.Select(models.Nutrition{},`select  *  From Nutrition`)
	if err != nil {
		panic(err)
	}
	var nutritions  []*models.Nutrition
	for _,r  := range(res) {
		b := r.(*models.Nutrition)
		nutritions = append(nutritions,b)
	}
	return c.RenderJSON(nutritions)
}

func(c Category) CategoryById(id int) revel.Result {
	category := models.Category{}
	err := c.Txn.SelectOne(&category,`Select  *  From Category Where CategoryId = ? limit 1`,id)
	if err!=nil {
		panic(err)
	}
	return c.RenderJSON(category)
}
