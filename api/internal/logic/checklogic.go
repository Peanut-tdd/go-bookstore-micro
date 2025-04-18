package logic

import (
	"bookstore/rpc/check/check"
	"context"
	"fmt"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	// todo: add your logic here and delete this line

	resp1, err := l.svcCtx.Check.Check(l.ctx, &check.CheckReq{
		Book: req.Book,
	})

	if err != nil {
		fmt.Println("err:", err)
	}

	resp = &types.CheckResp{
		Found: resp1.Found,
		Price: resp1.Price,
	}

	return
}
