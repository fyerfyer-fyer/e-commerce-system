package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCommentLogic {
	return &EditCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditCommentLogic) EditComment(in *comment.EditCommentRequest) (*comment.Empty, error) {
	// todo: add your logic here and delete this line

	return &comment.Empty{}, nil
}
