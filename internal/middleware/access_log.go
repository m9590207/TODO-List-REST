package middleware

import (
	"bytes"
	"time"

	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/pkg/logger"

	"github.com/gin-gonic/gin"
)

//實現http.ResponseWriter interface struct 和 Write方法就可以取得response body
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//實現http.ResponseWriter Write方法
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

//自訂的log訊息
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
