// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type User struct {
	ID              int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Username        string    `json:"username,omitempty" comment:"用户名"`
	Alias           string    `json:"alias,omitempty" comment:"昵称"`
	Avatar          string    `json:"avatar,omitempty" comment:"头像"`
	Email           string    `json:"email,omitempty" comment:"邮箱"`
	Mobile          string    `json:"mobile,omitempty" comment:"手机号码"`
	Password        string    `json:"password,omitempty" comment:"密码"`
	ConfirmPassword string    `json:"confirm_password,omitempty" comment:"确认密码"`
	Pwd             string    `json:"pwd,omitempty" comment:"数据库存储加密密码"`
	IsAdmin         bool      `json:"is_admin,omitempty" comment:"超级用户"`
	IsActive        bool      `json:"is_active,omitempty" comment:"用户有效"`
	Roles           []*Role   `json:"roles" comment:"用户角色"`
	RoleList        []int     `gorm:"-" json:"role_list,omitempty" comment:"角色ID列表"`
}

func (mh *User) TableName() string {
	return "users"
}
