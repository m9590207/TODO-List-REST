package app

import (
	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/pkg/convert"

	"github.com/gin-gonic/gin"
)

//從query string取出分頁頁數
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//從query string取出每頁最大的筆數,比預設還大或為0重設為系統預設值
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

//計算偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
