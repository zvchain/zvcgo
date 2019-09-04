ZVChain API library for Go
=========================


[![GoDoc](https://godoc.org/github.com/zvchain/zvcgo?status.svg)](https://godoc.org/github.com/zvchain/zvcgo)

This library provides simple access to data structures and API calls to an ZVChain RPC server.  

## Install

```
go get -u github.com/zvchain/zvcgo
```

## Basic usage

```go
api := zvcgo.NewApi("http://node1.zvchain.io")

height, _ := api.BlockHeight()
fmt.Printf("Current block height: %d", height)
```

## Transfer ZVC Example 
```go
package main

import (
	"fmt"
	"github.com/zvchain/zvcgo"
)

func main() {
	mnemonic, _ := zvcgo.NewMnemonic(zvcgo.Mnemonic12WordBitSize)
	fmt.Println("Mnemonic: ", mnemonic)
	wallet := zvcgo.NewWallet(mnemonic)
	if wallet == nil {
		panic("wrong mnemonic")
	}
	acc, _ := wallet.DeriveAccount(0)
	// acc, err := zvcgo.NewAccountFromString("you private key")
	fmt.Println("account address: ", acc.Address().String())
	api := zvcgo.NewApi("endpoint host")
	api.SetSigner(acc)
	asset, _ := zvcgo.NewAssetFromString("123 zvc")
	target, _ := zvcgo.NewAddressFromString("you address")
	tx := zvcgo.NewTransferTransaction(*acc.Address(), target, asset, []byte{})
	nonce, _ := api.GetNonce(*acc.Address())
	tx.SetNonce(nonce)
	hash, err := api.SignAndSendTransaction(tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction hash: ", hash.String())
}


```



## License

GPL
