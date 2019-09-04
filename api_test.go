package zvcgo

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var api = NewApi("http://120.77.155.204:8102")
var (
	hash1, _  = hex.DecodeString("7a41f2571af87bf9e798828f4b1cd9638b770b105f5821cc060ff6209d41127c")
	blockHash = Hash{
		data: hash1,
	}

	hash2, _ = hex.DecodeString("b0f136eca631083fadd3333b30614de7ba4d1a67e4411005833667ae1f687270")
	txHash   = Hash{
		data: hash2,
	}

	addr1, _    = hex.DecodeString("74800b982ff90fb5b36709d886554f28e43458ac8cf1974d7cfe4d683f5818e5")
	accountAddr = Address{
		data: addr1,
	}

	addr2, _  = hex.DecodeString("afebb47a1abeb2bbe2470ffadc90ee5ec546ba711380d33fe8e66e6daa40a186")
	minerAddr = Address{
		data: addr2,
	}

	addr3, _     = hex.DecodeString("f90a829af7fd9196f98fb57d412586ff4949214d639ada48a047ba028fa585d9")
	contractAddr = Address{
		data: addr3,
	}
)

func TestApi_GetBlockByHeight(t *testing.T) {
	block, err := api.GetBlockByHeight(10)
	fmt.Println("block:", block)
	fmt.Println(err)
}

//func TestApi_GetBlockDetailByHeight(t *testing.T) {
//	blockDetail, err := api.GetBlockDetailByHeight(2119)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(blockDetail)
//}

func TestApi_GetBlockByHash(t *testing.T) {
	block, err := api.GetBlockByHash(blockHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(block)
}

//func TestApi_GetBlockDetailByHash(t *testing.T) {
//	blockDetail, err := api.GetBlockDetailByHash(blockHash)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(blockDetail)
//	for _, v := range blockDetail.GenRewardTx.TargetIDs {
//		res, _ := v.MarshalJSON()
//		fmt.Println(string(res))
//	}
//}

func TestApi_GetTransactionByHash(t *testing.T) {
	tx, err := api.GetTransactionByHash(txHash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tx)
}

func TestApi_BlockHeight(t *testing.T) {
	height, err := api.BlockHeight()
	fmt.Println("height:", height)
	fmt.Println("err:", err)
}

func TestApi_GetBlockHashByHeight(t *testing.T) {
	hash, err := api.GetBlockHashByHeight(2119)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("hash:", hash)
}

func TestApi_GetNonce(t *testing.T) {
	nonce, err := api.GetNonce(accountAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("nonce:", nonce)
}

//func TestApi_GetCode(t *testing.T) {
//	addr,err := Decode("0xedf103af25f0bb800fdd057e201200f891ece964a3d7d14df72be12a8017fccb")
//	if err != nil {
//		t.Error(err)
//	}
//	addr1 := Address{
//		data:addr,
//	}
//	res,err := api.GetCode(addr1)
//	fmt.Println(res)
//	fmt.Println(err)
//}
//
func TestApi_MinerInfo(t *testing.T) {
	res, err := api.MinerInfo(minerAddr, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*res)
}

func TestApi_Balance(t *testing.T) {
	balance, err := api.Balance(minerAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balance)
}
