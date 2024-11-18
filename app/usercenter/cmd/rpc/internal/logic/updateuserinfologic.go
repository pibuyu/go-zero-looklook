package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {

	//先根据id或者mobile查询是否存在这个user
	user := &model.User{}
	err := l.svcCtx.DB.Table("user").Where("id = ? or mobile = ?", in.UserId, in.Mobile).
		Find(&user).Error

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "query userinfo from database failed,id=%d,err:%v", in.UserId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "User with mobile:%s or id:%d not exists", in.Mobile, in.UserId)
	}
	if in.Nickname != "" {
		user.Nickname = in.Nickname
	}
	if in.Password != "" {
		user.Password = in.Password
	}
	if in.Avatar != "" {
		user.Avatar = in.Avatar
	}
	user.Sex = int64(in.Sex)
	if in.Info != "" {
		user.Info = in.Info
	}

	//再将user更新回去
	_, err = l.svcCtx.UserModel.Update(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update userinfo failed,id=%d,err:%v", in.UserId, err)
	}

	return &pb.UpdateUserInfoResp{
		EffectedRows: 1,
	}, nil
}
