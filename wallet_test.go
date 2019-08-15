package zvlib

import (
	"fmt"
	"testing"
)

func TestNewWallet(t *testing.T) {
	mnemonic := "mixed example surround name fat cigar hidden text arrest advice muscle diamond"
	wallet := NewWallet(mnemonic)
	account1, _ := wallet.DeriveAccount(0)
	if account1.Address().String() != "zv7168515e5868df9bc31d922e984b8d8bdc6099829e32b5e947e70adf2760285c" {
		fmt.Println(account1.Address())
	}
}
