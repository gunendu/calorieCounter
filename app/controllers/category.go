package controllers

import (
	"github.com/revel/revel"
	"github.com/gunendu/calorieCounter/app/models"
	"fmt"
)

type Category struct {
	Application
}

type Response struct {
	category Category
	nutrition Nutrition
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
	res, err  :=  c.Txn.Select(models.Nutrition{},`select * From Nutrition`)
	if err != nil {
		panic(err)
	}	
	var nutritions  []*models.Nutrition
	for _,r  := range(res) {
		b := r.(*models.Nutrition)
		nutritions = append(nutritions,b)
	}
	fmt.Print(nutritions)
	return c.RenderJSON(nutritions)
}




