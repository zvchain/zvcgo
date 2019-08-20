package zvlib

import (
	"math/big"
)

const SignLength = 65 //length of signature，32 bytes r & 32 bytes s & 1 byte recid.

type KeyBag struct {
	keys map[string]*Account
}

func (kb KeyBag) AvailableAddresses() []*Address {
	pks := make([]*Address, 0)
	for _, value := range kb.keys {
		pks = append(pks, value.Address())
	}
	return pks
}

func (kb KeyBag) AvailablePublicKeys() []*PublicKey {
	pks := make([]*PublicKey, 0)
	for _, value := range kb.keys {
		pks = append(pks, value.PublicKey())
	}
	return pks
}

func (kb *KeyBag) ImportPrivateKey(key []byte) (err error) {
	account, err := NewAccount(key)
	if err != nil {
		return err
	}
	kb.keys[account.Address().String()] = account
	return nil
}

func (kb *KeyBag) ImportPrivateKeyFromString(k string) (err error) {
	bs, err := Decode(k)
	if err != nil {
		return err
	}
	return kb.ImportPrivateKey(bs)
}

func (kb KeyBag) Sign(tx Transaction) (*Sign, error) {
	txr := tx.ToRawTransaction()
	if txr.Source == nil {
		return nil, ErrorSourceEmpty
	}
	account := kb.keys[txr.Source.String()]
	if account == nil {
		return nil, ErrorSignerNotFound
	}
	return account.Sign(tx)
}

func NewKeyBag() *KeyBag {
	return &KeyBag{make(map[string]*Account)}
}

type Signer interface {
	Sign(tx Transaction) (*Sign, error)
}

// Sign Data struct
type Sign struct {
	r     big.Int
	s     big.Int
	recid byte
}

func newSign(b []byte) *Sign {
	if len(b) == 65 {
		var r, s big.Int
		br := b[:32]
		r = *r.SetBytes(br)

		sr := b[32:64]
		s = *s.SetBytes(sr)

		recid := b[64]
		return &Sign{r, s, recid}
	}
	return nil
}

func (s Sign) Bytes() []byte {
	rb := s.r.Bytes()
	sb := s.s.Bytes()
	r := make([]byte, SignLength)
	copy(r[32-len(rb):32], rb)
	copy(r[64-len(sb):64], sb)
	r[64] = s.recid
	return r
}
