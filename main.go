package main

import (
	"calculator/route"
)

func main(){

	r := route.SetUpRount()
	r.Run(":9000")
}



