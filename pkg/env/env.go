package env

import "github.com/dhanusaputra/anywhat-server/util/envutil"

const (
	defaultAuthEnable = true

	defaultExpiredTimeInMinute = 15
)

var (
	// AuthEnable ...
	AuthEnable bool

	// Key ...
	Key []byte
	// JWTExpiredTimeInMinute ...
	JWTExpiredTimeInMinute int
)

// Init ...
func Init() {
	AuthEnable = envutil.GetEnvAsBool("AUTH_ENABLE", defaultAuthEnable)

	Key = []byte(envutil.GetEnv("KEY", ""))
	JWTExpiredTimeInMinute = envutil.GetEnvAsInt("JWT_EXPIRED_TIME_IN_MINUTE", defaultExpiredTimeInMinute)
}
