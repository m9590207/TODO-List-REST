CREATE DATABASE
IF 
  NOT EXISTS todo_list_service DEFAULT CHARACTER
  SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE todo_list_service;
CREATE TABLE todo_list (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  item varchar(2000) DEFAULT '' COMMENT '待辦事項內容',
  date_added int(10) unsigned DEFAULT '0' COMMENT '加入時間',
  created_by varchar(100) DEFAULT '' COMMENT '建立人',
  date_modified int(10) unsigned DEFAULT '0' COMMENT '更新時間', 
  `state` tinyint(3) unsigned DEFAULT '0' COMMENT '狀態 0待處理、1已完成',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='待辦事項';

CREATE TABLE todo_auth (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  app_key varchar(20) DEFAULT '' COMMENT 'Key',
  app_secret varchar(50) DEFAULT '' COMMENT 'Secret',
  date_added int(10) unsigned DEFAULT '0' COMMENT '建立時間',
  created_by varchar(100) DEFAULT '' COMMENT '建立人',   
  date_modified int(10) unsigned DEFAULT '0' COMMENT '刪除時間',  
  is_del tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0未刪除、1已刪除',
  PRIMARY KEY (id) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='認證管理';

INSERT INTO todo_auth(id, app_key, app_secret, date_added, created_by, date_modified, is_del)
VALUES(1, 'appkey1','appsecret1', UNIX_TIMESTAMP(),'todolistapp', UNIX_TIMESTAMP(),0);