package zvlib

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

const ZVCPath = "m/44'/372'/0'/0/%d" // 372

type Wallet struct {
	key *hdkeychain.ExtendedKey
}

func NewWallet(mnemonic string) *Wallet {
	if mnemonic == "" {
		return nil
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil
	}
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil
	}
	return &Wallet{masterKey}
}

func (m Wallet) DeriveAccount(index int) (*Account, error) {
	path, err := parseDerivationPath(fmt.Sprintf(ZVCPath, index))
	if err != nil {
		return nil, err
	}
	key := m.key
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKeyECDSA := privateKey.ToECDSA()
	return NewAccount(privateKeyECDSA.D.Bytes()), nil
}
