package zvlib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"github.com/zvchain/zvchain/common/secp256k1"
	"github.com/zvchain/zvchain/storage/sha3"
	"math/big"
)

type PrivateKey struct {
	pk *ecdsa.PrivateKey
}

func (key PrivateKey) String() string {
	return Encode(key.pk.D.Bytes())
}

func (key PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{key.pk.D.Bytes()}
}

type PublicKey struct {
	key []byte
}

func (key PublicKey) String() string {
	return Encode(key.key)
}

type Account struct {
	pk *ecdsa.PrivateKey
}

func NewAccountFromString(key string) (*Account, error) {
	k, err := Decode(key)
	if err != nil {
		return nil, ErrorInvalidPrivateKey
	}
	return NewAccount(k)
}

func NewAccount(privateKey []byte) (*Account, error) {
	account := Account{&ecdsa.PrivateKey{}}
	var one = new(big.Int).SetInt64(1)

	params := secp256k1.S256().Params()
	d := new(big.Int).SetBytes(privateKey)
	if d.Cmp(params.N) >= 0 || d.Cmp(one) < 0 {
		return nil, ErrorInvalidPrivateKey
	}

	account.pk.Curve = secp256k1.S256()
	account.pk.D = d
	account.pk.PublicKey.X, account.pk.PublicKey.Y = account.pk.Curve.ScalarBaseMult(privateKey)
	return &account, nil
}

func (a Account) Sign(tx RawTransaction) (*Sign, error) {
	pribytes := a.pk.D.Bytes()
	seckbytes := pribytes
	if len(pribytes) < 32 {
		seckbytes = make([]byte, 32)
		copy(seckbytes[32-len(pribytes):32], pribytes) //make sure that the length of seckey is 32 bytes
	}

	sig, err := secp256k1.Sign(tx.ToRawTransaction().GenHash().Bytes(), seckbytes)
	if err != nil {
		return nil, err
	} else {
		return newSign(sig), nil
	}
}

func (a Account) Address() *Address {
	x := a.pk.PublicKey.X.Bytes()
	y := a.pk.PublicKey.Y.Bytes()
	x = append(x, y...)

	addrBuf := sha3.Sum256(x)
	addr, err := NewAddressFromBytes(addrBuf[:])
	if err != nil {
		panic("error address length")
	}
	return &addr
}

func (a Account) PublicKey() *PublicKey {
	buf := elliptic.Marshal(a.pk.PublicKey.Curve, a.pk.PublicKey.X, a.pk.PublicKey.Y)
	return &PublicKey{buf}
}
