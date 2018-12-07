// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

type Role struct {
	ID         int         `json:"id" comment:"主键ID"`
	CreatedAt  time.Time   `json:"created_at,omitempty" comment:"记录创建时间"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty" comment:"记录更新时间"`
	Name       string      `json:"name" db:"not null;unique_index" binding:"required" comment:"角色名称"`
	Code       string      `json:"code" db:"not null;unique_index" comment:"角色编码"`
	InheritIds []int       `json:"inherit_ids,omitempty" db:"-" comment:"所继承的角色ID"`
	Inherits   []*Role     `json:"inherits,omitempty" db:"-" comment:"继承的角色"`
	RoleMenus  []*RoleMenu `json:"role_menus" yaml:"role_menus" comment:"角色拥有的菜单"`
}
