package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReplyLogic {
	return &AddReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddReplyLogic) AddReply(in *reply.AddReplyRequest) (*reply.Empty, error) {
	// todo: add your logic here and delete this line

	return &reply.Empty{}, nil
}
