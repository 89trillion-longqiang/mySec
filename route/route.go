package route

import (
	"strconv"
	"strings"

	"calculator/calcula"
	"github.com/gin-gonic/gin"
)

func SetUpRoute() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/C1")

	c1.GET("/calculator", getCalculator)
	return r
}
func HandleGetCalculator(uri string)map[string]string{
	retMap := make(map[string]string,2)
	newExp := strings.Replace(uri,"/C1/calculator?exp=","",1)///将获取的url切分，仅保留表达式部分

	ret,err := calcula.Calculate(newExp)
	retString := strconv.Itoa(ret)
	if err == nil{
		retMap["condition"] = "pass"
		retMap["result"]= retString
	}else {
		retMap["condition"] = "error"
		retMap["result"]  = err.Error()
	}


	return retMap
}
func getCalculator(c *gin.Context){
	uri := c.Request.RequestURI///获取url
	ret := HandleGetCalculator(uri)

		c.JSON(200,gin.H{
			"condition": ret["condition"],
			"result":   ret["result"],
		})
}