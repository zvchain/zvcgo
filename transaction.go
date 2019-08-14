package zvlib

import (
	"bytes"
	"github.com/zvchain/zvchain/common"
	"github.com/zvchain/zvchain/middleware/types"
)

type Transaction struct {
	Data   []byte
	Value  uint64
	Nonce  uint64
	Target *Address
	Type   int8

	GasLimit uint64
	GasPrice uint64
	Hash     Hash

	ExtraData []byte
	Sign      []byte
	Source    *Address
}

func (t Transaction) ToRawTransaction() Transaction {
	return t
}

func (t Transaction) GenHash() Hash {
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
		buffer.Write(t.ExtraData)
	}
	return Hash{common.Sha256(buffer.Bytes())}
}

type RawTransaction interface {
	ToRawTransaction() Transaction
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
	TargetIDs    string `json:"target_ids"`
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
