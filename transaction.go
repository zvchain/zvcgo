package zvlib

import (
	"bytes"
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/middleware/types"
)

const (
	DefaultGasPrice = 500
	DefaultGasLimit = 3000
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
	Sign      []byte   `json:"sign"`
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

func NewTrasnferTransaction(target, value string) (*RawTransaction, error) {

	if !hasZVPrefix(target) {
		return nil, ErrorInvalidZVString
	}

	asset, err := NewAssetFromString(value)
	if err != nil {
		return nil, err
	}

	return &RawTransaction{
		Value:    asset.value,
		GasPrice: DefaultGasPrice,
		GasLimit: DefaultGasLimit,
	}, nil
}

func (t RawTransaction) GenHash() Hash {
	buffer := bytes.Buffer{}
	if len(t.Data) != 0 {
		buffer.Write(t.Data)
	}
	valueB := types.NewBigInt(t.Value)
	// value
	buffer.Write(valueB.GetBytesWithSign())
	// nonce
	buffer.Write(common.Uint64ToByte(t.Nonce))
	if t.Target != nil {
		buffer.Write(t.Target.Bytes())
	}
	// type
	buffer.WriteByte(byte(t.Type))
	gasLimitB := types.NewBigInt(t.GasLimit)
	gasPriceB := types.NewBigInt(t.GasPrice)
	// gasLimit
	buffer.Write(gasLimitB.GetBytesWithSign())
	// gasPrice
	buffer.Write(gasPriceB.GetBytesWithSign())
	if len(t.ExtraData) != 0 {
		buffer.Write([]byte(t.ExtraData))
	}
	return Hash{common.Sha256(buffer.Bytes())}
}

type Transaction interface {
	ToRawTransaction() RawTransaction
}

type TransferTransaction struct {
	From      Address
	To        Address
	Value     Asset
	ExtraData []byte
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

type ContractCallTransaction struct {
}

type ContractDeployTransaction struct {
}

type MinerStakeTransaction struct {
}
