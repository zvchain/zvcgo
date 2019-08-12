package zvlib

type Transaction struct {
	Data   []byte
	Value  uint64
	Nonce  uint64
	Target Address
	Type   int8

	GasLimit uint64
	GasPrice uint64
	Hash     Hash

	ExtraData []byte
	Sign      []byte
	Source    Address
}

type RawTransaction interface {
	ToRawTransaction() Transaction
	GenHash() Hash
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
