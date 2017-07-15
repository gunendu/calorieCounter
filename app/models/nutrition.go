package models

import (

)

type  Nutrition  struct {
	 Id 				 	 int64	 `db:"Id, primarykey, autoincrement"`
	 CategoryId    int
	 Name 		   string
	 Energy        int
	 Moisture    int
	 Proteins     int
	 Fat               int
	 Minerals    int
	 Fibre			 int
	 Carbos       int
	 Calcium     int
	 Phosphorous   int
	 Iron					  int
}