// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IMenuRepository interface {
	//获得某些角色的菜单信息
	FindMenuByRoles(DB *gorm.DB, roleSlice []string) (menus []*models.MenuList, err error)
	//根据ID查找
	FindMenuByID(DB *gorm.DB, id string) (menu models.MenuDetail, err error)
	// 创建菜单
	InsertMenu(DB *gorm.DB, menu models.MenuCreate) (models.MenuDetail, error)
	// 菜单修改
	UpdateMenu(DB *gorm.DB, menu models.MenuUpdate) (models.MenuDetail, error)
	// 删除菜单
	DeleteMenu(DB *gorm.DB, ids []string) error
	// 查询菜单
	FindAllMenu(DB *gorm.DB, offset, limit int, order string, query string, queryArgs ...interface{}) (menus []*models.MenuList, count int, err error)
}
