package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update user info
func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (*types.UpdateUserInfoResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	result, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &usercenter.UpdateUserInfoReq{
		UserId:   userId,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Info:     req.Info,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var resp types.UpdateUserInfoResp
	_ = copier.Copy(&resp, result)
	return &resp, nil
}
