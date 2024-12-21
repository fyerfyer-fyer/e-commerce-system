package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsLogic) GetComments(in *comment.GetCommentsRequest) (*comment.GetCommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetCommentsResponse{}, nil
}
