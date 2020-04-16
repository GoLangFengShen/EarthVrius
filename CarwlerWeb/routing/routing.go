package routing

import (
	"CarwlerWeb/contronller"
	"CarwlerWeb/model"
	"github.com/gin-gonic/gin"
)

func Routing()*gin.Engine{
	//创建gin的路由引擎
	model.R=gin.Default()
	//加载静态文件
	model.R.Static("/static","./static")
	//加载模板
	model.R.LoadHTMLFiles("static/earth.html","static/china.html","static/china2.html")

	model.R.GET("/earth", contronller.Earth)
	model.R.GET("/earth/:name", contronller.EarthData)
	model.R.GET("/china",contronller.EarthChina)
	model.R.GET("/china/:cityname",contronller.ChinaCity)
	return model.R
}