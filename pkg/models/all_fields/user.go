// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type User struct {
	ID              int       `json:"id" comment:"主键ID"`
	CreatedAt       time.Time `json:"created_at,omitempty" comment:"记录创建时间"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" comment:"记录更新时间"`
	Username        string    `comment:"用户名"`
	Alias           string    `comment:"昵称"`
	HeadImage       string    `comment:"头像"`
	Email           string    `comment:"邮箱"`
	Mobile          string    `comment:"手机号码"`
	Password        string    `comment:"密码"`
	ConfirmPassword string    `comment:"确认密码"`
	Pwd             string    `comment:"数据库存储加密密码"`
	IsAdmin         bool      `comment:"超级用户"`
	IsActive        bool      `comment:"用户有效"`
}
