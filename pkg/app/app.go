package app

import (
	"net/http"

	"github.com/m9590207/TODO-List-REST/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//回應處理
type Response struct {
	Ctx *gin.Context
}

//分頁
type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//無分頁
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//有分頁
func (r *Response) ToResponseList(list interface{}, totalRows int64) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//回傳自訂的錯誤代碼,訊息,呼叫堆疊發生的錯誤, 自訂代碼轉HTTP Status Code
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
