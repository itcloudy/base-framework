// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package models

import "time"

type User struct {
	ID              int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Username        string    `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Alias           string    `json:"alias,omitempty" gorm:"column:alias" comment:"昵称"`
	HeadImage       string    `json:"head_image,omitempty" gorm:"column:head_image" comment:"头像"`
	Email           string    `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile          string    `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
	Password        string    `json:"password,omitempty" gorm:"-:" comment:"密码"`
	ConfirmPassword string    `json:"confirm_password,omitempty" gorm:"-:" comment:"确认密码"`
	Pwd             string    `json:"pwd,omitempty" gorm:"column:pwd" comment:"数据库存储加密密码"`
	IsAdmin         bool      `json:"is_admin" gorm:"column:is_admin" comment:"超级用户"`
	IsActive        bool      `json:"is_active" gorm:"column:is_active" comment:"用户有效"`
}

// TableName returns name of table
func (mh *User) TableName() string {
	return "users"
}

//  用户创建
type UserCreate struct {
	ID              int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Username        string `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Alias           string `json:"alias,omitempty" gorm:"column:alias" comment:"昵称"`
	HeadImage       string `json:"head_image,omitempty" gorm:"column:head_image" comment:"头像"`
	Email           string `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile          string `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
	Password        string `json:"password,omitempty" gorm:"-:" comment:"密码"`
	ConfirmPassword string `json:"confirm_password,omitempty" gorm:"-:" comment:"确认密码"`
	Pwd             string `json:"pwd,omitempty" gorm:"column:pwd" comment:"数据库存储加密密码"`
	IsAdmin         bool   `json:"is_admin" gorm:"column:is_admin" comment:"超级用户"`
	IsActive        bool   `json:"is_active" gorm:"column:is_active" comment:"用户有效"`
}

func (mh *UserCreate) TableName() string {
	return "users"
}

//用户更新
type UserUpdate struct {
	ID        int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Alias     string `json:"alias,omitempty" gorm:"column:alias" comment:"昵称"`
	HeadImage string `json:"head_image,omitempty" gorm:"column:head_image" comment:"头像"`
	Email     string `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile    string `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
}

func (mh *UserUpdate) TableName() string {
	return "users"
}

//用户详情
type UserDetail struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Username  string    `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Alias     string    `json:"alias,omitempty" gorm:"column:alias" comment:"昵称"`
	HeadImage string    `json:"head_image,omitempty" gorm:"column:head_image" comment:"头像"`
	Email     string    `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile    string    `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
	Pwd       string    `json:"-" gorm:"column:pwd" comment:"数据库存储加密密码"`
	IsAdmin   bool      `json:"is_admin" gorm:"column:is_admin" comment:"超级用户"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"用户有效"`
}

func (mh *UserDetail) TableName() string {
	return "users"
}

//用户列表
type UserList struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Username  string    `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Alias     string    `json:"alias,omitempty" gorm:"column:alias" comment:"昵称"`
	HeadImage string    `json:"head_image,omitempty" gorm:"column:head_image" comment:"头像"`
	Email     string    `json:"email,omitempty" gorm:"column:email" comment:"邮箱"`
	Mobile    string    `json:"mobile,omitempty" gorm:"column:mobile" comment:"手机号码"`
	IsAdmin   bool      `json:"is_admin" gorm:"column:is_admin" comment:"超级用户"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active" comment:"用户有效"`
}

func (mh *UserList) TableName() string {
	return "users"
}

//用户登录和密码修改
type UserAuth struct {
	ID              int    `json:"id" gorm:"column:id" comment:"主键ID"`
	Username        string `json:"username,omitempty" gorm:"column:username" comment:"用户名"`
	Password        string `json:"password,omitempty" gorm:"-:" comment:"密码"`
	ConfirmPassword string `json:"confirm_password,omitempty" gorm:"-:" comment:"确认密码"`
}

func (mh *UserAuth) TableName() string {
	return "users"
}
