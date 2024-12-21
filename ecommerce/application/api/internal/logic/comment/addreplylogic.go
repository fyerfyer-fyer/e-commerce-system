package comment

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加回复
func NewAddReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReplyLogic {
	return &AddReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReplyLogic) AddReply(req *types.AddReplyReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
