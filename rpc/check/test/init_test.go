package test

import (
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/check/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"strconv"
	"testing"
)

var configFile = flag.String("f", "../etc/check.yaml", "the config file")

func TestConn(t *testing.T) {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	err := ctx.RedisClient.SetCtx(context.Background(), "age", strconv.Itoa(10))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ctx.RedisClient.Get("age"))

	fmt.Println(111)

}
