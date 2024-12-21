package comment

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑回复
func NewEditReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditReplyLogic {
	return &EditReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditReplyLogic) EditReply(req *types.EditReplyReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
