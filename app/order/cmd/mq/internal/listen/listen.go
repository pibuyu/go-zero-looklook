package listen

import (
	"context"
	"looklook/app/order/cmd/mq/internal/config"
	"looklook/app/order/cmd/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// back to all consumers
// 管理不同类别的mq,然后在main中调用listen.Mqs可以获取所有mq一起start
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：pub sub
	services = append(services, KqMqs(c, ctx, svcContext)...)

	return services
}
