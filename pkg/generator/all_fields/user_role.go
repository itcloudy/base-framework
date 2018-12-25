// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type UserRole struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Role      *Role     `json:"role" gorm:"-" comment:"角色"`
	RoleID    int       `json:"role_id" gorm:"column:role_id" comment:"角色ID"`
	User      *User     `json:"user" gorm:"-" comment:"用户"`
	UserID    int       `json:"user_id" gorm:"column:user_id" comment:"用户ID"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *UserRole) TableName() string {
	return "user_role"
}
