package middleware

import "github.com/gin-gonic/gin"

//使用gin.Context提供的setter和getter在context加入自訂的資訊
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "todo_list_service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
