ZVChain API library for Go
=========================

[点击查看中文版](./README-cn.md)

[![GoDoc](https://godoc.org/github.com/zvchain/zvlib?status.svg)](https://godoc.org/github.com/zvchain/zvlib)

This library provides simple access to data structures and API calls to an ZVChain RPC server.  

## Install

```
go get -u github.com/zvchain/zvlib
```

## Basic usage

```go
api := zvlib.NewApi("http://node1.zvchain.io")

height, _ := api.BlockHeight()
fmt.Printf("Current block height: %d", height)
```

## Example

### Reference


## License

GPL
