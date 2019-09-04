package zvcgo

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

const ZVCPath = "m/44'/372'/0'/0/%d" // 372
const Mnemonic12WordBitSize = 128
const Mnemonic24WordBitSize = 256

type Wallet struct {
	KeyBag
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
	return &Wallet{*NewKeyBag(), masterKey}
}

// DeriveAccount  derive account by index
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
	return NewAccount(privateKeyECDSA.D.Bytes())
}

// NewMnemonic Generate Mnemonic, returns of number of words depends on input: bitSize(eg: Mnemonic12WordBitSize, Mnemonic24WordBitSize)
func NewMnemonic(bitSize int) (string, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}
