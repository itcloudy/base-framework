// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type Role struct {
	ID             int         `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt      time.Time   `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Name           string      `json:"name" gorm:"column:name" binding:"required" comment:"角色名称"`
	Code           string      `json:"code" gorm:"column:code" comment:"角色编码"`
	InheritIds     []int       `json:"inherit_ids,omitempty" gorm:"-" comment:"所继承的角色ID"`
	InheritStrings string      `json:"inherit_strings" gorm:"column:inherit_strings" comment:"所继承角色ID逗号分隔"`
	Inherits       []*Role     `json:"inherits,omitempty" gorm:"-" comment:"继承的角色"`
	RoleMenus      []*RoleMenu `json:"role_menus" gorm:"-" comment:"角色拥有的菜单"`
	IsActive       bool        `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *Role) TableName() string {
	return "menu"
}
