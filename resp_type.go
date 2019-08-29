package zvlib

import (
	"math/big"
)

// ABIVerify stores the contract function name and args types,
// in order to facilitate the abi verify
type ABIVerify struct {
	FuncName string
	Args     []string
}

type AccountMsg struct {
	Balance   *big.Int               `json:"balance"`
	Nonce     uint64                 `json:"nonce"`
	Type      uint32                 `json:"type"`
	CodeHash  string                 `json:"code_hash"`
	ABI       []ABIVerify            `json:"abi"`
	Code      string                 `json:"code"`
	StateData map[string]interface{} `json:"state_data"`
}

type MortGage struct {
	Stake              uint64 `json:"stake"`
	ApplyHeight        uint64 `json:"apply_height"`
	Type               string `json:"type"`
	Status             string `json:"miner_status"`
	StatusUpdateHeight uint64 `json:"status_update_height"`
}

type StakeDetail struct {
	Value        uint64 `json:"value"`
	UpdateHeight uint64 `json:"update_height"`
	MType        string `json:"m_type"`
	Status       string `json:"stake_status"`
}

type MinerStakeDetails struct {
	Overview []*MortGage               `json:"overview,omitempty"`
	Details  map[string][]*StakeDetail `json:"details,omitempty"`
}
