package crud

import "CarwlerWeb/model"

func CrudFind(name string)(model.ProvinceData){
	var  data model.ProvinceData
	model.Db.Where(model.ProvinceData{ProvinceName: name}).First(&data)
	return data
}

func CrudFindProvinceId(procinceid int)([]model.VirusData){
	var  data []model.VirusData
	model.Db.Where(model.VirusData{ProvinceDataID:procinceid}).Find(&data)
	return data
}