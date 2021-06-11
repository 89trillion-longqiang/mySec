package route

import (
	"calculator/handle"
	"github.com/gin-gonic/gin"
)

func SetUpRoute() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/C1")

	c1.GET("/calculator", getCalculator)
	return r
}

func getCalculator(c *gin.Context){
	uri := c.Request.RequestURI///获取url
	ret := handle.HandleGetCalculator(uri)

		c.JSON(200,gin.H{
			"condition": ret["condition"],
			"result":   ret["result"],
		})
}