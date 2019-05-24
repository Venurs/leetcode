package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	const INT_MAX = int(^uint32(0) >> 1)
	const INT_MIN = ^INT_MAX
	var res = ""
	var base = 1
	if x < 0{
		base = -1
		x = base * x
	}else{
		base = 1
	}
	strX := strconv.Itoa(x)
	for  i := len(strX) - 1; i >= 0; i-- {
		res = res + string(strX[i])
	}
	val, err :=strconv.Atoi(res)
	if val > INT_MAX || val < INT_MIN{
		return 0
	}
	if err != nil{
		return 0
	}else{
		return val * base
	}
}

func main() {
	fmt.Println(reverse(1534236469))
}
