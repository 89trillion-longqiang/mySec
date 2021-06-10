package main

import (
	"calculator/route"
)

func main(){

	r := route.SetUpRoute()
	r.Run(":9000")
}



