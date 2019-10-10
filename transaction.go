package zvcgo

import (
	"encoding/json"
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/middleware/types"
)

const (
	DefaultGasPrice        = 500
	DefaultGasLimit        = 3000
	CodeBytePrice          = 19073 //1.9073486328125
	CodeBytePricePrecision = 10000
)

const (
	TransactionTypeTransfer       = 0
	TransactionTypeContractCreate = 1
	TransactionTypeContractCall   = 2
)

type SimpleTx struct {
	Data   []byte   `json:"data"`
	Value  float64  `json:"value"`
	Nonce  uint64   `json:"nonce"`
	Source *Address `json:"source"`
	Target *Address `json:"target"`
	Type   int8     `json:"type"`

	GasLimit uint64 `json:"gas_limit"`
	GasPrice uint64 `json:"gas_price"`
	Hash     Hash   `json:"hash"`

	ExtraData string `json:"extra_data"`
}

type RawTransaction struct {
	Data   []byte   `json:"data"`
	Value  uint64   `json:"value"`
	Nonce  uint64   `json:"nonce"`
	Target *Address `json:"target"`
	Type   int8     `json:"type"`

	GasLimit uint64 `json:"gas_limit"`
	GasPrice uint64 `json:"gas_price"`
	Hash     Hash   `json:"hash"`

	ExtraData []byte   `json:"extra_data"`
	Sign      *Sign    `json:"sign"`
	Source    *Address `json:"source"`
}

func (t RawTransaction) ToRawTransaction() RawTransaction {
	return t
}

func (t *RawTransaction) SetGasPrice(gasPrice uint64) *RawTransaction {
	t.GasPrice = gasPrice
	return t
}

func (t *RawTransaction) SetGasLimit(gasLimit uint64) *RawTransaction {
	t.GasPrice = gasLimit
	return t
}

func (t *RawTransaction) SetNonce(nonce uint64) *RawTransaction {
	t.Nonce = nonce
	return t
}

func (t *RawTransaction) SetData(data []byte) *RawTransaction {
	t.Data = data
	return t
}

func (t *RawTransaction) SetExtraData(extraData []byte) *RawTransaction {
	t.ExtraData = extraData
	return t
}

func (t RawTransaction) GenHash() Hash {
	var target, source *common.Address
	if t.Target == nil {
		target = nil
	} else {
		tmp := common.BytesToAddress(t.Target.Bytes())
		target = &tmp
	}
	if t.Source == nil {
		source = nil
	} else {
		tmp := common.BytesToAddress(t.Source.Bytes())
		source = &tmp
	}
	tx := types.RawTransaction{
		Data:      t.Data,
		Value:     types.NewBigInt(t.Value),
		Nonce:     t.Nonce,
		Target:    target,
		Source:    source,
		Type:      t.Type,
		GasLimit:  types.NewBigInt(t.GasLimit),
		GasPrice:  types.NewBigInt(t.GasPrice),
		ExtraData: t.ExtraData,
	}
	return Hash{tx.GenHash().Bytes()}
}

type Transaction interface {
	ToRawTransaction() RawTransaction
}

type TransferTransaction struct {
	RawTransaction
}

func NewTransferTransaction(from, to Address, value Asset, data []byte) *TransferTransaction {
	gaslimit := len(data)*CodeBytePrice/CodeBytePricePrecision + DefaultGasLimit
	tx := RawTransaction{
		Source:   &from,
		Value:    value.Ra(),
		Target:   &to,
		GasLimit: uint64(gaslimit),
		Type:     TransactionTypeTransfer,
		GasPrice: DefaultGasPrice,
		Data:     data,
	}
	return &TransferTransaction{tx}
}

func (tt *TransferTransaction) ToRawTransaction() RawTransaction {
	return tt.RawTransaction
}

type ContractCallTransaction struct {
	RawTransaction
}

func NewContractCallTransaction(from, to Address, value Asset, funcName string, params ...interface{}) *ContractCreateTransaction {
	abi := ABI{FuncName: funcName, Args: params}
	d, _ := json.Marshal(abi)
	tx := RawTransaction{
		Source:   &from,
		Value:    value.Ra(),
		Target:   &to,
		GasLimit: 60000,
		GasPrice: DefaultGasPrice,
		Type:     TransactionTypeContractCall,
		Data:     d,
	}
	return &ContractCreateTransaction{tx}
}

func (tt *ContractCallTransaction) ToRawTransaction() RawTransaction {
	return tt.RawTransaction
}

type ContractCreateTransaction struct {
	RawTransaction
}

func NewContractCreateTransaction(from Address, code string, contractName string, value Asset) *ContractCreateTransaction {
	contract := Contract{Code: string(code), ContractName: contractName, ContractAddress: nil}
	d, _ := json.Marshal(contract)
	tx := RawTransaction{
		Source:   &from,
		Value:    value.Ra(),
		Target:   nil,
		GasLimit: 60000,
		GasPrice: DefaultGasPrice,
		Type:     TransactionTypeContractCreate,
		Data:     d,
	}
	return &ContractCreateTransaction{tx}
}

func (tt *ContractCreateTransaction) ToRawTransaction() RawTransaction {
	return tt.RawTransaction
}

type RewardTransaction struct {
	Hash         Hash   `json:"hash"`
	BlockHash    Hash   `json:"block_hash"`
	GroupSeed    Hash   `json:"group_id"`
	TargetIDs    []ID   `json:"target_ids"`
	Value        uint64 `json:"value"`
	PackFee      uint64 `json:"pack_fee"`
	StatusReport string `json:"status_report"`
	Success      bool   `json:"success"`
}
