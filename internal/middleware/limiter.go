package middleware

import (
	"github.com/m9590207/TODO-List-REST/pkg/app"
	"github.com/m9590207/TODO-List-REST/pkg/errcode"
	"github.com/m9590207/TODO-List-REST/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			//佔用一個權杖,如沒有可用權限count=0
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
