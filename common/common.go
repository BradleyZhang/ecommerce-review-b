package common

import (
	"flag"
)

var (
	Name              = "review-b"
	ReviewServiceName = "review-service" // 需要与 github.com/BradleyZhang/ecommerce-review-service 的 common.Name 一致
)
var (
	PortGRPC = *flag.Int("port", 9091, "the gRPC listening port")
	Port     = 8081
)
var (
	DebugMode = false
)

func initCommonEnv() {
	PortGRPC = GetEnvOrDefault("PORT_GRPC", PortGRPC)
	Port = GetEnvOrDefault("PORT", Port)

	DebugMode = GetEnvOrDefaultBool("DEBUG_MODE", DebugMode)
}
