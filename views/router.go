package views

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
)

func ConfRouter(r *gin.Engine)  {
	//r.LoadHTMLGlob("templates/*") // 静态模板
	//r.Static("/static", "./static") // 静态文件
	r.MaxMultipartMemory = 32 << 20  // 32M
	kingAdminRouter(r)
}

func kingAdminRouter(r *gin.Engine)  {
	//g.Use(middleware.Authentication()) // 认证数据
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/version", Version)
	r.GET("/",Index)
	g := r.Group("/kingadmin")

	service(g)

}

func service(r *gin.RouterGroup)  {
	g := r.Group("/service")
	g.GET("/list.html",ServiceList)

}

