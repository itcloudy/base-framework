// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mysql

import "strings"

var Init = strings.Replace(`
CREATE TABLE users (
id int(11) NOT NULL COMMENT '主键',
username varchar(255) NOT NULL COMMENT '用户名',
password varchar(255) NULL COMMENT '密码',
avatar varchar(255) NULL COMMENT '头像',
is_active tinyint(1) NULL DEFAULT 1 COMMENT '有效',
is_admin tinyint(1) NULL COMMENT '超级用户',
create_at datetime NULL COMMENT '创建时间',
update_at datetime NULL COMMENT '更新时间',
PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT = '用户表';

CREATE TABLE user_profile (
id int(11) NOT NULL COMMENT '主键',
create_at datetime NULL COMMENT '创建时间',
update_at datetime NULL COMMENT '更新时间',
user_id int(11) NULL COMMENT '用户ID',
PRIMARY KEY (id)
);


ALTER TABLE user_profile ADD CONSTRAINT fk_user_profile FOREIGN KEY (user_id) REFERENCES user (id);

`, `\n`, " ", -1)
