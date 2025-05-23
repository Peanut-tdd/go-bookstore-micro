package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Adder  adder.Adder // rpc client
	Check  checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder:  adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Check:  checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
