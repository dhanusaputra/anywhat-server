package constant

import "github.com/dhanusaputra/anywhat-server/util/envutil"

const (
	defaultAuthEnable = true
)

var (
	// AuthEnable ...
	AuthEnable bool
)

// Init ...
func Init() {
	AuthEnable = envutil.GetEnvAsBool("AUTH_ENABLE", defaultAuthEnable)
}
