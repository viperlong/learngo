package main

import (
	"fmt"
)

func add(x,y int) int {return x+y}


func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}


func Fun() func(string) string {
    a := "world"
    return func(args string) string {
        a += args
        return  a
    }
}

func sum(vals ...int) (total int) {
	total = 0
	for _,val := range vals {
		total += val;
	}
	return total;
} 

func sumwithdefer(vals ...int) (total int) {
	defer func () {total=total +1}()
	total = 0
	for _,val := range vals {
		total += val;
	}
	return total;
} 
// func main() {
//     a := Fun()
//     b:=a("hello ")
//     c:=a("hello ")
//     fmt.Println(b)//worldhello 
//     fmt.Println(c)//worldhello hello 
// }

func main() {
	f := squares();
	fmt.Printf("%T\n",add)
	fmt.Println(add(2,3))
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(sum(1,2,3,4))
	fmt.Println(sumwithdefer(1,2,3,4))
}