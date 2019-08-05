package amazon

import (
	"fmt"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)
	fmt.Println(finder.FindMedian())
	finder.AddNum(3)
	fmt.Println(finder.FindMedian())
}