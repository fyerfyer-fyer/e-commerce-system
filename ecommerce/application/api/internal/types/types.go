// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddAddressReq struct {
	UserId    int64  `json:"userId"`          // 用户ID
	Line1     string `json:"line1"`           // 地址行1
	Line2     string `json:"line2,omitempty"` // 地址行2
	City      string `json:"city"`            // 城市
	State     string `json:"state"`           // 省份/州
	Zip       string `json:"zip"`             // 邮政编码
	Country   string `json:"country"`         // 国家
	IsDefault bool   `json:"isDefault"`       // 是否默认地址
}

type AddCommentReq struct {
	ProductId int64  `json:"productId"` // 商品ID
	UserId    int64  `json:"userId"`    // 用户ID
	Rating    int32  `json:"rating"`    // 评分 (1-5)
	Content   string `json:"content"`   // 评论内容
}

type AddProductReq struct {
	Name          string      `json:"name"`                    // 商品名称
	Price         float64     `json:"price"`                   // 商品价格
	OriginalPrice float64     `json:"originalPrice,omitempty"` // 原价
	Description   string      `json:"description"`             // 商品描述
	ImageUrls     []string    `json:"imageUrls"`               // 商品图片列表
	Stock         int         `json:"stock"`                   // 库存数量
	CategoryId    int64       `json:"categoryId"`              // 分类ID
	Attributes    []Attribute `json:"attributes"`              // 商品属性
}

type AddReplyReq struct {
	CommentId int64  `json:"commentId"` // 评论ID
	UserId    int64  `json:"userId"`    // 用户ID
	Content   string `json:"content"`   // 回复内容
}

type AddToCartReq struct {
	UserId    int64 `json:"userId"`    // 用户ID
	ProductId int64 `json:"productId"` // 商品ID
	Quantity  int   `json:"quantity"`  // 购买数量
}

type Address struct {
	Id        int64  `json:"id"`              // 地址ID
	Line1     string `json:"line1"`           // 地址行1
	Line2     string `json:"line2,omitempty"` // 地址行2
	City      string `json:"city"`            // 城市
	State     string `json:"state"`           // 省份/州
	Zip       string `json:"zip"`             // 邮政编码
	Country   string `json:"country"`         // 国家
	IsDefault bool   `json:"isDefault"`       // 是否默认
}

type AddressListResp struct {
	Addresses []Address `json:"addresses"` // 地址列表
}

type Attribute struct {
	Name  string `json:"name"`  // 属性名
	Value string `json:"value"` // 属性值
}

type BaseResponse struct {
	Code int    `json:"code"` // Response code
	Msg  string `json:"msg"`  // Response message
}

type CartItemResp struct {
	ProductId   int64   `json:"productId"`   // 商品ID
	Name        string  `json:"name"`        // 商品名称
	ImageUrl    string  `json:"imageUrl"`    // 商品主图URL
	Price       float64 `json:"price"`       // 单价
	Quantity    int     `json:"quantity"`    // 数量
	Selected    bool    `json:"selected"`    // 是否选中
	TotalAmount float64 `json:"totalAmount"` // 商品小计
}

type CartListResp struct {
	Items       []CartItemResp `json:"items"`       // 购物车商品列表
	TotalAmount float64        `json:"totalAmount"` // 总金额
}

type ClearCartReq struct {
	UserId int64 `json:"userId"` // 用户ID
}

type CommentListResp struct {
	Comments []CommentResp `json:"comments"` // 评论列表
	Total    int           `json:"total"`    // 评论总数
}

type CommentResp struct {
	Id        int64  `json:"id"`        // 评论ID
	ProductId int64  `json:"productId"` // 商品ID
	UserId    int64  `json:"userId"`    // 用户ID
	Username  string `json:"username"`  // 评论用户的名称
	Rating    int32  `json:"rating"`    // 评分 (1-5)
	Content   string `json:"content"`   // 评论内容
	CreatedAt string `json:"createdAt"` // 评论创建时间
	UpdatedAt string `json:"updatedAt"` // 评论更新时间
}

type DeleteCartReq struct {
	UserId    int64 `json:"userId"`    // 用户ID
	ProductId int64 `json:"productId"` // 商品ID
}

type DeleteCommentReq struct {
	CommentId int64 `json:"commentId"` // 评论ID
	UserId    int64 `json:"userId"`    // 用户ID，确保是评论的作者
}

type DeleteProductReq struct {
	ProductId int64 `json:"productId"` // 商品ID
}

type DeleteReplyReq struct {
	ReplyId int64 `json:"replyId"` // 回复ID
	UserId  int64 `json:"userId"`  // 用户ID，确保是回复的作者
}

type EditAddressReq struct {
	Line1     string `json:"line1"`           // 地址行1
	Line2     string `json:"line2,omitempty"` // 地址行2
	City      string `json:"city"`            // 城市
	State     string `json:"state"`           // 省份/州
	Zip       string `json:"zip"`             // 邮政编码
	Country   string `json:"country"`         // 国家
	IsDefault bool   `json:"isDefault"`       // 是否默认地址
}

type EditCommentReq struct {
	CommentId int64  `json:"commentId"` // 评论ID
	UserId    int64  `json:"userId"`    // 用户ID，确保是评论的作者
	Content   string `json:"content"`   // 更新后的评论内容
}

type EditProductReq struct {
	ProductId     int64       `json:"productId"`               // 商品ID
	Name          string      `json:"name,omitempty"`          // 商品名称
	Price         float64     `json:"price,omitempty"`         // 商品价格
	OriginalPrice float64     `json:"originalPrice,omitempty"` // 原价
	Description   string      `json:"description,omitempty"`   // 商品描述
	ImageUrls     []string    `json:"imageUrls,omitempty"`     // 商品图片列表
	Stock         int         `json:"stock,omitempty"`         // 库存数量
	CategoryId    int64       `json:"categoryId,omitempty"`    // 分类ID
	Attributes    []Attribute `json:"attributes,omitempty"`    // 商品属性
}

type EditReplyReq struct {
	ReplyId int64  `json:"replyId"` // 回复ID
	UserId  int64  `json:"userId"`  // 用户ID，确保是回复的作者
	Content string `json:"content"` // 更新后的回复内容
}

type Empty struct {
	Msg string `json:"msg"` // Generic response message
}

type GetCommentsReq struct {
	ProductId int64 `json:"productId"`           // 商品ID
	Page      int   `json:"page,default=1"`      // 页码
	PageSize  int   `json:"pageSize,default=10"` // 每页数量
}

type GetPaymentStatusReq struct {
	PaymentId int64 `json:"paymentId"` // 支付ID
}

type GetProductDetailReq struct {
	ProductId int64 `json:"productId"` // 商品ID
}

type GetProductDetailResp struct {
	Product ProductDetail `json:"product"` // 商品详细信息
}

type GetProductListReq struct {
	CategoryId int64  `json:"categoryId,omitempty"` // 分类ID，可选
	Keyword    string `json:"keyword,omitempty"`    // 搜索关键词，可选
	Page       int    `json:"page,default=1"`       // 页码
	PageSize   int    `json:"pageSize,default=10"`  // 每页数量
	SortBy     string `json:"sortBy,omitempty"`     // 排序字段，例如: price, createdAt
	Order      string `json:"order,omitempty"`      // 排序方式: asc/desc
}

type GetProductListResp struct {
	Products []ProductSummary `json:"products"` // 商品摘要列表
	Total    int              `json:"total"`    // 总数
}

type GetRepliesReq struct {
	CommentId int64 `json:"commentId"`           // 评论ID
	Page      int   `json:"page,default=1"`      // 页码
	PageSize  int   `json:"pageSize,default=10"` // 每页数量
}

type GetSeckillEventsReq struct {
	Page     int `json:"page,default=1"`      // 页码
	PageSize int `json:"pageSize,default=10"` // 每页数量
}

type GetSeckillEventsResp struct {
	Events []SeckillEventResp `json:"events"` // 秒杀活动列表
	Total  int                `json:"total"`  // 秒杀活动总数
}

type GetSeckillProductsReq struct {
	EventId  int64 `json:"eventId"`             // 秒杀活动ID
	Page     int   `json:"page,default=1"`      // 页码
	PageSize int   `json:"pageSize,default=10"` // 每页数量
}

type GetSeckillProductsResp struct {
	Products []SeckillProductResp `json:"products"` // 秒杀商品列表
	Total    int                  `json:"total"`    // 商品总数
}

type LoginHistoryEntry struct {
	Id        int64  `json:"id"`        // 记录ID
	LoginTime string `json:"loginTime"` // 登录时间
	LoginIp   string `json:"loginIp"`   // 登录IP地址
}

type LoginHistoryResp struct {
	History []LoginHistoryEntry `json:"history"` // 登录历史记录列表
}

type LoginReq struct {
	Identifier string `json:"identifier"` // 用户名/邮箱/手机号
	Password   string `json:"password"`   // 密码
}

type LoginResp struct {
	Token string `json:"token"` // JWT Token
}

type MakePaymentReq struct {
	OrderId       int64   `json:"orderId"`       // 订单ID
	UserId        int64   `json:"userId"`        // 用户ID
	Amount        float64 `json:"amount"`        // 支付金额
	PaymentMethod string  `json:"paymentMethod"` // 支付方式 (credit_card, debit_card, paypal, bank_transfer, wallet)
}

type OrderAddress struct {
	Id      int64  `json:"id"`      // 地址ID
	Line1   string `json:"line1"`   // 地址行1
	Line2   string `json:"line2"`   // 地址行2
	City    string `json:"city"`    // 城市
	State   string `json:"state"`   // 省份/州
	Zip     string `json:"zip"`     // 邮政编码
	Country string `json:"country"` // 国家
}

type OrderDetails struct {
	Id          int64           `json:"id"`          // 订单ID
	OrderNumber string          `json:"orderNumber"` // 订单编号
	Status      string          `json:"status"`      // 订单状态
	TotalAmount float64         `json:"totalAmount"` // 订单总金额
	Address     OrderAddress    `json:"address"`     // 收货地址信息
	Items       []OrderItemResp `json:"items"`       // 订单商品项列表
	CreatedAt   string          `json:"createdAt"`   // 订单创建时间
	UpdatedAt   string          `json:"updatedAt"`   // 订单更新时间
}

type OrderDetailsResp struct {
	Order OrderDetails `json:"order"` // 订单详情
}

type OrderItemReq struct {
	ProductId int64 `json:"productId"` // 商品ID
	Quantity  int   `json:"quantity"`  // 购买数量
}

type OrderItemResp struct {
	ProductId int64   `json:"productId"` // 商品ID
	Name      string  `json:"name"`      // 商品名称
	Price     float64 `json:"price"`     // 单价
	Quantity  int     `json:"quantity"`  // 数量
	Total     float64 `json:"total"`     // 商品小计
}

type OrderListResp struct {
	Orders []OrderResp `json:"orders"` // 订单列表
}

type OrderResp struct {
	Id          int64   `json:"id"`          // 订单ID
	OrderNumber string  `json:"orderNumber"` // 订单编号
	Status      string  `json:"status"`      // 订单状态
	TotalAmount float64 `json:"totalAmount"` // 订单总金额
	CreatedAt   string  `json:"createdAt"`   // 订单创建时间
}

type PaginationReq struct {
	Page     int `json:"page,default=1"`      // Page number
	PageSize int `json:"pageSize,default=10"` // Page size
}

type PaymentResp struct {
	PaymentId     int64  `json:"paymentId"`     // 支付ID
	TransactionId string `json:"transactionId"` // 交易流水号
	Status        string `json:"status"`        // 支付状态 (pending, completed, failed)
}

type PaymentStatusResp struct {
	PaymentId int64  `json:"paymentId"` // 支付ID
	Status    string `json:"status"`    // 支付状态 (pending, completed, failed, refunded)
}

type ProductDetail struct {
	Id            int64       `json:"id"`                      // 商品ID
	Name          string      `json:"name"`                    // 商品名称
	Price         float64     `json:"price"`                   // 商品价格
	OriginalPrice float64     `json:"originalPrice,omitempty"` // 原价（如有折扣）
	Description   string      `json:"description"`             // 商品描述
	ImageUrls     []string    `json:"imageUrls"`               // 商品图片列表
	Stock         int         `json:"stock"`                   // 库存数量
	CategoryId    int64       `json:"categoryId"`              // 分类ID
	CategoryName  string      `json:"categoryName"`            // 分类名称
	Attributes    []Attribute `json:"attributes"`              // 商品属性列表
	CreatedAt     string      `json:"createdAt"`               // 商品创建时间
	UpdatedAt     string      `json:"updatedAt"`               // 商品更新时间
}

type ProductSummary struct {
	Id          int64   `json:"id"`          // 商品ID
	Name        string  `json:"name"`        // 商品名称
	Price       float64 `json:"price"`       // 商品价格
	ImageUrl    string  `json:"imageUrl"`    // 商品主图URL
	Description string  `json:"description"` // 简短描述
	Stock       int     `json:"stock"`       // 商品库存
}

type RefundReq struct {
	PaymentId    int64   `json:"paymentId"`    // 支付ID
	RefundAmount float64 `json:"refundAmount"` // 退款金额
	RefundReason string  `json:"refundReason"` // 退款原因
}

type RefundResp struct {
	RefundStatus string `json:"refundStatus"` // 退款状态 (pending, processed, failed)
}

type RegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Avatar   string `json:"avatar"`   // 用户头像URL
}

type ReplyListResp struct {
	Replies []ReplyResp `json:"replies"` // 回复列表
	Total   int         `json:"total"`   // 回复总数
}

type ReplyResp struct {
	Id        int64  `json:"id"`        // 回复ID
	CommentId int64  `json:"commentId"` // 评论ID
	UserId    int64  `json:"userId"`    // 用户ID
	Username  string `json:"username"`  // 回复用户的名称
	Content   string `json:"content"`   // 回复内容
	CreatedAt string `json:"createdAt"` // 回复创建时间
	UpdatedAt string `json:"updatedAt"` // 回复更新时间
}

type SeckillEventResp struct {
	Id        int64  `json:"id"`        // 秒杀活动ID
	Name      string `json:"name"`      // 秒杀活动名称
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
	Status    string `json:"status"`    // 活动状态 (upcoming/ongoing/ended)
}

type SeckillOrderReq struct {
	EventId   int64 `json:"eventId"`   // 秒杀活动ID
	ProductId int64 `json:"productId"` // 商品ID
	UserId    int64 `json:"userId"`    // 用户ID
	Quantity  int   `json:"quantity"`  // 购买数量
}

type SeckillOrderResp struct {
	OrderId int64 `json:"orderId"` // 秒杀订单ID
}

type SeckillProductResp struct {
	Id            int64   `json:"id"`            // 秒杀商品ID
	EventId       int64   `json:"eventId"`       // 秒杀活动ID
	ProductId     int64   `json:"productId"`     // 商品ID
	ProductName   string  `json:"productName"`   // 商品名称
	SeckillPrice  float64 `json:"seckillPrice"`  // 秒杀价格
	OriginalPrice float64 `json:"originalPrice"` // 商品原价
	Stock         int     `json:"stock"`         // 秒杀库存
	ImageUrl      string  `json:"imageUrl"`      // 商品图片URL
}

type SelectCartItemsReq struct {
	UserId     int64   `json:"userId"`     // 用户ID
	ProductIds []int64 `json:"productIds"` // 商品ID列表
	Selected   bool    `json:"selected"`   // 是否选中
}

type SubmitOrderReq struct {
	UserId    int64          `json:"userId"`    // 用户ID
	AddressId int64          `json:"addressId"` // 收货地址ID
	Items     []OrderItemReq `json:"items"`     // 订单商品列表
}

type SubmitOrderResp struct {
	OrderId int64 `json:"orderId"` // 订单ID
}

type UpdateCartReq struct {
	UserId    int64 `json:"userId"`    // 用户ID
	ProductId int64 `json:"productId"` // 商品ID
	Quantity  int   `json:"quantity"`  // 更新后的数量
	Selected  bool  `json:"selected"`  // 是否选中
}

type UpdateUserInfoReq struct {
	UserId   int64  `json:"userId"`             // 用户ID
	Username string `json:"username,omitempty"` // 新用户名
	Email    string `json:"email,omitempty"`    // 新邮箱
	Phone    string `json:"phone,omitempty"`    // 新手机号
	Avatar   string `json:"avatar,omitempty"`   // 新头像
}

type UserInfo struct {
	Id        int64  `json:"id"`        // 用户ID
	Username  string `json:"username"`  // 用户名
	Email     string `json:"email"`     // 邮箱
	Phone     string `json:"phone"`     // 手机号
	Avatar    string `json:"avatar"`    // 用户头像URL
	Role      string `json:"role"`      // 用户角色 (user/admin/super_admin)
	CreatedAt string `json:"createdAt"` // 用户注册时间
}

type UserInfoResp struct {
	User UserInfo `json:"user"` // 用户信息
}