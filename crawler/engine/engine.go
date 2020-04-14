package engine

import (
	"crawler/conndatabase"
	"crawler/enginetype"
	"crawler/fetcher"
	"fmt"
	"time"
)
type ConcurrentEngine struct{
	Scheduler Scheduler
	LenConcurrent int
}
type Scheduler interface {
	UrlToChan(enginetype.Data)
	EqualChan(chan enginetype.Data)
}
//计数
var i int
func (e ConcurrentEngine)Run(Parameter enginetype.Parameter) {
	//先做转换成utf8
	url:=fetcher.FetcherContent(Parameter)
	//确定输入输出
	in:=make(chan enginetype.Data)
	out:=make(chan enginetype.DatasParsre)
	//接收要转换成in的数据
	e.Scheduler.EqualChan(in)
	//将url扔进去转化
	e.Scheduler.UrlToChan(url)
	//开CreateWorker的goroutine
	for i := 0; i < e.LenConcurrent; i++{
		CreateWorker(in,out)
	}
//循环接收数据
	for {
		select {
			case citydata := <-out:

			for _, virus := range citydata.Datas {
				                                                                          //fmt.Printf("城市:%s   新增确诊:%s  累计确诊:%s  疑似:%s  治愈人数:%s  死亡:%s \n",virus.Name,virus.Todayconfirm,virus.Confirm,virus.Suspect,virus.Heal,virus.Dead)
				//l--
				if virus.Request != nil {
					e.Scheduler.UrlToChan(virus)
					//fmt.Printf("数据 %s \n\n",virus.Request)
				} else {
					if virus.Result != nil {
						i++
						fmt.Printf("数据 %v 计数%d\n\n", virus.Result, i)
						conndatabase.SaveData(virus.Result)
					}
				}
			}
		                                                                                       	//做准备，不然前面没得到数据这直接关了
		case <-time.After(1*time.Second):
			return

			}
		}
}

func CreateWorker(in chan enginetype.Data,out chan enginetype.DatasParsre){
	go func() {
		Url:=<-in
		data:=Worker(Url)
		out<-data
	}()
}