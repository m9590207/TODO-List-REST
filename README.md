# 待辦事項REST APIs 
* gin
* GORM 2.0
* middleware
  * 自訂log記錄資訊
  * 統一逾時控制
  * JWT
  * API限流
  * Jaeger 鏈路追蹤
* viper config管理
* docker 
  * 部署 Jaeger
  * 部署 MySql
## API
### 取得交易權杖
回傳JWT取得在權杖有效期限內進行其他API存取的權限，此功能使用github.com/juju/ratelimi提供的token bucket實作來實現api限流控管

**GET /auth/{appKey}/{appSecret}**
* Path params
  * appKey: 每次交易前用來獲取權杖 
  * appSecret: 同上  
  
### 查詢待辦事項清單
**GET /api/v1/{createdBy}/{state}/todos?page={int}&pageSize={int}**
* Headers
  * token: 傳入/auth 所取得的JWT
* Path params
  * createdBy: 待辦事項建立的人 
  * state: 0: 待處理、 1: 已完成
* Query params
  * page: 分頁第幾頁 
  * pageSize: 每頁最大筆數
  
### 建立待辦事項
**POST /api/v1/todo**
* Headers
  * token: 傳入/auth 所取得的JWT
  * Content-Type: multipart/form-data
* Body params
  * item: 待辦事項內容 
  * createdBy: 待辦事項建立的人

### 刪除待辦事項
**DELETE /api/v1/todo/{id}**
* Headers
  * token: 傳入/auth 所取得的JWT
* Path params
  * id: 待辦事項的id 

### 更新待辦事項
**PUT /api/v1/todo/{id}**
* Headers
  * token: 傳入/auth 所取得的JWT
  * Content-Type: multipart/form-data
* Path params
  * id: 待辦事項的id 
* Body params
  * item: 待辦事項內容 
  	
### 更新待辦事項狀態
**PATCH /api/v1/todo/{id}**
* Headers
  * token: 傳入/auth 所取得的JWT
  * Content-Type: multipart/form-data
* Path params
  * id: 待辦事項的id 
* Body params
  * state: 0: 待處理、 1: 已完成 

