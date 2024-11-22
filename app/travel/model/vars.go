package model

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// 民宿活动类型

var HomestayActivityPreferredType = "preferredHomestay" //优选民宿
var HomestayActivityGoodBusiType = "goodBusiness"       //最佳房东

// 民宿活动上下架

var HomestayActivityDownStatus int64 = 0 //下架
var HomestayActivityUpStatus int64 = 1   //上架

// 民俗评论类型.todo:This is a special comment.注释和变量不能挨在一起，不然识别不到，不知道是为啥

var HomestayCommentExistStatus int64 = 0   //未删除
var HomestayCommentDeletedStatus int64 = 1 //已删除
