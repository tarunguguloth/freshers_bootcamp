package main
import (
	"fmt"
	"freshers_bootcamp/Day3/Config"
	"freshers_bootcamp/Day3/Models"
	"freshers_bootcamp/Day3/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main(){
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}