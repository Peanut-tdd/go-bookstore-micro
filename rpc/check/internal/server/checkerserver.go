// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: check.proto

package server

import (
	"context"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/logic"
	"bookstore/rpc/check/internal/svc"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
	check.UnimplementedCheckerServer
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *check.CheckReq) (*check.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
