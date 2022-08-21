package dao

import "github.com/m9590207/TODO-List-REST/internal/model"

//取得todo_auth的資料
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
