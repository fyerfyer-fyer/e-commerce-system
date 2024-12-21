package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReplyLogic {
	return &GetReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReplyLogic) GetReply(in *reply.GetReplyRequest) (*reply.Reply, error) {
	// todo: add your logic here and delete this line

	return &reply.Reply{}, nil
}
