package main

import (
	"fmt"
	"github.com/zhadowz99/calculator/addition"
)

func main(){
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b: ")
	fmt.Scanln(&b)
	c := addition.Sum(a,b)
	fmt.Print(c,"\n")
}
