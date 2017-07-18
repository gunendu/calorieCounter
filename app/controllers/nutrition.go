package controllers

import (
	"encoding/json"
	"encoding/csv"
	"io"
	"os"
	"strings"
	"github.com/revel/revel"
	"fmt"
	"strconv"
	"github.com/gunendu/calorieCounter/app/models"
	_"github.com/go-restit/lzjson"
	"io/ioutil"
)

type NutritionCtrl struct {
	Application
}

var (
	categoryType   string
	temp 				    string
	name 					string
	categoryId          int
)

func (c NutritionCtrl) Index() revel.Result {
	results, err := c.Txn.Select(models.Nutrition{},`select  *  From Nutrition `)

	if err !=nil {
		panic(err)
	}
	var nutritions  []*models.Nutrition
	for _, r  := range(results) {
		b := r.(*models.Nutrition)
		nutritions = append(nutritions,b)
	}
	return c.Render(nutritions)
}

func (c NutritionCtrl) SaveData() revel.Result {
	file,  err  :=  os.Open("nutrition.csv")
	if err != nil {
			panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	lineCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
			return nil
		}
		switch temp = strings.TrimSpace(record[0]) ; temp {
		case "common foods" :
			categoryType = temp
			categoryId  = 5
		case "PULSES AND LEGUMES":
			categoryType = temp
			categoryId = 6
		case "Leafy Vegetables ":
			categoryType = temp
			categoryId = 7
		case "Roots and Tubers":
			categoryType= temp
			categoryId = 8
		case "Other Vegetables":
			categoryType=temp
			categoryId = 9
		case "Nuts and Oilseeds":
			categoryType=temp
			categoryId = 10
		case "Condiments and Spices":
			categoryType = temp
			categoryId = 11
		case "Fruits":
			categoryType = temp
			categoryId = 12
		case "Milk and Milk Products":
			categoryType = temp
			categoryId = 13
		case  "Fats and Edible Oils":
			categoryId = 14
		case  "Sugars":
			categoryId = 15
		case  "Fatty Acids":
			categoryId = 16
		}
		if len(record) > 2  && len(record[1])>1 {
		kcal,err := strconv.ParseFloat(record[1],64)
		i, err  :=  strconv.ParseFloat(record[2],64)
		j,err :=  strconv.ParseFloat(record[3],64)
		k,err := strconv.ParseFloat(record[4],64)
		l,err :=  strconv.ParseFloat(record[5],64)
		m,err := strconv.ParseFloat(record[6],64)
		n,err := strconv.ParseFloat(record[7],64)
		o,err := strconv.ParseFloat(record[8],64)
		p,err := strconv.ParseFloat(record[9],64)
		q,err := strconv.ParseFloat(record[10],64)

		if err != nil {
			panic(err)
		}
		//fmt.Printf("%d %s %d %d %d %d %d %d %d %d %d %d",categoryId,record[0],kcal,i,j,k,l,m,n,o,p,q)
		nutrition1 := []*models.Nutrition{
			&models.Nutrition{0,categoryId,record[0],kcal,i,j,k,l,m,n,o,p,q},
		}

		for _,  nut := range nutrition1 {
			if err := Dbm.Insert(nut); err != nil{
				panic(err)
			}
   		}

		fmt.Println(lineCount)
		lineCount++
	}
	}
	return nil
}

func (c NutritionCtrl) Foods() revel.Result {
	    b, err := ioutil.ReadAll(c.Request.Body)
		if err !=nil {
			panic(err)
		}
    	revel.INFO.Println(string(b))
		nutrition := models.Nutrition{}
		err1 := json.Unmarshal([]byte(b), &nutrition)

		if err1 != nil {
			panic(err1)
		}
		return c.RenderJSON(nutrition)
}
