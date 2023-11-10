package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main()  {
	var s,sep string
	fmt.Println(os.Args[0:])
	fmt.Println(strings.Join(os.Args[1:], " "))
	for i:=1;i<len(os.Args);i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(os.Args[i])
		fmt.Printf("%d:%s\n",i,os.Args[i]);
	}
	fmt.Println(s)

	fmt.Print("hello world\n")
	fmt.Println("hello world")
	fmt.Printf("布尔：%t\n",0==2)
	fmt.Printf("十进制：%d\n",255)
	fmt.Printf("二进制：%b\n",255)
	fmt.Printf("八进制：%o\n",255)
	fmt.Printf("十六进制：%X\n",255)
	fmt.Printf("浮点：%f\n",math.Pi)
	fmt.Printf("字符串：%s\n","hello world")

}