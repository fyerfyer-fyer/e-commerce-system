package order

import (
	"net/http"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/logic/order"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取订单详情
func GetOrderDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Empty
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewGetOrderDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
