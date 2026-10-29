package data

import (
	"context"

	reviewPb "github.com/BradleyZhang/ecommerce-review-b/api/gen/go/review/v1"
	"github.com/BradleyZhang/ecommerce-review-b/common"
	"github.com/BradleyZhang/ecommerce-review-b/internal/biz"
)

type BusinessRepo struct {
	data *Data
}

// Reply implements [biz.BusinessRepo].
func (b *BusinessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (replyID int64, err error) {
	common.DebugLogf("[data] Reply,param:%v", param)
	res, err := b.data.rc.ReplyReview(ctx, &reviewPb.ReplyReviewRequest{
		ReviewId:  param.ReviewID,
		StoreId:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	common.DebugLogf("ReplayReview return,res:%v err:%v", res, err)
	if err != nil {
		return 0, err
	}
	return res.GetReplyId(), nil
}

func NewBusinessRepo(data *Data) biz.BusinessRepo {
	return &BusinessRepo{
		data: data,
	}
}
