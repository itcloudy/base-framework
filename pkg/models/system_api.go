// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

type SystemApiCreate struct {
	ID      int    `json:"id" yaml:"id" gorm:"column:id" comment:"主键ID"`
	Name    string `json:"name" yaml:"name" gorm:"column:name" validate:"required" comment:"API名称"`
	Address string `json:"address" yaml:"address" gorm:"column:address" validate:"required" comment:"API地址"`
	Method  string `json:"method" yaml:"method" gorm:"column:method" validate:"required" comment:"API请求方法"`
	Display string `json:"display" yaml:"-" gorm:"column:display" comment:"显示名称"`
}

func (mh *SystemApiCreate) TableName() string {
	return "system_api"
}

type SystemApiUpdate struct {
	ID      int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Name    string `json:"name" yaml:"name" gorm:"column:name" validate:"required" comment:"API名称"`
	Address string `json:"address" yaml:"address" gorm:"column:address" validate:"required" comment:"API地址"`
	Method  string `json:"method" yaml:"method" gorm:"column:method" validate:"required" comment:"API请求方法"`
	Display string `json:"display" yaml:"-" gorm:"column:display" comment:"显示名称"`
}

func (mh *SystemApiUpdate) TableName() string {
	return "system_api"
}

type SystemApiDetail struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Name      string    `json:"name" yaml:"name" gorm:"column:name" comment:"API名称"`
	Address   string    `json:"address" yaml:"address" gorm:"column:address" comment:"API地址"`
	Method    string    `json:"method" yaml:"method" gorm:"column:method" comment:"API请求方法"`
	Display   string    `json:"display" yaml:"-" gorm:"column:display" comment:"显示名称"`
	IsActive  bool      `json:"is_active" yaml:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *SystemApiDetail) TableName() string {
	return "system_api"
}

type SystemApiList struct {
	ID       int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Name     string `json:"name" yaml:"name" gorm:"column:name" comment:"API名称"`
	Address  string `json:"address" yaml:"address" gorm:"column:address" comment:"API地址"`
	Method   string `json:"method" yaml:"method" gorm:"column:method" comment:"API请求方法"`
	Display  string `json:"display" yaml:"-" gorm:"column:display" comment:"显示名称"`
	IsActive bool   `json:"is_active" yaml:"is_active" gorm:"column:is_active" comment:"有效"`
}

func (mh *SystemApiList) TableName() string {
	return "system_api"
}
