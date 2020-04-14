package enter

import "crawler/enginetype"

func Enter()enginetype.Parameters{
	//国内
	fineVirusRegexp:=enginetype.FineVirusRegexp{NameRegexp:`name\\":\\"([\W]*)\\",\\"today`,
		TodayConfirmRegexp:`today\\":{\\"confirm\\":([0-9]+)`,
		ConfirmRegexp:`confirm\\":([0-9]+),\\"suspect`,
		HealRegexp:`heal\\":([0-9]+),\\"healRate`,
		DeadRegexp:`dead\\":([0-9]+),\\"deadRate`,
	}
	src:=enginetype.Parameter{
		Address:"https://view.inews.qq.com/g2/getOnsInfo?name=disease_h5&callback=jQuery34106305350739632178_1584331066813&_=1584331066814",
		ProvinceRegexp:`{\\"name\\":[^>]*?}}]}`,
		CityRegexp:`{\\"name\\":\\"[\W]*\\",\\"today\\":{\\"confirm\\":[0-9]+[^>]*?},\\"total\\":{\\"nowConfirm\\":[0-9]+,\\"confirm\\":[0-9]+,\\"suspect\\":[0-9]+,\\"dead\\":[0-9]+,\\"deadRate\\":\\"[0-9]+.[0-9]+\\",\\"showRate\\":[a-zA-Z]*,\\"heal\\":[0-9]+,\\"healRate\\":\\"[0-9]+.[0-9]+\\",\\"showHeal\\":[a-zA-Z]*}`,
		FineVirusRegexp:fineVirusRegexp,
	}
	//国外
	fineAbroadVirusRegexp:=enginetype.FineVirusRegexp{
		NameRegexp:         `"name\\":\\"([\W]*)\\",`,
		TodayConfirmRegexp: `confirmAdd\\":([0-9]+),`,
		ConfirmRegexp:      `"confirm\\":([0-9]+),`,
		HealRegexp:         `"dead\\":([0-9]+),`,
		DeadRegexp:         `"heal\\":([0-9]+)`,
	}
	abroadSrc:=enginetype.Parameter{
		Address:"https://view.inews.qq.com/g2/getOnsInfo?name=disease_foreign&callback=jQuery34107360506572093108_1584806738559&_=1584806738560",
		//ProvinceRegexp:`{\\"name\\":\\"[\W]*\\",\\"continent\\":\\"[\W]*\\",\\"date\\":\\"[0-9]+.[0-9]+\\",\\"isUpdated\\":[a-zA-Z]*,\\"confirmAdd\\":[0-9]+,\\"confirmAddCut\\":[0-9]+,\\"confirm\\":[0-9]+,\\"suspect\\":[0-9]+,\\"dead\\":[0-9]+,\\"heal\\":[0-9]+,\\"nowConfirm\\":[0-9]+,\\"confirmCompare\\":[^\\]+,\\"nowConfirmCompare\\":[0-9]+,\\"healCompare\\":[0-9]+,\\"deadCompare\\":[0-9]+`,
		ProvinceRegexp:`[{\\"name\\":\\"[\W]*?\\",\\"continent[^>]*?healCompare\\":[0-9]+,\\"deadCompare\\":[0-9]+[^>]*?]*?`,
		CityRegexp:`\\"name\\":\\"[\W]*\\"[^>]*?confirmAdd\\":[0-9]+,\\"confirmAddCut\\":[0-9]+,\\"confirm\\":[0-9]+,\\"suspect\\":[0-9]+,\\"dead\\":[0-9]+,\\"heal\\":[0-9]+`,
		FineVirusRegexp:fineAbroadVirusRegexp,
	}
	parameters:=enginetype.Parameters{}
	parameters.Parameters=append(parameters.Parameters,src,abroadSrc)
	return parameters
}