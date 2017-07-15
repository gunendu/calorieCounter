package controllers

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
	"github.com/revel/revel"	
	"fmt"
	"strconv"
	"github.com/gunendu/calorieCounter/app/models"
)

type Nutrition struct {
	Application
}

var (
	categoryType   string
	temp 				    string
	name 					string
	categoryId          int
)

func (c Nutrition) Index() revel.Result {
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
		fmt.Printf("%V",record)	
		kcal,err := strconv.Atoi(record[1])
		i, err  :=  strconv.Atoi(record[2])
		j,err :=  strconv.Atoi(record[3])
		k,err := strconv.Atoi(record[4])
		l,err :=  strconv.Atoi(record[5])
		m,err := strconv.Atoi(record[6])
		n,err := strconv.Atoi(record[7])
		o,err := strconv.Atoi(record[8])
		p,err := strconv.Atoi(record[9])
		q,err := strconv.Atoi(record[10])
	
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %d %d %d %d %d %d %d %d %d %d",categoryId,record[0],kcal,i,j,k,l,m,n,o,p,q)
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




