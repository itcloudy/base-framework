// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

type RoleCreate struct {
	ID int `json:"id" gorm:"column:id" comment:"主键ID"`
}

func (mh *RoleCreate) TableName() string {
	return "role"
}

type RoleUpdate struct {
	ID int `json:"id" gorm:"column:id" comment:"主键ID"`
}

func (mh *RoleUpdate) TableName() string {
	return "role"
}

type RoleList struct {
	ID int `json:"id" gorm:"column:id" comment:"主键ID"`
}

func (mh *RoleList) TableName() string {
	return "role"
}

type RoleDetail struct {
	ID int `json:"id" gorm:"column:id" comment:"主键ID"`
}

func (mh *RoleDetail) TableName() string {
	return "role"
}
