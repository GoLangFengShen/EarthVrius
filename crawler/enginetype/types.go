package enginetype
//数据携带者，给每一步提供数据
type Data struct{
	Request  []byte
	Result interface{}
	Next func(Data)DatasParsre
	Parameter Parameter
}
//函数返回方便
type DatasParsre struct {
	Datas []Data
}
//细化数据的正则
type FineVirusRegexp struct {
	NameRegexp string
	TodayConfirmRegexp string
	ConfirmRegexp string
	HealRegexp string
	DeadRegexp string
}
//fetcher和parser所需的正则及FineVirusRegexp
type Parameter struct {
	Address string
	ProvinceRegexp string
	CityRegexp string
	FineVirusRegexp FineVirusRegexp
}
type Parameters struct {
	Parameters []Parameter
}