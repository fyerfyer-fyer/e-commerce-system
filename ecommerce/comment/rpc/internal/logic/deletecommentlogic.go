package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/comment"
	"github.com/fyerfyer/e-commerce-system/ecommerce/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.Empty, error) {
	// todo: add your logic here and delete this line

	return &comment.Empty{}, nil
}
