package conndatabase

import (
	"crawler/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SaveData(results interface{}) {
	Results:=results.(model.ProvinceData)
	model.Db.AutoMigrate(&model.VirusData{})
	model.Db.AutoMigrate(&model.ProvinceData{})
	//city:=Results.ProvinceVirus[1:]
	if Results.ProvinceVirus[0].Name=="中国"{
		bigProvince:=Results.ProvinceVirus[0]
		province:=Results.ProvinceVirus[1]
		CitysDelete:=Results.ProvinceVirus[2:]

		var citys []model.VirusData
		for _,city:=range CitysDelete {
			if city.Name!="境外输入"&&city.Name!="地区待确认"{
				citys=append(citys,city)
			}
		}
		var nilvalues []model.VirusData
		SaveProvince(bigProvince,model.Db,nilvalues)
		SaveProvince(province,model.Db,citys)
	}else {
		province:=Results.ProvinceVirus[0]
		//赋值
		CitysDelete:=Results.ProvinceVirus[1:]
		var citys []model.VirusData
		for _,city:=range CitysDelete {
			if city.Name!="境外输入"&&city.Name!="地区待确认"{
				citys=append(citys,city)
			}
		}
		//创建记录
		SaveProvince(province,model.Db,citys)
	}

}
//赋值
func GiveValue(data model.VirusData,citys []model.VirusData)model.ProvinceData{
	provinceData:=model.ProvinceData{
		ProvinceName:  data.Name,
		Todayconfirm:  data.Todayconfirm,
		Confirm:       data.Confirm,
		Heal:          data.Heal,
		Dead:          data.Dead,
		ProvinceVirus: citys,
	}
	return provinceData
}
//将数据保存到数据库
func SaveProvince(Province model.VirusData,db *gorm.DB,citys []model.VirusData){
	var beingProvince model.ProvinceData
	var city []model.VirusData
	bigProvinceData:=GiveValue(Province,citys)
	////查询数据
	db.Where(model.ProvinceData{ProvinceName:bigProvinceData.ProvinceName}).First(&beingProvince)
	//没有就创建
	if beingProvince.ProvinceName!=Province.Name{
		db.Create(&bigProvinceData)
	}else {
		db.Where(model.VirusData{ProvinceDataID:beingProvince.Id}).Delete(&city)
		//gorm创建主表会将外键也一并创建了，但删除更新主表，外键表不会跟着删除更新
		db.Model(&beingProvince).Update(bigProvinceData)
	}
}
