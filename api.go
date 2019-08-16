package zvlib

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	NoSignerError = errors.New("signer should be set before send a no-sign transaction")
)

type Api struct {
	host string
	Signer
}

func NewApi(url string) *Api {
	return &Api{
		host: url,
	}
}

func (api Api) GetBlockByHeight(height uint64) (*Block, error) {
	block := new(Block)
	err := api.request("Gzv", "getBlockByHeight", block, height)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (api Api) GetBlockDetailByHeight(height uint64) (*BlockDetail, error) {
	hash, err := api.GetBlockHashByHeight(height)
	if err != nil {
		return nil, err
	}
	return api.GetBlockDetailByHash(*hash)
}

func (api Api) GetBlockByHash(hash Hash) (*Block, error) {
	block := new(Block)
	err := api.request("Gzv", "getBlockByHash", block, hash.String())
	return block, err
}

func (api Api) GetBlockDetailByHash(hash Hash) (*BlockDetail, error) {
	blockDetail := new(BlockDetail)
	err := api.request("Dev", "blockDetail", blockDetail, hash.String())
	if err != nil {
		return nil, err
	}
	return blockDetail, nil
}

func (api Api) GetTransactionByHash(hash Hash) (*Transaction, error) {
	tx := new(Transaction)
	err := api.request("Gzv", "transDetail", tx, hash.String())
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (api Api) BlockHeight() (uint64, error) {
	var height uint64
	err := api.request("Gzv", "blockHeight", &height)
	return height, err
}

func (api Api) GetBlockHashByHeight(height uint64) (*Hash, error) {
	block := new(Block)
	err := api.request("Gzv", "getBlockByHeight", block, height)
	if err != nil {
		return nil, err
	}
	return &block.Hash, nil
}

func (api Api) GetNonce(address Address) (uint64, error) {
	var nonce uint64
	err := api.request("Gzv", "nonce", &nonce, address.String())
	return nonce, err
}

func (api Api) GetCode(address Address) (string, error) {
	accountMsg := new(AccountMsg)
	var code string
	err := api.request("Gzv", "viewAccount", accountMsg, address.String())
	if err != nil {
		return code, err
	}
	code = accountMsg.Code
	return code, err
}

//todo
func (api Api) GetData(address Address, key string) (interface{}, error) {
	accountMsg := new(AccountMsg)
	err := api.request("Gzv", "viewAccount", accountMsg, address.String(), key)
	if err != nil {
		return nil, err
	}
	return accountMsg, err
}

func (api Api) SetSigner(signer Signer) {
	api.Signer = signer
}

//todo
func (api Api) SendRawTransaction(tx RawTransaction) (*Hash, error) {

	// todo get nonce
	if tx.ToRawTransaction().Nonce == 0 {
		var nonce uint64
		err := api.request("Gzv", "nonce", &nonce, tx.ToRawTransaction().Source.String())
		if err != nil {
			return nil, err
		}
		tx.ToRawTransaction().Nonce = nonce
	}

	hash := new(Hash)
	jsonByte, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	//todo
	err = api.request("Gzv", "tx", hash, string(jsonByte))
	if err != nil {
		return nil, err
	}
	return hash, err
}

//todo
func (api Api) SignAndSendRawTransaction(tx RawTransaction) (*Hash, error) {

	sign, err := api.Sign(tx)
	fmt.Println("SIGN:", sign)
	if err != nil {
		return nil, err
	}
	rawTx := tx.ToRawTransaction()
	rawTx.Sign = sign.Bytes()
	fmt.Println(rawTx)

	hash, err := api.SendRawTransaction(rawTx)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (api Api) GetPastEvent(address Address, topic string, from, to uint64) ([]*Event, error) {
	return nil, nil
}

func (api Api) EventListen(address Address, topic string, from uint64, callBack EventCallBack) error {
	return nil
}

func (api Api) MinerInfo(address Address, detail interface{}) (*MinerStakeDetails, error) {
	stakeDetails := new(MinerStakeDetails)

	switch detail.(type) {
	case Address:
		detail = detail.(Address).String()
	case string:
		if detail != "" && detail != "all" {
			return nil, fmt.Errorf("params input err, please check it carefully")
		}
	}
	err := api.request("Gzv", "minerInfo", stakeDetails, address.String(), detail)
	return stakeDetails, err
}

func (api Api) Balance(address Address) (float64, error) {
	var balance float64
	err := api.request("Gzv", "balance", &balance, address.String())
	return balance, err
}

type EventCallBack func(event *Event)

func A(event *Event) {

}
