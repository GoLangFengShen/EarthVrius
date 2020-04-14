package parser

import (
	"crawler/enginetype"
	"regexp"
)

func CityParser(provincedata enginetype.Data) enginetype.DatasParsre {

		CityRegexp:=regexp.MustCompile(provincedata.Parameter.CityRegexp)
		province:=CityRegexp.FindAll(provincedata.Request,-1)
		//传入细化数据得到每个省对应的集合
		pr:=FineSelsctContent(province,provincedata)
		s:= enginetype.DatasParsre{}
		r:= enginetype.Data{Result:pr}
		s.Datas=append(s.Datas,r)

		return s
}
