package contronller

import "C"
import (
	"CarwlerWeb/contronller/crud"
	"CarwlerWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


//earth.html主要疫情城市黄色圆点大小代表其严重程度
func Earth(c *gin.Context) {
	names:=[]string{"美国","西班牙","意大利","德国","法国","中国","伊朗","英国","土耳其"}
	h:=ProvinceFind(names)

	c.HTML(http.StatusOK,"earth.html",h)
}
//点击中国跳转，点击外国直接给数据
func EarthData(c *gin.Context){
	name:=c.Param("name")
	if name!="中国"{
	var s model.ProvinceData
	model.Db.Where(model.ProvinceData{ProvinceName:name}).First(&s)

	names:=[]string{"美国","西班牙","意大利","德国","法国","中国","伊朗","英国","土耳其"}
		h:=ProvinceFind(names)

		h["name"]=s.ProvinceName
		h["todayconfirm"]=s.Todayconfirm
		h["confirm"]=s.Confirm
		h["heal"]=s.Heal
		h["dead"]=s.Dead

	c.HTML(http.StatusOK,"earth.html",h)
	}else {
		c.Request.URL.Path="/china"
		model.R.HandleContext(c)
	}

}
//各省确诊数
func EarthChina(c *gin.Context){
	provinces:=[]string{"黑龙江","吉林","辽宁","内蒙古","北京","天津","河北","山东","山西","陕西","甘肃","新疆","青海","西藏","四川","河南","江苏","安徽","湖北","重庆","贵州","湖南","江西","上海","浙江","福建","云南","广西","广东","海南","香港","澳门","台湾","宁夏"}
	h:=ProvinceFind(provinces)

	c.HTML(http.StatusOK,"china.html",h)
}
//China2.html需要两部分数据
type Parameter struct {
	ParameterProvince map[string]interface{}
	ParameterCity map[string]interface{}
}
//各省确诊数+下属城市疫情数据
func ChinaCity(c *gin.Context){
	provinces:=[]string{"黑龙江","吉林","辽宁","内蒙古","北京","天津","河北","山东","山西","陕西","甘肃","新疆","青海","西藏","四川","河南","江苏","安徽","湖北","重庆","贵州","湖南","江西","上海","浙江","福建","云南","广西","广东","海南","香港","澳门","台湾","宁夏"}
	h:=ProvinceFind(provinces)

	m:=make(map[string]interface{})
	cityname:=c.Param("cityname")
	province:=crud.CrudFind(cityname)
	//fmt.Printf("打印id %v  %v",province.Id,province.ProvinceName)
	citys:=crud.CrudFindProvinceId(province.Id)
	//fmt.Printf("打印 %v",citys)
	for _,city:=range citys{
		m[city.Name]=city
	}
	parameter:=Parameter{ParameterProvince:h,
		ParameterCity:m,
	}
	c.HTML(http.StatusOK,"china2.html",parameter)
}
//多次使用的代码
func ProvinceFind(provinces []string)map[string]interface{}{
	h:=make(map[string]interface{})
	for _,province:=range provinces{
		data:=crud.CrudFind(province)
		h[data.ProvinceName]=data.Confirm
	}
	return h
}








