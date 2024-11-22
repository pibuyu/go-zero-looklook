package homestayComment

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"log"
	"looklook/app/travel/model"
	"looklook/common/xerr"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentListLogic {
	return CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req types.CommentListReq) (*types.CommentListResp, error) {

	whereBuilder := l.svcCtx.HomestayCommentModel.SelectBuilder().Where(squirrel.Eq{
		"del_state": model.HomestayCommentExistStatus,
		"user_id":   req.LastId,
	})
	res, err := l.svcCtx.HomestayCommentModel.FindPageListByPage(l.ctx, whereBuilder, req.PageNo, req.PageSize, "create_time desc")
	log.Printf("查询comment得到的结果为:%v", res)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get homestay comment by id fail:err : %v", err)
	}
	var resp []types.HomestayComment
	//用mapreduce批量转化
	if len(res) > 0 {
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, val := range res {
				source <- *val
			}
		}, func(item interface{}, writer mr.Writer[model.HomestayComment], cancel func(error)) {
			writer.Write(item.(model.HomestayComment))
		}, func(pipe <-chan model.HomestayComment, cancel func(error)) {
			for comment := range pipe {
				var typesComment types.HomestayComment
				_ = copier.Copy(&typesComment, &comment)
				resp = append(resp, typesComment)
			}
		})
	}
	return &types.CommentListResp{
		List: resp,
	}, nil
}
