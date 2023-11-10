package main

import (
	"fmt"
	"math"
)


func main() {
	var s string = "hello world"
	var n int = 11
	var i,j,k int = 1,2,3
	var b = true
	a := 199
	const (
		pi1 = 3.14
		pi = math.Pi
	)
	fmt.Printf("a type:%T\n",a)
	fmt.Printf("pi type:%T\n",pi)
	fmt.Printf("pi1 type:%T\n",pi1)
	fmt.Println(s,n,i,j,k,b,a,pi1,pi)
}