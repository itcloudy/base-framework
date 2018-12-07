// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type Menu struct {
	ID        int       `json:"id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" comment:"记录更新时间"`
	Parent    *Menu     `json:"parent,omitempty" comment:"上级菜单"`
	ParentID  int       `json:"parent_id" comment:"上级菜单ID"`
	Name      string    `json:"name" db:"index" yaml:"name" validate:"required" comment:"菜单名称"`
	Route     string    `json:"route,omitempty" yaml:"route" comment:"菜单路由"`
	Component string    `json:"component,omitempty" yaml:"component" comment:"菜单组件"`
	Icon      string    `json:"icon,omitempty" yaml:"icon" validate:"required" comment:"菜单样式类"`
	Sequence  int       `json:"sequence" yaml:"sequence" validate:"required" comment:"菜单顺序"`
	Tree      string    `json:"-" yaml:"tree" comment:"菜单继承树"`
	Children  []*Menu   `json:"children,omitempty" yaml:"children" comment:"子菜单"`
	UniqueTag string    `json:"-" db:"unique_index" validate:"required" comment:"菜单唯一标识"`
}
