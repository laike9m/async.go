package async

import (
	"fmt"
	"testing"
)

func TestEach(t *testing.T) {
	a := []int{1, 2, 3}
	Each(a, func(e int) {
		fmt.Println(e)
	})
	b := []string{"1", "2", "3"}
	Each(b, func(e string) {
		fmt.Println(e)
	})
}
