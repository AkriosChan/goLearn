package main

import (
	_ "fmt"
	"time"
)

const (
	pi = 3.1415926
)

func main() {
	//fmt.Println(pi)
	forTest()
}

func forTest() {
	sum := 0
	//count := 0
	start := time.Now().UnixNano()
	for sum < 10000 {
		println(">>>>>>>>>>")
		sum++
		//count++
	}
	end := time.Now().UnixNano()
	print((end - start) / 1000000)

}
