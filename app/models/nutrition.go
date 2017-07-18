package models

import (

)

type  Nutrition  struct {
	 Id 				 	 int64	 `db:"Id, primarykey, autoincrement"`
	 CategoryId    int
	 Name 		   string
	 Energy        float64
	 Moisture    float64
	 Proteins     float64
	 Fat               float64
	 Minerals    float64
	 Fibre			 float64
	 Carbos       float64
	 Calcium     float64
	 Phosphorous   float64
	 Iron					  float64
}