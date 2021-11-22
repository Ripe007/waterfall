package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync/atomic"
)

type RPCClientRoundRobin struct {
	clis   []*ethclient.Client
	robin  int64
	length int64
}

func NewRPCClientRoundRobin(urls []string) (*RPCClientRoundRobin, error) {
	rpcs := &RPCClientRoundRobin{
		robin:  0,
		length: int64(len(urls)),
	}
	for _, u := range urls {
		cli, err := ethclient.Dial(u)
		if err != nil {
			return nil, fmt.Errorf("[ethclient]%e", err)
		}
		rpcs.clis = append(rpcs.clis, cli)

	}
	return rpcs, nil
}
func (r *RPCClientRoundRobin) Next() *ethclient.Client {
	if r.length == 0 {
		return nil
	}
	robin := atomic.LoadInt64(&r.robin)
	nrobin := (robin + 1) % r.length
	for !atomic.CompareAndSwapInt64(&r.robin, robin, nrobin) {

	}
	return r.clis[nrobin]
}
func (r *RPCClientRoundRobin) Close() {
	for i := range r.clis {
		r.clis[i].Close()
	}
}
