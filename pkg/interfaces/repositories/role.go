// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IRoleRepository interface {
	//根据ID查找
	FindRoleByID(DB *gorm.DB, id int) (role models.RoleDetail, err error)
	// 创建角色
	InsertRole(DB *gorm.DB, role models.RoleCreate) (result models.RoleDetail, err error)
	// 修改角色
	UpdateRole(DB *gorm.DB, role models.RoleUpdate) (result models.RoleDetail, err error)
	// 删除角色
	DeleteRole(DB *gorm.DB, ids []int) error
	// 查询角色
	FindAllRole(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (roles []*models.RoleList, total int, err error)
}
