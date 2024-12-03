package handlers

import (
	"fmt"
	"time"
	"yug_server/global"
	"yug_server/internal/dto"
	"yug_server/internal/libs"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type UserHandler struct {
	uc     *services.UserUseCase
	logger *zap.Logger
}

func NewUserHandler(uc *services.UserUseCase, logger *zap.Logger) *UserHandler {
	return &UserHandler{uc: uc, logger: logger}
}

func (h *UserHandler) Register(ctx *gin.Context) {

	req := new(dto.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.BadRequestResponse(ctx, "请求数据无效")
		return
	}

	if err := h.uc.Register(ctx, req); err != nil {
		h.logger.Error("注册失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "注册失败")
		return
	}

	libs.SuccessResponse(ctx, "注册成功", nil)
}

func (h *UserHandler) Login(ctx *gin.Context) {

	req := new(dto.LoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.BadRequestResponse(ctx, "请求数据无效")
		return
	}

	if libs.ValidateEmpty(req.Username) {
		h.logger.Error("用户名不能为空", zap.Error(err))
		libs.BadRequestResponse(ctx, "用户名不能为空")
		return
	}

	if libs.ValidateEmpty(req.Password) {
		h.logger.Error("密码不能为空", zap.Error(err))
		libs.BadRequestResponse(ctx, "密码不能为空")
		return
	}

	userInfo, err := h.uc.Login(ctx, req.Username, req.Password)
	if err != nil {
		h.logger.Error("用户不存在", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "用户不存在")
		return
	}

	token, err := libs.GenToken(uint64(userInfo.ID), userInfo.Username)
	if err != nil {
		h.logger.Error("生成token失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "登陆失败")
		return
	}

	redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userInfo.ID)

	err = libs.RedisSet(ctx, redisKey, token, time.Hour*time.Duration(viper.GetInt64("redis.expires")))
	if err != nil {
		h.logger.Error("设置redis失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "登录失败")
		return
	}

	res := dto.LoginResponse{
		UserID:    cast.ToString(userInfo.ID),
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		AvatarUrl: userInfo.AvatarUrl,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Bio:       userInfo.Bio,
		Online:    userInfo.Online,
	}

	// 使用 AddExtra 添加 token 作为额外字段
	response := libs.NewResponse(libs.CodeSuccess, "登录成功", true, res, nil)
	libs.AddExtra(ctx, response, map[string]interface{}{"token": token})
}

// func (h *UserHandler) AddFriend(ctx *gin.Context) {
// 	userID := ctx.MustGet("id").(uint64)

// 	// 获取请求参数
// 	req := new(dto.AddFriendRequest)
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		h.logger.Error("请求数据无效", zap.Error(err))
// 		libs.BadRequestResponse(ctx, "请求数据无效")
// 		return
// 	}

// 	if err := h.uc.AddFriend(ctx, userID, cast.ToUint64(req.FriendID)); err != nil {
// 		h.logger.Error("添加好友失败", zap.Error(err))
// 		libs.InternalServerErrorResponse(ctx, err.Error())
// 		return
// 	}

// 	libs.SuccessResponse(ctx, "添加好友成功", nil)
// }

func (h *UserHandler) GetUserSeesionList(ctx *gin.Context) {

}

// func (h *UserHandler) GetFriends(ctx *gin.Context) {
// 	userID := ctx.MustGet("id").(uint64)

// 	res, err := h.uc.GetFriends(ctx, userID)
// 	if err != nil {
// 		h.logger.Error("获取好友列表失败", zap.Error(err))
// 		libs.InternalServerErrorResponse(ctx, "获取好友列表失败")
// 	}
// 	libs.SuccessResponse(ctx, "获取好友列表成功", res)
// }

func (h *UserHandler) Logout(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userID)
	libs.RedisDelete(ctx, redisKey)
	libs.SuccessResponse(ctx, "退出成功", nil)
}

// func GithubAuthURL(ctx *gin.Context) {
// 	oauth2Config := &oauth2.Config{
// 		ClientID:     "Iv23liTPWZse8DM9fHSR",
// 		ClientSecret: "221698c4c550fa6c1a7f80371fd680f1c4ec8576",
// 		RedirectURL:  "http://localhost:5173/login",
// 		Scopes:       []string{"user:email"},
// 		Endpoint:     github.Endpoint,
// 	}

// 	req := new(dto.GithubAuthURLRequest)
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		libs.BadRequestResponse(ctx, "请求数据无效")
// 		pkg.Error("请求数据无效", err)
// 		return
// 	}

// 	url := oauth2Config.AuthCodeURL(req.State, oauth2.AccessTypeOffline)

// 	data := map[string]interface{}{
// 		"url": url,
// 	}

// 	libs.SuccessResponse(ctx, "github登录成功", data)

// }

// func GithubCallback(ctx *gin.Context) {

// 	db := ctx.MustGet("db").(*gorm.DB)
// 	res := new(dto.GithubLoginResponse)

// 	oauth2Config := &oauth2.Config{
// 		ClientID:     "Iv23liTPWZse8DM9fHSR",
// 		ClientSecret: "221698c4c550fa6c1a7f80371fd680f1c4ec8576",
// 		RedirectURL:  "http://localhost:5173/login",
// 		Scopes:       []string{"user:email"},
// 		Endpoint:     github.Endpoint,
// 	}
// 	// 获取code
// 	code := ctx.Query("code")
// 	if code == "" {
// 		libs.BadRequestResponse(ctx, "code不能为空")
// 		return
// 	}

// 	// 获取access_token
// 	authToken, err := oauth2Config.Exchange(ctx, code)
// 	if err != nil {
// 		pkg.Error("获取access_token失败", err)
// 		libs.InternalServerErrorResponse(ctx, "获取access_token失败")
// 		return
// 	}

// 	client := oauth2Config.Client(context.Background(), authToken)
// 	resp, err := client.Get("https://api.github.com/user")
// 	if err != nil {
// 		http.Error(ctx.Writer, fmt.Sprintf("Failed to fetch user info: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	var user map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
// 		http.Error(ctx.Writer, fmt.Sprintf("Failed to decode user info: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = services.RegisterUser(db, &dto.RegisterRequest{
// 		Username: user["login"].(string),
// 		Avatar:   user["avatar_url"].(string),
// 		Nickname: user["name"].(string),
// 	})

// 	if err != nil {
// 		// 检查是否是唯一性约束错误
// 		if strings.Contains(err.Error(), "Duplicate entry") {
// 			// 用户已注册，直接登录
// 			userInfo, err := services.LoginByThirdParty(db, user["login"].(string))
// 			if err != nil {
// 				libs.InternalServerErrorResponse(ctx, "用户不存在")
// 				pkg.Error("用户不存在", err)
// 				return
// 			}

// 			token, err := libs.GenToken(userInfo.ID, userInfo.Username)
// 			if err != nil {
// 				libs.InternalServerErrorResponse(ctx, "登陆失败")
// 				pkg.Error("生成token失败", err)
// 				return
// 			}

// 			redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userInfo.ID)

// 			err = libs.RedisSet(ctx, redisKey, token, time.Hour*time.Duration(viper.GetInt64("redis.expires")))
// 			if err != nil {
// 				libs.InternalServerErrorResponse(ctx, "登录失败")
// 				pkg.Error("设置redis失败", err)
// 				return
// 			}

// 			res.Token = token
// 			libs.SuccessResponse(ctx, "github登录成功", res)
// 			return
// 		}

// 		pkg.Error("注册用户失败", err)
// 		libs.InternalServerErrorResponse(ctx, "登陆失败")
// 		return
// 	}

// 	userInfo, err := services.LoginByThirdParty(db, user["login"].(string))
// 	if err != nil {
// 		libs.InternalServerErrorResponse(ctx, "用户不存在")
// 		pkg.Error("用户不存在", err)
// 		return
// 	}

// 	token, err := libs.GenToken(userInfo.ID, userInfo.Username)
// 	if err != nil {
// 		libs.InternalServerErrorResponse(ctx, "登陆失败")
// 		pkg.Error("生成token失败", err)
// 		return
// 	}

// 	redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userInfo.ID)

// 	err = libs.RedisSet(ctx, redisKey, token, time.Hour*time.Duration(viper.GetInt64("redis.expires")))
// 	if err != nil {
// 		libs.InternalServerErrorResponse(ctx, "登录失败")
// 		pkg.Error("设置redis失败", err)
// 		return
// 	}

// 	res.Token = token

// 	// 成功获取用户信息后重定向
// 	// ctx.Redirect(http.StatusSeeOther, returnedState)
// 	libs.SuccessResponse(ctx, "github登录成功", res)

// }
