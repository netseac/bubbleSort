package bubble

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{10, 9, 7, 4, 3, 2, 1, 8, 6, 5}
	s := sort(a)
	fmt.Println(s)
}
