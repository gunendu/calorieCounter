package controllers

import (

	"github.com/revel/revel"
	"io/ioutil"
	"encoding/json"
	"github.com/gunendu/calorieCounter/app/models"
	"fmt"
)

type PreperationsCtrl struct{
	Application
}

type IngredientsTemp struct{
	PreperationsId 	 int
	NutritionIds	 [] int
}

func(c PreperationsCtrl) Save() revel.Result {
	req,err := ioutil.ReadAll(c.Request.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Print(string(req))
	preperation := models.Preperations{}
	err = json.Unmarshal([]byte(req),&preperation)
	if err!=nil {
		panic(err)
	}
	if err := Dbm.Insert(&preperation); err!=nil {
		panic(err)
	}
	return c.RenderJSON(preperation)
}

func(c PreperationsCtrl) SaveIngredients() revel.Result {
	req,err := ioutil.ReadAll(c.Request.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Print(string(req))
	ingredients := IngredientsTemp{}
	err = json.Unmarshal([]byte(req),&ingredients)
	if err!= nil {
		panic(err)
	}
	fmt.Print(ingredients)
	prepId := int64(ingredients.PreperationsId)
	for _,r  := range(ingredients.NutritionIds) {
		ingredient := models.Ingredients{prepId,int64(r)}
		//fmt.Printf(ingredient))
		if err := Dbm.Insert(&ingredient);err!=nil {
			panic(err)
		}
	}
	return c.RenderJSON(nil)
}


