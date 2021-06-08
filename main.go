package main

import (
	"route"
)

func main(){

	r := route.SetUpRount()
	r.Run(":9000")
}



