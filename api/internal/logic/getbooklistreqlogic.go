package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookListReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookListReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookListReqLogic {
	return &GetBookListReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookListReqLogic) GetBookListReq(req *types.GetBookListReq) (resp *types.GetBookListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
