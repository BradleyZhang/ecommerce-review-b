package server

import (
	"context"
	"net/http"
	"strconv"

	v1 "github.com/BradleyZhang/ecommerce-review-b/api/gen/go/business/v1"
	"github.com/BradleyZhang/ecommerce-review-b/common"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails" // 注册 google.rpc.ErrorInfo 到 protobuf 的全局类型注册表。因为 review-b 中没有使用 errdetails 的，而 review-service 的 gerr 包返回错误会使用errdetails.ErrorInfo，不导入初始化的化，gateway 会Failed to marshal

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHTTPHandler(ctx context.Context) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	err := mux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := "localhost:" + strconv.Itoa(common.PortGRPC)
	err = v1.RegisterBusinessServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return nil, err
	}
	return mux, nil
}
