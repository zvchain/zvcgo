package zvlib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrPassword    = fmt.Errorf("password error")
	ErrUnlocked    = fmt.Errorf("please unlock the Account first")
	ErrUnConnected = fmt.Errorf("please connect to one node first")
	ErrInternal    = fmt.Errorf("internal error")
)

// Result is rpc requestGzv successfully returns the variable parameter
type Result struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// RPCReqObj is complete rpc requestGzv body
type RPCReqObj struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Jsonrpc string        `json:"jsonrpc"`
	ID      uint          `json:"id"`
}

type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

type RPCResObj struct {
	Result RawMessage   `json:"result,omitempty"`
	Error  *ErrorResult `json:"error,omitempty"`
}

// ErrorResult is rpc requestGzv error returned variable parameter
type ErrorResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func opError(err error) *Result {
	ret, _ := failResult(err.Error())
	return ret
}

func opSuccess(data interface{}) *Result {
	ret, _ := successResult(data)
	return ret
}

func successResult(data interface{}) (*Result, error) {
	return &Result{
		Message: "success",
		Data:    data,
		Status:  0,
	}, nil
}

func failResult(err string) (*Result, error) {
	return &Result{
		Message: err,
		Data:    nil,
		Status:  -1,
	}, nil
}

func (api *Api) request(nameSpace, method string, params ...interface{}) (*RawMessage, error) {
	if api.host == "" {
		return nil, fmt.Errorf("ErrUnConnected")
	}

	param := RPCReqObj{
		Method:  nameSpace + "_" + method,
		Params:  params[:],
		ID:      1,
		Jsonrpc: "2.0",
	}

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(api.host, "application/json", bytes.NewReader(paramBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	ret := &RPCResObj{}
	fmt.Println(string(responseBytes))
	if err := json.Unmarshal(responseBytes, ret); err != nil {
		return nil, err
	}
	if ret.Error != nil {
		return nil, fmt.Errorf(ret.Error.Message)
	}
	return &ret.Result, nil
}
