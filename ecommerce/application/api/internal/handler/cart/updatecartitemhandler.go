package cart

import (
	"net/http"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/logic/cart"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新购物车商品
func UpdateCartItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewUpdateCartItemLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCartItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
