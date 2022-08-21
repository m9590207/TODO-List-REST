package global

import (
	"github.com/m9590207/TODO-List-REST/pkg/logger"
	"github.com/m9590207/TODO-List-REST/pkg/setting"
)

//服務的全域變數
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
