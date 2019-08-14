package zvlib

import (
	"fmt"
	"testing"
)

var api = NewApi("http://node2.zvchain.io:8101")

func TestApi_BlockHeight(t *testing.T) {
	height, err := api.BlockHeight()
	fmt.Println("height:", height)
	fmt.Println("err:", err)
}

//func TestApi_GetNonce(t *testing.T) {
//	nonce, err := api.GetNonce("")
//	fmt.Println("nonce:",nonce)
//	fmt.Println("err:",err)
//}
//
//func TestApi_GetCode(t *testing.T) {
//	code, err := api.GetCode("")
//	fmt.Println("code:",code)
//	fmt.Println("err:",err)
//}

func TestApi_GetBlockHashByHeight(t *testing.T) {
	blockDetail, err := api.GetBlockDetailByHeight(100)
	fmt.Println(blockDetail)
	fmt.Println(err)
}

func TestApi_GetBlockByHeight(t *testing.T) {
	block, err := api.GetBlockByHeight(100)
	fmt.Println(block)
	fmt.Println(err)
}
