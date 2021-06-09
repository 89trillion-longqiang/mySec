package main

import (
	"fmt"
	"log"
	"testing"

	"calcula"
)

func Test_Calculate(t *testing.T) {

	exp := "3+3*4/2+3"
	ret, err := calcula.Calculate(exp)

	if ret == 12 && err != nil {
		fmt.Println("Test_Calculate pass")
	} else {
		log.Fatal("Test_Calculate error")
	}
}
