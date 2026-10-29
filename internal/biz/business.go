package biz

import (
	"context"

	"github.com/BradleyZhang/ecommerce-review-b/common"
)

type BusinessRepo interface {
	Reply(ctx context.Context, param *ReplyParam) (replyID int64, err error)
}
type BusinessUsecase struct {
	repo BusinessRepo
}

func NewBusinessUsecase(repo BusinessRepo) *BusinessUsecase {
	return &BusinessUsecase{repo: repo}
}

type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

func (b *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	common.DebugLogf("[biz] CreateReply param:%v", param)
	return b.repo.Reply(ctx, param)
}
