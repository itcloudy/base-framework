// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IRoleRepository interface {
	//根据ID查找
	FindMenuByID(DB *gorm.DB, id string) (menu models.RoleDetail, err error)
	// 创建角色
	InsertMenu(DB *gorm.DB, role models.RoleCreate) (models.RoleDetail, error)
	// 修改角色
	UpdateMenu(DB *gorm.DB, role models.RoleUpdate) (models.RoleDetail, error)
	// 获得所有角色
	FindAllRole(DB *gorm.DB) (roles []*models.RoleList, err error)
}
