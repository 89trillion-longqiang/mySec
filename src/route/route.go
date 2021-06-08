package route

import (
	"strings"

	"calcula"
	"github.com/gin-gonic/gin"
)

func SetUpRount() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/C1")

	c1.GET("/calculator", getCalculator)
	return r
}
func getCalculator(c *gin.Context){
	uri := c.Request.RequestURI///获取url
	newExp := strings.Replace(uri,"/C1/calculator?exp=","",1)///将获取的url切分，仅保留表达式部分

	ret,err := calcula.Calculate(newExp)
	if err == nil{
		c.JSON(200,gin.H{
			"condition" : "pass",
			"result" : ret ,
		})
	}else {
		c.JSON(200,gin.H{
			"condition" : "error",
			"result"  : err.Error(),
		})
	}
}