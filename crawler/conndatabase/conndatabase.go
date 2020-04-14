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
		SaveTitys(citys,model.Db)

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

		SaveTitys(citys,model.Db)
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
//将省和外国国家数据保存到数据库
func SaveProvince(Province model.VirusData,db *gorm.DB,citys []model.VirusData){
	var beingProvince model.ProvinceData
	bigProvinceData:=GiveValue(Province,citys)
	////查询数据
	db.Where(model.ProvinceData{ProvinceName:bigProvinceData.ProvinceName}).First(&beingProvince)
	////没有就创建
	if beingProvince.ProvinceName!=Province.Name{
		db.Create(&bigProvinceData)
	}else {
		db.Model(&beingProvince).Update(bigProvinceData)
	}
}
//保存个省所属城市数据
func SaveTitys(citys []model.VirusData,db *gorm.DB) {
	//
	//var beingTity model.VirusData
	//
	//	for _,city:=range citys{
	/*	sql:="REPLACE INTO `virus_data` (`Name`,`Todayconfirm`,`Confirm`,`Heal`,`Dead`) VALUES"
		for k,result:=range citys{
			//db.Where(model.VirusData{Name:result.Name}).First(&CheckCity)
			//if CheckCity.Name!=result.Name{
			if len(citys)-1==k{
				sql+=fmt.Sprintf("('%s','%s','%s','%s','%s');",result.Name,result.Todayconfirm,result.Confirm,result.Heal,result.Dead)
			}else {
				sql+=fmt.Sprintf("('%s','%s','%s','%s','%s'),",result.Name,result.Todayconfirm,result.Confirm,result.Heal,result.Dead)
			}*/

	//virusdata:=model.VirusData{
	//Name:city.Name,
	//Todayconfirm:city.Todayconfirm,
	//Confirm:city.Confirm,
	//Heal:city.Heal,
	//Dead:city.Dead,
	//}
	//db.Where(model.VirusData{Name:virusdata.Name}).First(&beingTity)
	//if beingTity.Name!=""{

	db.Model(model.VirusData{}).Updates(&citys)

}
//			//db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
//			db.Model(&beingTity).Update(virusdata)
//		}else {
//			db.Create(&virusdata)
//		}
//		}
//}
	/*db.Exec(sql)
}*/