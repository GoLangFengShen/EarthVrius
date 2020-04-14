package parser

import (
	"crawler/enginetype"
	"crawler/model"
	"regexp"
)

//在是精筛选出详细的对应数据数据
func FindSub(rexp *regexp.Regexp,finedata []byte)string{
	data:=rexp.FindSubmatch(finedata)
	//return string(data[1])
	if len(data)>=2{
		return string(data[1])
	}else {
		return ""
	}
}

func FineSelsctContent(provincedata [][]byte,CityRegexp enginetype.Data)model.ProvinceData{
	var Rname=regexp.MustCompile(CityRegexp.Parameter.FineVirusRegexp.NameRegexp)
	var RtodayConfirm=regexp.MustCompile(CityRegexp.Parameter.FineVirusRegexp.TodayConfirmRegexp)
	var Rconfirm=regexp.MustCompile(CityRegexp.Parameter.FineVirusRegexp.ConfirmRegexp)

	var Rheal=regexp.MustCompile(CityRegexp.Parameter.FineVirusRegexp.HealRegexp)
	var Rdead=regexp.MustCompile(CityRegexp.Parameter.FineVirusRegexp.DeadRegexp)
	v:=model.ProvinceData{}
	//遍历得到每个地方的数据
	for _,Finedata:=range provincedata{

		//在具体的城市找数据
		virusdata:=model.VirusData{}
		name:=FindSub(Rname,Finedata)
		virusdata.Name=name
		todayconfirm:=FindSub(RtodayConfirm,Finedata)
		virusdata.Todayconfirm=todayconfirm
		confirm:=FindSub(Rconfirm,Finedata)
		virusdata.Confirm=confirm
		heal:=FindSub(Rheal,Finedata)
		virusdata.Heal=heal
		dead:=FindSub(Rdead,Finedata)
		virusdata.Dead=dead

		v.ProvinceVirus=append(v.ProvinceVirus,virusdata)

	}
	return v
}

