// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type MenuRepository struct {
}

func (repo *MenuRepository) FindMenuByRoles(DB *gorm.DB, roleSlice []string) (menus []*models.MenuList, err error) {
	err = DB.Where("role_id IN (?)", roleSlice).Find(&menus).Error
	return
}

//根据ID查找
func (repo *MenuRepository) FindMenuByID(DB *gorm.DB, id string) (menu models.MenuDetail, err error) {
	err = DB.Where("id = ?", id).First(&menu).Error
	return
}

// 创建菜单
func (repo *MenuRepository) InsertMenu(DB *gorm.DB, model models.MenuCreate) (result models.MenuDetail, err error) {
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.FindMenuByID(DB, strconv.Itoa(model.ID))
	}
	return
}

// 修改
func (repo *MenuRepository) UpdateMenu(DB *gorm.DB, menu models.MenuUpdate) (result models.MenuDetail, err error) {
	err = DB.Updates(menu).Error
	return
}
func (repo *MenuRepository) DeleteMenu(DB *gorm.DB, ids []string) (err error) {
	return DB.Where("id IN (?)", ids).Delete(models.MenuDetail{}).Error
}

// 查询菜单
func (repo *MenuRepository) FindAllMenu(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (menus []*models.MenuList, count int, err error) {
	if len(order) == 0 {
		order = "id desc"
	}
	err = DB.Order(order).Offset((page - 1) * size).Limit(size).Find(&menus).Error
	return
}
