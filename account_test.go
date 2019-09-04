package zvcgo

import (
	"testing"
)

var privKey = "0xf32d35d2b40da6582aa6a54bfd248a39be2eb3620d33100bc5082f2ab5fbcb6d"
var pubKey = "0x04f5f322e57871aca6e0a940083f06a1de136919fd093937b22eeccf648ffdf8eac86eb0ca19e531f9473a37bce33005f8ca5f75fea6a949a13a3f3ddbc153887a"
var addr = "zvfd2f4785793a506225773ba7b036bfd55bb593c9dd16eaf52147a4ec963867af"

func TestAccount(t *testing.T) {
	account, err := NewAccountFromString(privKey)
	if Encode(account.PrivateKey().Bytes()) != privKey {
		t.Error("import key error")
	}
	if err != nil {
		t.Error(err)
	}
	publicKey := account.PrivateKey().PublicKey()
	if publicKey.String() != pubKey {
		t.Errorf("wanted: %s, got: %s", pubKey, publicKey)
	}
	address := publicKey.Address()
	if address.String() != addr {
		t.Errorf("wanted: %s, got: %s", addr, address)
	}
}
