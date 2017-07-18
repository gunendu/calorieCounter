package controllers

import (

	"strings"
	"fmt"
	"time"	
	"database/sql"
    
	"github.com/go-gorp/gorp"	
	 _ "github.com/go-sql-driver/mysql"
    
	r "github.com/revel/revel"
	
    "github.com/gunendu/calorieCounter/app/models"
	"github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
	createdAt int64
)

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == ""{
			revel.ERROR.Fatal("could not find paramter " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host  :=  getParamString("db.host",  "")
	port := getParamString("db,.port", "3306")
	user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "auction")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", 
        user, pass, protocol, host, port, dbname, dbargs)

}

var  InitDB  func() = func()  {

	connectionString := getConnectionString()
	
	if db,  err  :=  sql.Open("mysql", connectionString);  err  != nil {
		 revel.ERROR.Fatal(err)
	}  else {
		Dbm = &gorp.DbMap{
            Db: db, 
            Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 20,
		"Name":     100,
	})

	t = Dbm.AddTable(models.Hotel{}).SetKeys(true, "HotelId")
	setColumnSizes(t, map[string]int{
		"Name":    50,
		"Address": 100,
		"City":    40,
		"State":   6,
		"Zip":     6,
		"Country": 40,
	})

	t = Dbm.AddTable(models.Booking{}).SetKeys(true, "BookingId")
	t.ColMap("User").Transient = true
	t.ColMap("Hotel").Transient = true
	t.ColMap("CheckInDate").Transient = true
	t.ColMap("CheckOutDate").Transient = true
	setColumnSizes(t, map[string]int{
		"CardNumber": 16,
		"NameOnCard": 50,
	})

	Dbm.AddTable(models.Category{}).SetKeys(true, "CategoryId")

	Dbm.AddTable(models.Nutrition{})
	
	Dbm.TraceOn("[gorp]", r.INFO)
		
	Dbm.CreateTablesIfNotExists()
	
	createdAt  = time.Now().UnixNano() / int64(time.Millisecond)
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
