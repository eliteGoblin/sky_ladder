package main

import "fmt"

func multiply(num1 string, num2 string) string {
	res := ""
	multiDigits := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			p1, p2 := i+j, i+j+1
			sum := mul + multiDigits[p2]
			multiDigits[p1] += sum / 10
			multiDigits[p2] = sum % 10
			fmt.Printf("i %d j %d ", i, j)
			print(multiDigits)
		}
	}
	for _, v := range multiDigits {
		if len(res) > 0 || v != 0 {
			res += string(v + '0')
		}
	}
	if res == "" {
		return "0"
	}
	return res
}

func print(arr []int) {
	for _, v := range arr {
		fmt.Printf(" %d", v)
	}
	fmt.Println("")
}

func main() {
	res := multiply("89", "76")
	fmt.Println(res)
}
