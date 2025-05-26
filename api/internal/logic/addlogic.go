package logic

import (
	"bookstore/rpc/add/add"
	"context"
	"fmt"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 逻辑层
func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	// todo: add your logic here and delete this line

	//调用微服务
	resp1, err := l.svcCtx.Adder.Add(l.ctx, &add.AddReq{
		Book:  req.Book.Name,
		Price: req.Book.Price,
	})

	fmt.Printf("resp:%+v, err:%+v", resp1, err)

	if resp1 == nil || err != nil {
		resp = &types.AddResp{}
	}

	resp = &types.AddResp{
		Ok: resp1.Ok,
	}
	return
}
