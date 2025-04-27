package logic

import (
	"context"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookListLogic {
	return &GetBookListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookListLogic) GetBookList(in *add.GetBookListReq) (*add.GetBookListResp, error) {
	// todo: add your logic here and delete this line

	return &add.GetBookListResp{}, nil
}
