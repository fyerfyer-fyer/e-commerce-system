package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRepliesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepliesLogic {
	return &GetRepliesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRepliesLogic) GetReplies(in *reply.GetRepliesRequest) (*reply.GetRepliesResponse, error) {
	// todo: add your logic here and delete this line

	return &reply.GetRepliesResponse{}, nil
}
