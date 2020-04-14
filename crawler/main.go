package main

import (
	"crawler/engine"
	"crawler/enter"
	"crawler/model"
	"crawler/scheduler"
)
func main(){
	e:=engine.ConcurrentEngine{
		//实例结构体，让引擎的方法用起来
		Scheduler:     &scheduler.SimpleScheduler{},
		//开的太少不够玩，没有分配到goroutine
		LenConcurrent: 200,
	}
	parameters:=enter.Enter()
	err:=model.InitMysql()
	if err!=nil{
		panic(err)
	}
	defer model.Db.Close()
	for _,parameter:=range parameters.Parameters{
		e.Run(parameter)

	}
}

