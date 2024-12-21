package comment

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑评论
func NewEditCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCommentLogic {
	return &EditCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditCommentLogic) EditComment(req *types.EditCommentReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
