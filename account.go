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

// String serialize private key to hex string
func (key PrivateKey) String() string {
	return Encode(key.pk.D.Bytes())
}

// Bytes serialize private key to bytes
func (key PrivateKey) Bytes() []byte {
	return key.pk.D.Bytes()
}

func (key PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{key.pk.PublicKey}
}

// ImportPrivateKey import private key by bytes of private key
func (key *PrivateKey) ImportPrivateKey(privateKey []byte) (err error) {
	key, err = NewPrivateKey(privateKey)
	return err
}

func NewPrivateKey(privateKey []byte) (*PrivateKey, error) {
	p := &PrivateKey{&ecdsa.PrivateKey{}}
	var one = new(big.Int).SetInt64(1)

	params := secp256k1.S256().Params()
	d := new(big.Int).SetBytes(privateKey)
	if d.Cmp(params.N) >= 0 || d.Cmp(one) < 0 {
		return nil, ErrorInvalidPrivateKey
	}

	p.pk.Curve = secp256k1.S256()
	p.pk.D = d
	p.pk.PublicKey.X, p.pk.PublicKey.Y = p.pk.Curve.ScalarBaseMult(privateKey)
	return p, nil
}

type PublicKey struct {
	key ecdsa.PublicKey
}

// String serialize publicKey key to hex string
func (pk PublicKey) String() string {
	buf := elliptic.Marshal(pk.key.Curve, pk.key.X, pk.key.Y)
	return Encode(buf)
}

func (pk PublicKey) Address() *Address {
	x := pk.key.X.Bytes()
	y := pk.key.Y.Bytes()
	x = append(x, y...)

	addrBuf := sha3.Sum256(x)
	addr, err := NewAddressFromBytes(addrBuf[:])
	if err != nil {
		panic("error address length")
	}
	return &addr
}

type Account struct {
	pk *PrivateKey
}

func NewAccountFromString(key string) (*Account, error) {
	k, err := Decode(key)
	if err != nil {
		return nil, ErrorInvalidPrivateKey
	}
	return NewAccount(k)
}

func NewAccount(privateKey []byte) (*Account, error) {
	key, err := NewPrivateKey(privateKey)
	return &Account{key}, err
}

func (a Account) Sign(tx RawTransaction) (*Sign, error) {
	pribytes := a.pk.Bytes()
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
	return a.PublicKey().Address()
}

func (a Account) PublicKey() *PublicKey {
	return a.pk.PublicKey()
}

func (a Account) PrivateKey() *PrivateKey {
	return a.pk
}
