package views

import "github.com/gin-gonic/gin"

func Version(g *gin.Context)  {
	data := ResponseData{Status:true,Date:"0.0.1"}
	g.JSON(200,data)
}



type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Cid     int64   `json:"cid"`
	Date    interface{} `json:"data"`
}