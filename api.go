package zvlib

import (
	"errors"
)

var (
	NoSignerError = errors.New("signer should be set before send a no-sign transaction")
)

type Api struct {
	host string
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

// todo
func (api Api) GetBlockByHash(hash Hash) (*Block, error) {
	block := new(Block)
	err := api.request("Gzv", "getBlockByHash", block, hash)
	return block, err
}

//todo
func (api Api) GetBlockDetailByHash(hash Hash) (*BlockDetail, error) {
	blockDetail := new(BlockDetail)
	err := api.request("Dev", "blockDetail", blockDetail, hash)
	if err != nil {
		return nil, err
	}
	return blockDetail, nil
}

func (api Api) GetTransactionByHash(hash Hash) (*Transaction, error) {
	return nil, nil
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
	err := api.request("Gzv", "nonce", &nonce, address)
	return nonce, err
}

// for contract? todo
func (api Api) GetCode(address Address) (string, error) {
	accountMsg := new(AccountMsg)
	var code string
	err := api.request("Gzv", "viewAccount", accountMsg, address)
	if err != nil {
		return code, err
	}
	code = accountMsg.Code
	return code, err
}

//todo
func (api Api) GetData(address Address, key string) (interface{}, error) {
	accountMsg := new(AccountMsg)
	err := api.request("Gzv", "viewAccount", accountMsg, address, key)
	return accountMsg, err
}

func (api Api) SetSigner(signer Signer) {

}

func (api Api) SendRawTransaction(tx RawTransaction) (*Hash, error) {
	return nil, nil
}

func (api Api) SignAndSendRawTransaction(tx RawTransaction) (*Hash, error) {
	return nil, nil
}

func (api Api) GetPastEvent(address Address, topic string, from, to uint64) ([]*Event, error) {
	return nil, nil
}

func (api Api) EventListen(address Address, topic string, from uint64, callBack EventCallBack) error {
	return nil
}

func (api Api) MinerInfo(address Address) error {
	return nil
}

func (api Api) Balance(address Address) error {
	return nil
}

type EventCallBack func(event *Event)

func A(event *Event) {

}
