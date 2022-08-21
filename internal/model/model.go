package model

import (
	"fmt"
	"time"

	"github.com/m9590207/TODO-List-REST/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

//公用欄位
type Model struct {
	ID           uint32 `gorm:"primary_key" json:"id"`
	DateAdded    uint32 `json:"date_added"`
	CreatedBy    string `json:"created_by"`
	DateModified uint32 `json:"date_modified"`
}

//根據自己的需要自訂gorm的call back操作
//統一處理共用欄位, 取代現有的call back
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用單數表名
		},
	})
	if err != nil {
		return nil, err
	} else {
		dbConfig, _ := db.DB()
		dbConfig.SetMaxOpenConns(databaseSetting.MaxOpenConns)
		dbConfig.SetMaxIdleConns(databaseSetting.MaxIdleConns)
		dbConfig.SetConnMaxLifetime(time.Hour)
	}
	db.Callback().Create().Replace("gorm:before_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:before_update", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	nowTime := time.Now().Unix()
	DateAddedField := db.Statement.Schema.LookUpField("DateAdded")
	if DateAddedField != nil {
		db.Statement.SetColumn("date_added", nowTime)
	}

	DateModifiedField := db.Statement.Schema.LookUpField("DateModified")
	if DateModifiedField != nil {
		db.Statement.SetColumn("date_modified", nowTime)
	}

}
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	DateModifiedField := db.Statement.Schema.LookUpField("DateModified")
	if DateModifiedField != nil {
		db.Statement.SetColumn("date_modified", time.Now().Unix())
	}
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if db.Statement.Schema != nil {
			db.Statement.SQL.Grow(100)

			deleteField := db.Statement.Schema.LookUpField("IsDel")
			if !db.Statement.Unscoped && deleteField != nil {
				//軟刪除 sql
				if db.Statement.SQL.String() == "" {
					db.Statement.AddClause(
						clause.Set{{
							Column: clause.Column{Name: deleteField.DBName},
							Value:  1,
						}},
					)
					db.Statement.AddClauseIfNotExists(clause.Update{})
					db.Statement.Build("UPDATE", "SET", "WHERE")
				}
			} else {
				//Delete sql如果不是軟刪除
				if db.Statement.SQL.String() == "" {
					db.Statement.AddClauseIfNotExists(clause.Delete{})
					db.Statement.AddClauseIfNotExists(clause.From{})
					db.Statement.Build("DELETE", "FROM", "WHERE")
				}
			}

			//檢查 WHERE 子句
			if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok {
				db.AddError(gorm.ErrMissingWhereClause)
				return
			}

			//執行SQL
			if !db.DryRun && db.Error == nil {
				result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
				if err == nil {
					db.RowsAffected, _ = result.RowsAffected()
				} else {
					db.AddError(err)
				}
			}
		}
	}
}
