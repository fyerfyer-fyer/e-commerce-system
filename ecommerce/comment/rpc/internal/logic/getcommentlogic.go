package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *comment.GetCommentRequest) (*comment.Comment, error) {
	// todo: add your logic here and delete this line

	return &comment.Comment{}, nil
}
