package data

import (
	"context"

	reviewPb "github.com/BradleyZhang/ecommerce-review-b/api/gen/go/review/v1"
	"github.com/BradleyZhang/ecommerce-review-b/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	rc reviewPb.ReviewServiceClient
}

func NewData(rc reviewPb.ReviewServiceClient) (d *Data, cleanup func(), err error) {
	cleanup = func() {
		common.SysLog("closing the data resources")
	}
	return &Data{
		rc: rc,
	}, cleanup, nil
}
func NewReviewServiceClient(ctx context.Context) reviewPb.ReviewServiceClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	if err != nil {
		common.FatalLog("can't connect to review service gRPC server,err:", err)
	}
	defer func() {
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				common.ErrorLogf("gRPC conn close err:%v", cerr)
			}
		}()
	}()
	client := reviewPb.NewReviewServiceClient(conn)
	return client
}
