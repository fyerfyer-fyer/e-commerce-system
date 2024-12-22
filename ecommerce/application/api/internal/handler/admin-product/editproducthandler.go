package admin-product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/logic/admin-product"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

)

// 编辑商品
func EditProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin-product.NewEditProductLogic(r.Context(), svcCtx)
		resp, err := l.EditProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
