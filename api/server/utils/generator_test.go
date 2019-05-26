package utils

import (
	"fmt"
	"testing"
)

func TestGenerateOrderNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := GenerateOrderNumber(10)
		fmt.Println(res)
	}
}
