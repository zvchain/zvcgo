package zvlib

import "errors"

var (
	NoSignerError = errors.New("signer should be set before send a no-sign transaction")
)

type Api struct {
	host string
}

func NewApi(url string) *Api {
	return nil
}

func (api Api) GetBlockByHeight(height uint64) (*Block, error) {
	return nil, nil
}

func (api Api) GetBlockDetailByHeight(height uint64) (*BlockDetail, error) {
	return nil, nil
}

func (api Api) GetBlockByHash(hash Hash) (*Block, error) {
	return nil, nil
}

func (api Api) GetBlockDetailByHash(hash Hash) (*BlockDetail, error) {
	return nil, nil
}

func (api Api) GetTransactionByHash(hash Hash) (*Transaction, error) {
	return nil, nil
}

func (api Api) BlockHeight() (uint64, error) {
	api.request("blockHeight")
	return 0, nil
}

func (api Api) GetBlockHash() (*Hash, error) {
	return nil, nil
}

func (api Api) GetNonce(address Address) (uint64, error) {
	return 0, nil
}

func (api Api) GetCode(address Address) (uint64, error) {
	return 0, nil
}

func (api Api) GetData(address Address, key string) (interface{}, error) {
	return 0, nil
}

func (api Api) SetSigner(signer Signer) {

}

func (api Api) SendRawTransaction(tx RawTransaction) (*Hash, error) {
	return nil, nil
}

func (api Api) SignAndSendRawTransaction(tx RawTransaction) (*Hash, error) {
	return nil, nil
}

func (api Api) GetPastEvent(address Address, topic string, from, to uint64) ([]*Event, error) {
	return nil, nil
}

func (api Api) EventListen(address Address, topic string, from uint64, callBack EventCallBack) error {
	return nil
}

func (api Api) MinerInfo(address Address) error {
	return nil
}

func (api Api) Balance(address Address) error {
	return nil
}

type EventCallBack func(event *Event)

func A(event *Event) {

}
