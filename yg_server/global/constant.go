package global

import "fmt"

// redis key
const (
	UniqueKeySeed string = "Key12368975490255" // UniqueKeySeed 随机数相关的种子
)

var (
	UserRedisPrefix       string = "user"                                              // UserRedisPrefix 用户相关的前缀
	UserRedisSessionKey   string = fmt.Sprintf("%s:%s", UserRedisPrefix, "session")    // UserRedisSessionKey 会话相关的前缀
	UserRedisDataCacheKey string = fmt.Sprintf("%s:%s", UserRedisPrefix, "data_cache") // UserRedisDataCacheKey 缓存相关的前缀
	UserRedisConfigKey    string = fmt.Sprintf("%s:%s", UserRedisPrefix, "config")     // UserRedisConfigKey 配置相关的前缀
	UserRedisOnline       string = fmt.Sprintf("%s:%s", UserRedisPrefix, "online")     // UserRedisOnline 用户在线状态
)
var (
	ChatRedisPrefix       string = "chat"                                              // ChatRedisPrefix redis存储随机前缀
	ChatRedisUserKey      string = fmt.Sprintf("%s:%s", ChatRedisPrefix, "user")       // ChatRedisUserKey 用户相关的前缀
	ChatRedisSessionKey   string = fmt.Sprintf("%s:%s", ChatRedisUserKey, "session")   // ChatRedisSessionKey 会话相关的前缀
	ChatRedisDataCacheKey string = fmt.Sprintf("%s:%s", ChatRedisPrefix, "data_cache") // ChatRedisDataCacheKey 缓存相关的前缀
	ChatRedisOnline       string = fmt.Sprintf("%s:%s", ChatRedisPrefix, "online")     // ChatRedisOnline 用户在线状态
)
