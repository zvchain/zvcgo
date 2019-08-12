package zvlib

type Address struct {
}

func NewAddressFromString(s string) (Address, error) {
	return Address{}, nil
}

type Asset struct {
}

func NewAssetFromString(s string) (Asset, error) {
	return Asset{}, nil
}

type Hash struct {
}

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
