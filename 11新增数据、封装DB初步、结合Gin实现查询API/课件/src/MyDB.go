package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DBHelper *gorm.DB
var err error
func init()  {
	DBHelper,err= gorm.Open("mysql", "root:123123@/gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err)
	}
	DBHelper.LogMode(true)
}
