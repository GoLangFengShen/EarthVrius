package scheduler

import (
	"crawler/enginetype"
)

type SimpleScheduler struct {
	WorkerChan chan enginetype.Data
}
//type ParserScheduler struct {
//	ProvinceChan chan model.ProvinceData
//}
func (s *SimpleScheduler)UrlToChan(url enginetype.Data){
	//开启goroutine供应CreateWorker
	go func() {
		s.WorkerChan <- url
	}()
}
//输出转输入
func (s *SimpleScheduler)EqualChan(in chan enginetype.Data){
	s.WorkerChan=in
}


