package zvlib

import (
	"math/big"
	"time"
)

type Block struct {
	Height      uint64    `json:"height" mapstructure:"height"`
	Hash        Hash      `json:"hash" mapstructure:"hash"`
	PreHash     Hash      `json:"pre_hash" mapstructure:"pre_hash"`
	CurTime     time.Time `json:"cur_time" mapstructure:"cur_time,squash"`
	PreTime     time.Time `json:"pre_time" mapstructure:"pre_time,squash"`
	Castor      big.Int   `json:"castor" mapstructure:"castor"`
	Group       Hash      `json:"group_id" mapstructure:"group_id"`
	Prove       string    `json:"prove" mapstructure:"prove"`
	TotalQN     uint64    `json:"total_qn" mapstructure:"total_qn"`
	Qn          uint64    `json:"qn" mapstructure:"qn"`
	TxNum       uint64    `json:"txs" mapstructure:"txs"`
	StateRoot   Hash      `json:"state_root" mapstructure:"state_root"`
	TxRoot      Hash      `json:"tx_root" mapstructure:"tx_root"`
	ReceiptRoot Hash      `json:"receipt_root" mapstructure:"receipt_root"`
	ProveRoot   Hash      `json:"prove_root" mapstructure:"prove_root"`
	Random      string    `json:"random" mapstructure:"random"`
}

type BlockDetail struct {
	Block
	GenRewardTx   RewardTransaction   `json:"gen_reward_tx"`
	Trans         []Transaction       `json:"trans"`
	BodyRewardTxs []RewardTransaction `json:"body_reward_txs"`
	PreTotalQN    uint64              `json:"pre_total_qn"`
}
