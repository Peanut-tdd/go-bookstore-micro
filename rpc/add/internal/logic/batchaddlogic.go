package logic

import (
	"context"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddLogic {
	return &BatchAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchAddLogic) BatchAdd(in *add.BatchAddReq) (*add.BatchAddResp, error) {
	// todo: add your logic here and delete this line

	return &add.BatchAddResp{}, nil
}
