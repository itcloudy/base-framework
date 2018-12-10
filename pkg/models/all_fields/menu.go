// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type Menu struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Parent    *Menu     `json:"parent,omitempty" gorm:"-" comment:"上级菜单"`
	ParentID  int       `json:"parent_id" gorm:"column:parent_id" comment:"上级菜单ID"`
	Name      string    `json:"name" gorm:"column:name" validate:"required" comment:"菜单名称"`
	Route     string    `json:"route,omitempty" gorm:"column:route" comment:"菜单路由"`
	Component string    `json:"component,omitempty" gorm:"column:component"  comment:"菜单组件"`
	Icon      string    `json:"icon,omitempty" validate:"required" comment:"菜单样式类"`
	Sequence  int       `json:"sequence" gorm:"column:sequence"  validate:"required" comment:"菜单顺序"`
	Tree      string    `json:"-" gorm:"column:tree" comment:"菜单继承树"`
	Children  []*Menu   `json:"children,omitempty" gorm:"-" comment:"子菜单"`
	UniqueTag string    `json:"-" gorm:"column:unique_tag" validate:"required" comment:"菜单唯一标识"`
}

func (mh *Menu) TableName() string {
	return "menu"
}
