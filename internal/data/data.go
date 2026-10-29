package data

import (
	"context"
	"fmt"

	reviewPb "github.com/BradleyZhang/ecommerce-review-b/api/gen/go/review/v1"
	"github.com/BradleyZhang/ecommerce-review-b/common"
	_ "github.com/mbobakov/grpc-consul-resolver" // gRPC resolver for services registered in consul
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
	opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin": {}}]}`))

	connString := fmt.Sprintf("consul://%s/%s?[healthy=true]&[wait=]&[near=]&[insecure=]&[limit=]&[tag=]&[token=]", common.ConsulAddress, common.ReviewServiceName)
	conn, err := grpc.NewClient(
		connString,
		opts...)
	if err != nil {
		common.FatalLog("failed to connect review service gRPC server,err:", err)
	}
	defer func() {
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				common.ErrorLogf("gRPC conn close err:%v", cerr)
			}
			common.SysLog("gRPC conn close")
		}()
	}()
	client := reviewPb.NewReviewServiceClient(conn)
	return client
}
