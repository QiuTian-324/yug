package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"
	"yug_server/global"
	"yug_server/internal/libs"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var blacklistedIPs = make(map[string]bool)

type RateLimitInfo struct {
	Requests      int
	LastTime      time.Time
	ExceededCount int
}

const (
	RateLimit          = 2
	RateLimitTime      = time.Minute
	BlacklistThreshold = 5
)

// 分布式速率限制
func isRateLimited(ip string) bool {
	key := fmt.Sprintf("rate_limit:%s", ip)
	val, err := global.RedisClient.Get(context.Background(), key).Int()
	if err != nil && err != redis.Nil {
		pkg.Error("Redis错误", err)
		return false
	}

	if val >= RateLimit {
		return true
	}

	global.RedisClient.Incr(context.Background(), key)
	global.RedisClient.Expire(context.Background(), key, RateLimitTime)
	return false
}

// AuthToken token校验
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// 检查IP是否在黑名单中
		if blacklistedIPs[ip] {
			pkg.Error("IP在黑名单中", nil)
			libs.ForbiddenResponse(c, fmt.Sprintf("访问被拒绝，IP %s 在黑名单中！", ip))
			return
		}

		// 分布式速率限制检查
		if isRateLimited(ip) {
			pkg.Error("请求过于频繁", nil)
			libs.TooManyRequestsResponse(c, fmt.Sprintf("请求过于频繁，请稍后再试，IP %s 已记录！", ip))
			return
		}

		// 优先从��求头获取 token
		token := c.Request.Header.Get("Authorization")

		// 如果请求头中没有 token，尝试从查询参数中获取
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			ip := c.ClientIP()
			pkg.Error("token 不存在", nil)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，请停止任何非法操作，IP %s 已记录！", ip))
			return
		}

		// 分割 token
		parts := strings.SplitN(token, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			token = parts[1]
		}
		// 尝试解密 token
		currentUser, err := libs.ParseToken(token)
		if err != nil {
			pkg.Error("token 解密失败", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，解密 token 失败，IP %s 已记录！", ip))
			return
		}

		// 检查 Redis 中是否存在 token
		redisKey := fmt.Sprintf("%s:%d", global.UserRedisSessionKey, currentUser.ID)
		_, err = libs.RedisGet(c, redisKey)
		if err != nil {
			pkg.Error("Redis 中不存在 token", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，用户未认证，IP %s 已记录！", ip))
			return
		}

		// 将用户名设置到上下文中
		c.Set("id", currentUser.ID)
		c.Next()
	}
}
