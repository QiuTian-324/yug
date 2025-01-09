package handlers

import (
	"fmt"
	"time"
	"yug_server/global"
	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
	"yug_server/internal/libs"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type UserHandler struct {
	uc     *services.UserUseCase
	rds    *redis.Client
	logger *zap.Logger
}

func NewUserHandler(uc *services.UserUseCase, rds *redis.Client, logger *zap.Logger) *UserHandler {
	return &UserHandler{uc: uc, rds: rds, logger: logger}
}

// 注册
func (h *UserHandler) Register(ctx *gin.Context) {

	req := new(dto.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.ParamError(ctx, "请求数据无效")
		return
	}

	if err := h.uc.Register(ctx, req); err != nil {
		h.logger.Error("注册失败", zap.Error(err))
		libs.Internal(ctx, "注册失败")
		return
	}

	libs.OK(ctx, "注册成功", nil)
}

// 登录
func (h *UserHandler) Login(ctx *gin.Context) {

	req := new(dto.LoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.ParamError(ctx, "请求数据无效")
		return
	}

	if libs.ValidateEmpty(req.Username) {
		h.logger.Error("用户名不能为空", zap.Error(err))
		libs.ParamError(ctx, "用户名不能为空")
		return
	}

	if libs.ValidateEmpty(req.Password) {
		h.logger.Error("密码不能为空", zap.Error(err))
		libs.ParamError(ctx, "密码不能为空")
		return
	}

	userInfo, err := h.uc.Login(ctx, req.Username, req.Password)
	if err != nil {
		h.logger.Error("用户不存在", zap.Error(err))
		libs.Internal(ctx, "用户不存在")
		return
	}

	token, err := libs.GenToken(uint64(userInfo.ID), userInfo.Username)
	if err != nil {
		h.logger.Error("生成token失败", zap.Error(err))
		libs.Internal(ctx, "登陆失败")
		return
	}

	redisKey := fmt.Sprintf("%s:%d", global.UserRedisSessionKey, userInfo.ID)

	err = h.rds.Set(ctx, redisKey, token, time.Hour*time.Duration(viper.GetInt64("redis.expires"))).Err()
	if err != nil {
		h.logger.Error("设置redis失败", zap.Error(err))
		libs.Internal(ctx, "登录失败")
		return
	}

	res := dto.LoginResponse{
		UserID: cast.ToString(userInfo.ID),
		Token:  token,
	}

	libs.OK(ctx, "登录成功", res)
}

// 查询用户
func (h *UserHandler) QueryUser(ctx *gin.Context) {
	username := ctx.Query("username")
	email := ctx.Query("email")
	phone := ctx.Query("phone")

	if libs.ValidateEmpty(username) && libs.ValidateEmpty(email) && libs.ValidateEmpty(phone) {
		h.logger.Error("请求参数不能同时为空")
		libs.ParamError(ctx, "请输入用户名、邮箱或手机号")
		return
	}

	friendInfo, err := h.uc.QueryUser(ctx, username, email, phone)
	if err != nil {
		h.logger.Error("查询用户失败", zap.Error(err))
		libs.Internal(ctx, "查询用户失败")
		return
	}

	res := dto.QueryUserResponse{
		UserID:    friendInfo.ID,
		Username:  friendInfo.Username,
		Nickname:  friendInfo.Nickname,
		Email:     friendInfo.Email,
		Phone:     friendInfo.Phone,
		AvatarUrl: friendInfo.AvatarUrl,
	}

	libs.OK(ctx, "查询成功", res)
}

// 退出登录
func (h *UserHandler) Logout(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	redisKey := fmt.Sprintf("%s%d", global.UserRedisSessionKey, userID)
	h.rds.Del(ctx, redisKey)
	libs.OK(ctx, "退出成功", nil)
}

// 添加好友
func (h *UserHandler) AddFriend(ctx *gin.Context) {
	req := new(dto.AddFriendRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.ParamError(ctx, "请求数据无效")
		return
	}

	userID := ctx.MustGet("id").(uint64)

	if libs.ValidateEmpty(req.FriendID) {
		h.logger.Error("好友ID不能为空")
		libs.ParamError(ctx, "好友ID不能为空")
		return
	}

	if userID == cast.ToUint64(req.FriendID) {
		h.logger.Error("不能添加自己为好友")
		libs.ParamError(ctx, "不能添加自己为好友")
		return
	}

	err := h.uc.AddFriend(ctx, userID, cast.ToUint64(req.FriendID))
	if err != nil {
		h.logger.Error("添加好友失败", zap.Error(err))
		libs.Internal(ctx, err.Error())
		return
	}

	libs.OK(ctx, "添加成功", nil)
}

// 获取好友列表
func (h *UserHandler) GetFriends(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	friends, err := h.uc.GetFriends(ctx, userID)
	if err != nil {
		h.logger.Error("获取好友列表失败", zap.Error(err))
		libs.Internal(ctx, err.Error())
		return
	}

	libs.OK(ctx, "获取成功", friends)
}

// 获取用户信息
func (h *UserHandler) GetUserInfo(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	var user model.User
	global.DB.Model(&user).Where("id = ?", userID).First(&user)
	libs.OK(ctx, "获取成功", user)
}
