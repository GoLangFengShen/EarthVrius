package engine

import (
	"crawler/enginetype"
)


func Worker(Url enginetype.Data) enginetype.DatasParsre {
//用他自身带的函数，确定下一步
	L:=Url.Next(Url)
	return L
}


