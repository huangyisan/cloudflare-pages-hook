package logic

import (
	"cloudflare-pages-hook/notify/internal/svc"
	"cloudflare-pages-hook/notify/internal/types"
	"cloudflare-pages-hook/notify/pkg/notifier"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyLogic {
	return &NotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyLogic) Notify(req *types.NotifyReq) (resp *types.NotifyResp, err error) {
	// todo: add your logic here and delete this line
	log.Printf("%+v", req)
	notifier.N.SetChatID(req.Chat_id)
	err = notifier.N.Send(req.Text)
	if err != nil {
		return &types.NotifyResp{
			Status: err.Error(),
		}, err
	}
	return &types.NotifyResp{
		Status: "ok",
	}, nil
}
