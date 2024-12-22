package logic

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/model"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 参数验证
	if err := l.validateRegisterRequest(in); err != nil {
		return nil, err
	}

	// 检查用户名、邮箱、手机号是否已存在
	if err := l.checkUserExists(in); err != nil {
		return nil, err
	}

	// 如果提供了管理员注册码，验证它
	var role string = "user"
	if in.AdminCode != "" {
		if err := l.validateAdminCode(in.AdminCode); err != nil {
			return nil, err
		}
		role = "admin"
	}

	hashedPassword, err := l.hashPassword(in.Password)
	if err != nil {
		logx.Errorf("密码加密失败: %v", err)
		return nil, status.Error(codes.Internal, "内部服务错误")
	}

	// 开启事务进行用户注册
	// 事务会对涉及的行加锁，防止其他事务同时修改这些数据
	var userId int64
	err = l.svcCtx.UsersModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 5.1 创建用户
		result, err := l.svcCtx.UsersModel.Insert(ctx, &model.Users{
			Username: in.Username,
			Password: hashedPassword,
			Email:    in.Email,
			Phone:    in.Phone,
			Avatar:   sql.NullString{String: in.Avatar, Valid: in.Avatar != ""},
			Role:     role,
		})
		if err != nil {
			return err
		}

		userId, err = result.LastInsertId()
		if err != nil {
			return err
		}

		// 5.2 如果是管理员注册，标记注册码为已使用
		if in.AdminCode != "" {
			if err := l.svcCtx.AdminCodesModel.MarkCodeAsUsed(ctx, in.AdminCode); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logx.Errorf("用户注册失败: %v", err)
		return nil, status.Error(codes.Internal, "注册失败")
	}

	return &user.RegisterResponse{
		UserId: userId,
	}, nil
}

// 验证注册参数
func (l *RegisterLogic) validateRegisterRequest(in *user.RegisterRequest) error {
	if in.Username == "" || len(in.Username) < 3 {
		return status.Error(codes.InvalidArgument, "用户名不能为空且长度至少为3个字符")
	}
	if in.Password == "" || len(in.Password) < 6 {
		return status.Error(codes.InvalidArgument, "密码不能为空且长度至少为6个字符")
	}
	if in.Email == "" || !isValidEmail(in.Email) {
		return status.Error(codes.InvalidArgument, "邮箱格式不正确")
	}
	if in.Phone == "" || !isValidPhone(in.Phone) {
		return status.Error(codes.InvalidArgument, "手机号格式不正确")
	}
	return nil
}

// 检查用户是否已存在
func (l *RegisterLogic) checkUserExists(in *user.RegisterRequest) error {
	// 检查用户名
	if _, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.Username); err == nil {
		return status.Error(codes.AlreadyExists, "用户名已存在")
	}

	// 检查邮箱
	if _, err := l.svcCtx.UsersModel.FindOneByEmail(l.ctx, in.Email); err == nil {
		return status.Error(codes.AlreadyExists, "邮箱已被注册")
	}

	// 检查手机号
	if _, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone); err == nil {
		return status.Error(codes.AlreadyExists, "手机号已被注册")
	}

	return nil
}

// 验证管理员注册码
func (l *RegisterLogic) validateAdminCode(code string) error {
	adminCode, err := l.svcCtx.AdminCodesModel.FindValidCode(l.ctx, code)
	if err != nil {
		if err == model.ErrNotFound {
			return status.Error(codes.InvalidArgument, "无效的管理员注册码")
		}
		return status.Error(codes.Internal, "验证管理员注册码失败")
	}
	if adminCode.Used == 1 {
		return status.Error(codes.InvalidArgument, "该注册码已被使用")
	}
	return nil
}

// 密码加密
func (l *RegisterLogic) hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password+l.svcCtx.Config.Salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// 邮箱格式验证
func isValidEmail(email string) bool {
	// 使用正则表达式验证邮箱格式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 手机号格式验证
func isValidPhone(phone string) bool {
	// 这里使用简单的中国大陆手机号验证规则
	// 支持 13x, 14x, 15x, 16x, 17x, 18x, 19x 开头的手机号
	pattern := `^1[3-9]\d{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}
