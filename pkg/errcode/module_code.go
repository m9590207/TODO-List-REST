package errcode

//自訂的錯誤代碼,訊息 全域變數
var (
	ErrorGetTodoListFail = NewError(2101, "取得列表失敗")
	ErrorCreateTodoFail  = NewError(2102, "建立失敗")
	ErrorUpdateTodoFail  = NewError(2103, "更新失敗")
	ErrorDeleteTodoFail  = NewError(2104, "刪除失敗")
	ErrorCountTodoFail   = NewError(2105, "統計失敗")
)
