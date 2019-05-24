package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x <10 && x > 0{
		return true
	}
	if x < 0{
		return false
	}
	val := int(math.Log10(float64(x)))
	var end = 0

	wei := int(math.Ceil(float64(val + 1) / 2))
	for i := 1; i <= wei; i++{
		end = (x % 10) + end * 10
		if i == wei && val % 2 != 0{
			x = x / 10
			break
		}else if i == wei && val % 2 == 0{
			break
		}else {
			x = x / 10
		}

	}
	if x == end{
		return true
	}else{
		return false
	}
}

func main() {
	for i := -10; i < 10000; i++{
		//if !isPalindrome(121){
		fmt.Println(isPalindrome(121))
		//}
		break
	}

}
