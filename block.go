package zvlib

import (
	"math/big"
	"time"
)

type Block struct {
	Height      uint64    `json:"height"`
	Hash        Hash      `json:"hash"`
	PreHash     Hash      `json:"pre_hash"`
	CurTime     time.Time `json:"cur_time"`
	PreTime     time.Time `json:"pre_time"`
	Castor      big.Int   `json:"castor"`
	Group       Hash      `json:"group_id"`
	Prove       string    `json:"prove"`
	TotalQN     uint64    `json:"total_qn"`
	Qn          uint64    `json:"qn"`
	TxNum       uint64    `json:"txs"`
	StateRoot   Hash      `json:"state_root"`
	TxRoot      Hash      `json:"tx_root"`
	ReceiptRoot Hash      `json:"receipt_root"`
	ProveRoot   Hash      `json:"prove_root"`
	Random      string    `json:"random"`
}

type BlockDetail struct {
	Block
	GenRewardTx   RewardTransaction   `json:"gen_reward_tx"`
	Trans         []Transaction       `json:"trans"`
	BodyRewardTxs []RewardTransaction `json:"body_reward_txs"`
	PreTotalQN    uint64              `json:"pre_total_qn"`
}
