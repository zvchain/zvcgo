package zvlib

import (
	"bytes"
	"encoding/json"
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

// Result is rpc request successfully returns the variable parameter
type Result struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// RPCReqObj is complete rpc request body
type RPCReqObj struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Jsonrpc string        `json:"jsonrpc"`
	ID      uint          `json:"id"`
}

// RPCResObj is complete rpc response body
type RPCResObj struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      uint         `json:"id"`
	Result  *Result      `json:"result,omitempty"`
	Error   *ErrorResult `json:"error,omitempty"`
}

// ErrorResult is rpc request error returned variable parameter
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

func (a *Api) request(method string, params ...interface{}) *Result {
	if a.host == "" {
		return opError(ErrUnConnected)
	}

	param := RPCReqObj{
		Method:  "Gzv_" + method,
		Params:  params[:],
		ID:      1,
		Jsonrpc: "2.0",
	}

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return opError(err)
	}

	resp, err := http.Post(a.host, "application/json", bytes.NewReader(paramBytes))
	if err != nil {
		return opError(err)
	}
	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	ret := &RPCResObj{}
	if err := json.Unmarshal(responseBytes, ret); err != nil {
		return opError(err)
	}
	if ret.Error != nil {
		return opError(fmt.Errorf(ret.Error.Message))
	}
	return ret.Result
}
