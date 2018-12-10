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
func (repo *MenuRepository) InsertMenu(DB *gorm.DB, create models.MenuCreate) (result models.MenuDetail, err error) {
	err = DB.Create(&create).Error
	if err == nil {
		return repo.FindMenuByID(DB, strconv.Itoa(create.ID))
	}
	return
}

// 修改
func (repo *MenuRepository) UpdateMenu(DB *gorm.DB, menu models.MenuUpdate) (result models.MenuDetail, err error) {
	err = DB.Updates(menu).Error
	return
}

// 获得所有菜单
func (repo *MenuRepository) FindAllMenu(DB *gorm.DB) (menus []*models.MenuList, err error) {
	err = DB.Find(&menus).Error
	return
}
