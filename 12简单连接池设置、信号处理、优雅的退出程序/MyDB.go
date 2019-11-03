package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var DBHelper *gorm.DB
var err error
func init()  {

	DBHelper,err= gorm.Open("mysql", "shenyi:123123@/gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err)
	}
	DBHelper.DB().SetMaxIdleConns(10)

	DBHelper.DB().SetMaxOpenConns(100)

	DBHelper.DB().SetConnMaxLifetime(time.Hour)


}


