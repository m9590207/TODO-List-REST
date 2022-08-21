package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/internal/middleware"
	"github.com/m9590207/TODO-List-REST/internal/routers/api"
	v1 "github.com/m9590207/TODO-List-REST/internal/routers/api/v1"
	"github.com/m9590207/TODO-List-REST/pkg/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	//debug使用gin提供的logger和recovery輸出log跟例外錯誤處理
	//非debug就使用自訂的錯誤訊息格式,錯誤訊息輸出通知
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Tracing())

	r.GET("/auth/:appKey/:appSecret", api.GetAuth)
	todoList := v1.NewTodo()
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/:createdBy/:state/todos", todoList.List)
		apiv1.POST("/todo", todoList.Create)
		apiv1.DELETE("/todo/:id", todoList.Delete)
		apiv1.PUT("/todo/:id", todoList.Update)
		apiv1.PATCH("/todo/:id", todoList.Update)
	}
	return r
}
