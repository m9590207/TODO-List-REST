package dao

import (
	"github.com/m9590207/TODO-List-REST/internal/model"
	"github.com/m9590207/TODO-List-REST/pkg/app"
)

func (d *Dao) CountTodo(createdBy string, state uint8) (int64, error) {
	todo := model.TodoList{Model: &model.Model{CreatedBy: createdBy}, State: state}
	return todo.Count(d.engine)
}

func (d *Dao) GetTodoList(createdBy string, state uint8, page, pageSize int) ([]*model.TodoList, error) {
	todo := model.TodoList{Model: &model.Model{CreatedBy: createdBy}, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return todo.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTodo(id uint32, state uint8) (model.TodoList, error) {
	todo := model.TodoList{Model: &model.Model{ID: id}, State: state}
	return todo.Get(d.engine)
}

func (d *Dao) GetTodoByIDs(ids []uint32, state uint8) ([]*model.TodoList, error) {
	todo := model.TodoList{State: state}
	return todo.ListByIDs(d.engine, ids)
}

func (d *Dao) CreateTodo(item string, state uint8, createdBy string) error {
	todo := model.TodoList{
		Item:  item,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return todo.Create(d.engine)
}

func (d *Dao) UpdateTodo(id uint32, item string, state uint8) error {
	todo := model.TodoList{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state": state,
	}
	if item != "" {
		values["item"] = item
	}

	return todo.Update(d.engine, values)
}

func (d *Dao) DeleteTodo(id uint32) error {
	todo := model.TodoList{Model: &model.Model{ID: id}}
	return todo.Delete(d.engine)
}
