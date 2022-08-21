package service

import (
	"github.com/m9590207/TODO-List-REST/internal/model"
	"github.com/m9590207/TODO-List-REST/pkg/app"
)

//傳入參數定義 gin內建的模型綁定跟驗證使用go-playground/validator

type CountTodoListRequest struct {
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TodoListRequest struct {
	CreatedBy string `uri:"createdBy" binding:"required,min=2,max=100"`
	State     uint8  `uri:"state,default=1" binding:"oneof=0 1"`
}

type CreateTodoListRequest struct {
	Item      string `form:"item" binding:"max=2000"`
	CreatedBy string `form:"createdBy" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=0" binding:"oneof=0 1"`
}

type UpdateTodoListRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	Item  string `form:"item" binding:"max=2000"`
	State uint8  `form:"state" binding:"oneof=0 1"`
}

type DeleteTodoListRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTodo(param *CountTodoListRequest) (int64, error) {
	return svc.dao.CountTodo(param.CreatedBy, param.State)
}

func (svc *Service) GetTodoList(param *TodoListRequest, pager *app.Pager) ([]*model.TodoList, error) {
	return svc.dao.GetTodoList(param.CreatedBy, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTodo(param *CreateTodoListRequest) error {
	return svc.dao.CreateTodo(param.Item, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTodo(param *UpdateTodoListRequest) error {
	return svc.dao.UpdateTodo(param.ID, param.Item, param.State)
}

func (svc *Service) DeleteTodo(param *DeleteTodoListRequest) error {
	return svc.dao.DeleteTodo(param.ID)
}
