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
	host   string
	signer Signer
}

func NewApi(url string) *Api {
	return &Api{
		host: url,
	}
}

func (api Api) GetBlockByHeight(height uint64) (*Block, error) {
	block := new(Block)
	result, err := api.request("Gzv", "getBlockByHeight", height)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, block)
	return block, err
}

//func (api Api) GetBlockDetailByHeight(height uint64) (*BlockDetail, error) {
//	hash, err := api.GetBlockHashByHeight(height)
//	if err != nil {
//		return nil, err
//	}
//	return api.GetBlockDetailByHash(*hash)
//}

func (api Api) GetBlockByHash(hash Hash) (*Block, error) {
	block := new(Block)
	result, err := api.request("Gzv", "getBlockByHash", hash.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, block)
	return block, err
}

//func (api Api) GetBlockDetailByHash(hash Hash) (*BlockDetail, error) {
//	blockDetail := new(BlockDetail)
//	result, err := api.request("Dev", "blockDetail", hash.String())
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(*result, blockDetail)
//	return blockDetail, err
//}

func (api Api) GetTransactionByHash(hash Hash) (*RawTransaction, error) {
	tx := new(RawTransaction)
	result, err := api.request("Gzv", "transDetail", hash.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, tx)
	return tx, err
}

func (api Api) BlockHeight() (uint64, error) {
	var height uint64
	result, err := api.request("Gzv", "blockHeight")
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(*result, &height)
	return height, err
}

func (api Api) GetBlockHashByHeight(height uint64) (*Hash, error) {
	block := new(Block)
	result, err := api.request("Gzv", "getBlockByHeight", height)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, block)
	return &block.Hash, err
}

func (api Api) GetNonce(address Address) (uint64, error) {
	var nonce uint64
	result, err := api.request("Gzv", "nonce", address.String())
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(*result, &nonce)
	return nonce, err
}

func (api *Api) SetSigner(signer Signer) {
	api.signer = signer
}

func (api Api) SendTransaction(tx Transaction) (*Hash, error) {
	hash := new(Hash)
	jsonByte, err := json.Marshal(tx.ToRawTransaction())
	if err != nil {
		return nil, err
	}
	result, err := api.request("Gzv", "tx", string(jsonByte))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, hash)
	return hash, err
}

func (api Api) SignAndSendTransaction(tx Transaction) (*Hash, error) {
	rawTransaction := tx.ToRawTransaction()
	if rawTransaction.Source == nil {
		return nil, ErrorSignerNotFound
	}
	rawTransaction.Hash = rawTransaction.GenHash()
	sign, err := api.signer.Sign(rawTransaction)
	if err != nil {
		return nil, err
	}
	rawTransaction.Sign = sign
	hash, err := api.SendTransaction(rawTransaction)
	if err != nil {
		return nil, err
	}
	return hash, nil
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
	result, err := api.request("Gzv", "minerInfo", address.String(), detail)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(*result, stakeDetails)
	return stakeDetails, err
}

func (api Api) Balance(address Address) (float64, error) {
	var balance float64
	result, err := api.request("Gzv", "balance", address.String())
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(*result, &balance)
	return balance, err
}
