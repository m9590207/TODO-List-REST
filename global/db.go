package global

import "gorm.io/gorm"

//服務的全域變數
var (
	DBEngine *gorm.DB
)
