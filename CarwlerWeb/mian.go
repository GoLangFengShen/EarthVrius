package main

import (
	"CarwlerWeb/model"
	"CarwlerWeb/routing"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){

	err:=model.InitMysql()
	if err!=nil {
		panic(err)
	}
	defer model.Db.Close()
	r:=routing.Routing()



	r.Run(":9080")
}
