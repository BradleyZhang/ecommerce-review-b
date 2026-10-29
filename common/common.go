package common

import (
	"flag"
	"net"
)

var (
	Name              = "review-b"
	ReviewServiceName = "review-service" // 需要与 github.com/BradleyZhang/ecommerce-review-service 的 common.Name 一致
)

// ==== 下面的从 .env 更新 ====
var (
	PortGRPC = *flag.Int("grpcPort", 9091, "the gRPC listening port")
	Port     = 8081
	IP       = "127.0.0.1"
)
var (
	DebugMode = false
)

func initCommonEnv() {
	{
		PortGRPC = GetEnvOrDefault("PORT_GRPC", PortGRPC)
		Port = GetEnvOrDefault("PORT", Port)

		{ // 尝试获得 IP
			ip, err := getOutboundIP()
			if err == nil {
				IP = ip.String()
			}
		}
		IP = GetEnvOrDefaultString("IP", IP)
	}
	{
		DebugMode = GetEnvOrDefaultBool("DEBUG_MODE", DebugMode)
	}
}
func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
