// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

type RoleCreate struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Name      string    `json:"name" gorm:"column:name" binding:"required" comment:"角色名称"`
	Code      string    `json:"code" gorm:"column:code" comment:"角色编码"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *RoleCreate) TableName() string {
	return "role"
}

type RoleUpdate struct {
	ID       int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Name     string `json:"name" gorm:"column:name" binding:"required" comment:"角色名称"`
	IsActive bool   `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *RoleUpdate) TableName() string {
	return "role"
}

type RoleList struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Name      string    `json:"name" gorm:"column:name" binding:"required" comment:"角色名称"`
	Code      string    `json:"code" gorm:"column:code" comment:"角色编码"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *RoleCreate) RoleList() string {
	return "role"
}

type RoleDetail struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Name      string    `json:"name" gorm:"column:name" binding:"required" comment:"角色名称"`
	Code      string    `json:"code" gorm:"column:code" comment:"角色编码"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *RoleDetail) TableName() string {
	return "role"
}
