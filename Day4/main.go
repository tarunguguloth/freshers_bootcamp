package main

import (
	"fmt"
	"freshers_bootcamp/Day4/Config"
	"freshers_bootcamp/Day4/Models"
	"freshers_bootcamp/Day4/Routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{},&Models.Customer{},&Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
