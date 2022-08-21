package errcode

//自訂的錯誤代碼,訊息 全域變數
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "服務內部錯誤")
	InvalidParams             = NewError(1001, "導入參數錯誤")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "驗證失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1004, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(1005, "驗證失敗，Token逾時")
	UnauthorizedTokenGenerate = NewError(1006, "驗證失敗，Token產生失敗")
	TooManyRequests           = NewError(1007, "請求過多")
)
