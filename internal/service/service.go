package service

import (
	"context"

	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/internal/dao"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(WithContext(svc.ctx, global.DBEngine))
	return svc
}

//鏈路追蹤
func WithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	if ctx == nil {
		return db
	}
	//SpanFromContext 返回之前與 `ctx` 關聯的 `Span`，如果找不到這樣的 `Span`，則返回 `nil`。
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return db
	}
	return db.Set("opentracing:parent.span", parentSpan)
}
