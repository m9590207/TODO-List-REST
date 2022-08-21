package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/m9590207/TODO-List-REST/global"
	"github.com/m9590207/TODO-List-REST/internal/service"
	"github.com/m9590207/TODO-List-REST/pkg/app"
	"github.com/m9590207/TODO-List-REST/pkg/convert"
	"github.com/m9590207/TODO-List-REST/pkg/errcode"
)

type TodoList struct{}

func NewTodo() TodoList {
	return TodoList{}
}

func (t TodoList) Get(c *gin.Context) {}
func (t TodoList) List(c *gin.Context) {
	param := service.TodoListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.PathParamsBindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTodo(&service.CountTodoListRequest{CreatedBy: param.CreatedBy, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTodo err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTodoFail)
		return
	}
	todos, err := svc.GetTodoList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTodoList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTodoListFail)
		return
	}

	response.ToResponseList(todos, totalRows)
	return
}
func (t TodoList) Create(c *gin.Context) {
	param := service.CreateTodoListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTodo(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTodo err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTodoFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
func (t TodoList) Update(c *gin.Context) {
	param := service.UpdateTodoListRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTodo(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTodo err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTodoFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
func (t TodoList) Delete(c *gin.Context) {
	param := service.DeleteTodoListRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTodo(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTodo err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTodoFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
