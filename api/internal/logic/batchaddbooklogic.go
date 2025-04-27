package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchAddBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddBookLogic {
	return &BatchAddBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchAddBookLogic) BatchAddBook(req *types.BatchAddReq) (resp *types.BatchAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
