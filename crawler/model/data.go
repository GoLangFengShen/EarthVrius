package model

import "github.com/jinzhu/gorm"

type VirusData struct{
	Id int
	ProvinceDataID int
	Name string 		 `gorm:"UNIQUE_INDEX"`
	Todayconfirm string
	Confirm string
	Heal string
	Dead  string
}
type ProvinceData struct {
	Id int
	ProvinceName string `gorm:"UNIQUE_INDEX"`
	Todayconfirm string
	Confirm string
	Heal string
	Dead  string
	ProvinceVirus []VirusData `gorm:"foreignkey:ProvinceDataID"`
}
var Db *gorm.DB
func InitMysql()(err error){
	Db,err=gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/virus crawler?charset=utf8mb4")
	if err != nil {
		return
	}
	return Db.DB().Ping()
}
