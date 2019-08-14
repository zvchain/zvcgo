package zvlib

import (
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorInvalidAssetString   = errors.New("invalid asset string")
	ErrorInvalidAddressString = errors.New("invalid address string")
	ErrorInvalidPrivateKey    = errors.New("invalid private key")
	ErrorInvalid0xString      = errors.New("invalid hex String")
	ErrorSignerNotFound       = errors.New("signer not found")
	ErrorSourceEmpty          = errors.New("source should not be empty")
)

const AddrPrefix = "zv"
const AddressLength = 32

const (
	Ra  = 1
	kRa = 1000
	mRa = 1000000
	ZVC = 1000000000
)

type Address struct {
	data []byte
}

var addrReg = regexp.MustCompile("^[Zz][Vv][0-9a-fA-F]{64}$")

func ValidateAddress(addr string) bool {
	return addrReg.MatchString(addr)
}

func NewAddressFromString(s string) (addr Address, err error) {
	if !ValidateAddress(s) {
		return Address{}, ErrorInvalidAddressString
	}
	addr.data, err = hex.DecodeString(s[2:])
	return addr, err
}

func NewAddressFromBytes(bs []byte) (addr Address, err error) {
	if len(bs) != AddressLength {
		return addr, ErrorInvalidAddressString
	}
	addr.data = bs
	return addr, err
}

func (a Address) String() string {
	return AddrPrefix + hex.EncodeToString(a.data)
}

func (a Address) Bytes() []byte {
	return a.data
}

type Asset struct {
	value uint64
}

func NewAssetFromString(s string) (Asset, error) {
	asset := Asset{}
	var valuePart, symbol []byte
	reader := strings.NewReader(s)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if b == ' ' {
			continue
		} else if b == '.' || (b >= '0' && b <= '9') {
			if len(symbol) != 0 {
				return Asset{}, ErrorInvalidAssetString
			}
			valuePart = append(valuePart, b)
		} else if b >= 'a' && b <= 'z' {
			symbol = append(symbol, b)
		} else if b >= 'A' && b <= 'Z' {
			symbol = append(symbol, b+32)
		}
	}
	value, err := strconv.ParseFloat(string(valuePart), 64)
	if err != nil {
		return Asset{}, ErrorInvalidAssetString
	}
	switch string(symbol) {
	case "ra":
		asset.value = uint64(value * Ra)
	case "mra":
		asset.value = uint64(value * mRa)
	case "kra":
		asset.value = uint64(value * kRa)
	case "zvc":
		asset.value = uint64(value * ZVC)
	default:
		return Asset{}, ErrorInvalidAssetString
	}
	return asset, nil
}

func (a Asset) ZVC() uint64 {
	return a.value
}

func (a Asset) KRa() uint64 {
	return a.value / kRa
}

func (a Asset) MRa() uint64 {
	return a.value / mRa
}

func (a Asset) Ra() uint64 {
	return a.value
}

type Hash struct {
	data []byte
}

func (h Hash) Bytes() []byte {
	return h.data
}

type Event struct {
	Deleted bool
	Height  uint64
	TxHash  Hash
	Index   uint64
	Message map[string]interface{}
	Default []interface{}
}
