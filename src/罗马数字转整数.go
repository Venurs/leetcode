package main

import "fmt"

func romanToInt(s string) int {
	var y = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for index, str := range s{
		res := y[string(str)]
		if index != 0{
			prev := y[string(s[index - 1])]
			if res > prev{
				sum = sum - 2 * prev + res
			}else{
				sum = sum + res
			}
		}else{
			sum = sum + res
		}
	}

	return sum
}
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
