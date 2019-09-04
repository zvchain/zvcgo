package zvcgo

import (
	"fmt"
	"testing"
)

func TestNewAssetFromString(t *testing.T) {
	s, err := NewAssetFromString("1ZVC")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.value)
}
