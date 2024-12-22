package zeroerr

const (
	ErrUserExists         ErrorCode = iota + 2000 // 用户已存在
	ErrUserNotFound                               // 用户未找到
	ErrInvalidPassword                            // 密码无效
	ErrInvalidToken                               // 无效的令牌
	ErrExpiredToken                               // 令牌已过期
	ErrAdminCodeUsed                              // 管理员邀请码已使用
	ErrAdminCodeExpired                           // 管理员邀请码已过期
	ErrAdminCodeNotFound                          // 管理员邀请码未找到
	ErrAddressNotFound                            // 地址未找到
	ErrCollectionNotFound                         // 收藏未找到
	// ... 添加更多的错误码
)

var userMsg map[ErrorCode]string

func init() {
	userMsg = make(map[ErrorCode]string)

	userMsg[ErrUserExists] = "用户已存在"
	userMsg[ErrUserNotFound] = "用户未找到"
	userMsg[ErrInvalidPassword] = "密码无效"
	userMsg[ErrInvalidToken] = "无效的令牌"
	userMsg[ErrExpiredToken] = "令牌已过期"
	userMsg[ErrAdminCodeUsed] = "管理员邀请码已使用"
	userMsg[ErrAdminCodeExpired] = "管理员邀请码已过期"
	userMsg[ErrAdminCodeNotFound] = "管理员邀请码未找到"
	userMsg[ErrAddressNotFound] = "地址未找到"
	userMsg[ErrCollectionNotFound] = "收藏未找到"
}
