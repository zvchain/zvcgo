package zvlib

import "math/big"

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
