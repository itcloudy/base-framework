// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

//用户
type User struct {
	ID              int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Username        string    `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Alias           string    `json:"alias,omitempty"  gorm:"column:alias" comment:"昵称"`
	Avatar          string    `json:"avatar,omitempty" gorm:"column:avatar" comment:"头像"`
	Email           string    `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile          string    `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
	Password        string    `json:"password,omitempty" comment:"密码"`
	ConfirmPassword string    `json:"confirm_password,omitempty" comment:"确认密码"`
	Pwd             string    `json:"pwd,omitempty" gorm:"column:pwd"  comment:"数据库存储加密密码"`
	IsAdmin         bool      `json:"is_admin,omitempty" gorm:"column:is_admin" comment:"超级用户"`
	IsActive        bool      `json:"is_active,omitempty" gorm:"column:is_active" comment:"用户有效"`
	RoleID          int       `json:"role_id" gorm:"column:role_id" comment:"用户角色ID"`
	Role            *Role     `json:"role" comment:"用户角色"`
}

func (mh *User) TableName() string {
	return "users"
}
