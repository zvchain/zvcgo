package zvcgo

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

var defaultRootDerivationPath = []uint32{0x80000000 + 44, 0x80000000 + 60, 0x80000000 + 0, 0}

func parseDerivationPath(path string) ([]uint32, error) {
	var result []uint32

	// Handle absolute or relative paths
	components := strings.Split(path, "/")
	switch {
	case len(components) == 0:
		return nil, errors.New("empty derivation path")

	case strings.TrimSpace(components[0]) == "":
		return nil, errors.New("ambiguous path: use 'm/' prefix for absolute paths, or no leading '/' for relative ones")

	case strings.TrimSpace(components[0]) == "m":
		components = components[1:]

	default:
		result = append(result, defaultRootDerivationPath...)
	}
	// All remaining components are relative, append one by one
	if len(components) == 0 {
		return nil, errors.New("empty derivation path") // Empty relative paths
	}
	for _, component := range components {
		// Ignore any user added whitespace
		component = strings.TrimSpace(component)
		var value uint32

		// Handle hardened paths
		if strings.HasSuffix(component, "'") {
			value = 0x80000000
			component = strings.TrimSpace(strings.TrimSuffix(component, "'"))
		}
		// Handle the non hardened component
		bigval, ok := new(big.Int).SetString(component, 0)
		if !ok {
			return nil, fmt.Errorf("invalid component: %s", component)
		}
		max := math.MaxUint32 - value
		if bigval.Sign() < 0 || bigval.Cmp(big.NewInt(int64(max))) > 0 {
			if value == 0 {
				return nil, fmt.Errorf("component %v out of allowed range [0, %d]", bigval, max)
			}
			return nil, fmt.Errorf("component %v out of allowed hardened range [0, %d]", bigval, max)
		}
		value += uint32(bigval.Uint64())

		// Append and repeat
		result = append(result, value)
	}
	return result, nil
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func hasZVPrefix(input string) bool {
	return len(input) >= 2 && (input[0] == 'z' || input[0] == 'Z') && (input[1] == 'v' || input[1] == 'V')
}

// Decode decodes a hex string with 0x prefix.
func Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrorInvalid0xString
	}
	if !has0xPrefix(input) {
		return nil, ErrorInvalid0xString
	}
	b, err := hex.DecodeString(input[2:])
	if err != nil {
		err = ErrorInvalid0xString
	}
	return b, err
}

// Encode encodes b as a hex string with 0x prefix.
func Encode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

// Decode decodes a hex string with zv prefix.
func DecodeZV(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrorInvalid0xString
	}
	if !hasZVPrefix(input) {
		return nil, ErrorInvalid0xString
	}
	b, err := hex.DecodeString(input[2:])
	if err != nil {
		err = ErrorInvalid0xString
	}
	return b, err
}
