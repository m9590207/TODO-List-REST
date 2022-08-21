package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

//設定服務統一的逾時時間
//設定10秒由A發出請求所有服務需在10秒內返回到A
//如服務B花了7秒,在服務D只剩下3秒如超過時間,服務B即會主動斷開
//若某個服務要改變逾時時間只須呼叫context.WithTimeout重新設定再進行新的傳遞
//              服務A - 服務B - 服務D
//					\
//                    服務C
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//使用context.WithTimeout設定設定逾時時間,再重新指定給gin.Context
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
