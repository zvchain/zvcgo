package zvlib

import "math/big"

const SignLength = 65 //length of signature，32 bytes r & 32 bytes s & 1 byte recid.

type KeyBag struct {
}

func (ws KeyBag) AvailableAddresses() []*Address {
	return nil
}

func (ws KeyBag) AvailablePublicKeys() []*PublicKey {
	return nil
}

func (ws KeyBag) ImportPrivateKey() (err error) {
	return nil
}

func (ws KeyBag) Sign(tx Transaction) (*Sign, error) {
	return newSign(nil), nil
}

type Signer interface {
	Sign(tx RawTransaction) (*Sign, error)
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
