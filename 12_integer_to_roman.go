package amazon

var val []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var name []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	res := ""
	for i := range val {
		for num >= val[i] {
			res += name[i]
			num -= val[i]
		}
	}
	return res
}
