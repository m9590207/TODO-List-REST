package middleware

import (
	"fmt"
	"time"

	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/pkg/app"
	"github.com/m9590207/TODO-List-REST/pkg/email"
	"github.com/m9590207/TODO-List-REST/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("例外抛出，發生時間: %d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
