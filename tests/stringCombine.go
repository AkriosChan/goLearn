package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("aaaa")

	var i int = 12222
	var stra string = "a"
	var strb string = "b"
	var strc string = "c"
	//fmt.Println(stra + strb + strc + i)
	fmt.Println(stra+strb+strc, strconv.Itoa(i))
}
