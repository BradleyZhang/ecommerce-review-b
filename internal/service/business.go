package service

import (
	"context"

	pb "github.com/BradleyZhang/ecommerce-review-b/api/gen/go/business/v1"
	"github.com/BradleyZhang/ecommerce-review-b/common"
	"github.com/BradleyZhang/ecommerce-review-b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServiceServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}
func (bs *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewResponse, error) {
	common.DebugLogf("[service] ReplyReview req:%v", req)
	replyID, err := bs.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewId(),
		StoreID:   req.GetStoreId(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewResponse{ReplyId: replyID}, nil
}
