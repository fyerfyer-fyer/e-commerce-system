package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditReplyLogic {
	return &EditReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditReplyLogic) EditReply(in *reply.EditReplyRequest) (*reply.Empty, error) {
	// todo: add your logic here and delete this line

	return &reply.Empty{}, nil
}
