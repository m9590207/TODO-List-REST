package model

import (
	"github.com/m9590207/TODO-List-REST/pkg/app"

	"gorm.io/gorm"
)

type TodoList struct {
	*Model
	Item  string `json:"item"`
	State uint8  `json:"state"`
}

func (t TodoList) TableName() string {
	return "todo_list"
}

type TodoListSwagger struct {
	List  []*TodoList
	Pager *app.Pager
}

func (t TodoList) Count(db *gorm.DB) (int64, error) {
	var count int64
	db = db.Where("created_by = ?", t.CreatedBy)
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t TodoList) List(db *gorm.DB, pageOffset, pageSize int) ([]*TodoList, error) {
	var todos []*TodoList
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("created_by = ?", t.CreatedBy)
	db = db.Where("state = ?", t.State)
	if err = db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (t TodoList) ListByIDs(db *gorm.DB, ids []uint32) ([]*TodoList, error) {
	var todos []*TodoList
	db = db.Where("state = ? ", t.State)
	err := db.Where("id IN (?)", ids).Find(&todos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return todos, nil
}

func (t TodoList) Get(db *gorm.DB) (TodoList, error) {
	var todo TodoList
	err := db.Where("id = ?  AND state = ?", t.ID, t.State).First(&todo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return todo, err
	}

	return todo, nil
}

func (t TodoList) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t TodoList) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? ", t.ID).Updates(values).Error
}

func (t TodoList) Delete(db *gorm.DB) error {
	return db.Where("id = ? ", t.Model.ID).Delete(&t).Error
}
