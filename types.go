package zvlib

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	InvalidAssetString = errors.New("invalid asset string")
)

const (
	Ra  = 1
	mRa = 1000
	kRa = 1000000
	ZVC = 1000000000
)

const (
	AddressLength = 32 //Length of Address( golang.SHA3，256-bit)
	HashLength    = 32 //Length of Hash (golang.SHA3, 256-bit)。
)

//type Address struct {
//}
// Address data struct
type Address [AddressLength]byte

func NewAddressFromString(s string) (Address, error) {
	return Address{}, nil
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
				return Asset{}, InvalidAssetString
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
		return Asset{}, InvalidAssetString
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
		return Asset{}, InvalidAssetString
	}
	return asset, nil
}

func (a Asset) ZVC() uint64 {
	return 0
}

func (a Asset) KRa() uint64 {
	return 0
}

func (a Asset) MRa() uint64 {
	return 0
}

func (a Asset) Ra() uint64 {
	return 0
}

//type Hash struct {
//}

// Hash data struct (256-bits)
type Hash [HashLength]byte

func (h Hash) Bytes() []byte {
	return nil
}

type Event struct {
	Deleted bool
	Height  uint64
	TxHash  Hash
	Index   uint64
	Message map[string]interface{}
	Default []interface{}
}
