package parser

import (
	"crawler/enginetype"
	"regexp"
)


func SelectContent(content enginetype.Data) enginetype.DatasParsre {
	//先以省为单位正则筛选出所需主要内容
	e:=regexp.MustCompile(content.Parameter.ProvinceRegexp)
	c:= enginetype.DatasParsre{}
	ProvinceDatas:=e.FindAll(content.Request,-1)
	for _,provinceData:=range ProvinceDatas{

	c.Datas=append(c.Datas, enginetype.Data{
		provinceData,
			nil,
			CityParser,
			content.Parameter,
		})
	}
	return c
}