package rpc

import "github.com/btcsuite/btcd/btcjson"

type IRpc interface {
	NewAddress() (string, string)
	ListUnspent() []btcjson.ListUnspentResult
	Address2BalanceMap() map[string]float64
	Balance() float64
	ListUnspentByAddress(_address string) []btcjson.ListUnspentResult
	NewTx(_from []string, _to map[string]float64, _changeAddress string) string
	SignTx(_tx string) (string, bool)
	SendTx(_tx string) string
	Transfer(_from []string, _to map[string]float64, _changeAddress string) string
	ValidateAddress(_address string) bool
	Call(method string, params map[string]interface{}) interface{}
}
