package zvlib

import (
	"crypto/ecdsa"
	"github.com/zvchain/zvchain/common/secp256k1"
)

type PrivateKey struct {
}

type PublicKey struct {
}

type Account struct {
	pk *ecdsa.PrivateKey
}

func NewAccount(privateKey []byte) *Account {
	return nil
}

func (a Account) Sign(tx RawTransaction) (*Sign, error) {
	pribytes := a.pk.D.Bytes()
	seckbytes := pribytes
	if len(pribytes) < 32 {
		seckbytes = make([]byte, 32)
		copy(seckbytes[32-len(pribytes):32], pribytes) //make sure that the length of seckey is 32 bytes
	}

	sig, err := secp256k1.Sign(tx.GenHash().Bytes(), seckbytes)
	if err != nil {
		return nil, err
	} else {
		return newSign(sig), nil
	}
}
