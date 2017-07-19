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


