// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type MenuService struct {
	DB *gorm.DB
	repositories.IMenuRepository
}

func (service *MenuService) ServiceGetSelfMenu(roles []string) (results []*models.MenuList, err error) {
	return service.FindMenuByRoles(service.DB, roles)
}
func (service *MenuService) ServiceGetMenuByID(id string) (result models.MenuDetail, err error) {
	return service.FindMenuByID(service.DB, id)
}
func (service *MenuService) ServiceMenuCreate(userCreate models.MenuCreate) (result models.MenuDetail, err error) {
	userCreate.ID = 0
	return service.InsertMenu(service.DB, userCreate)

}
func (service *MenuService) ServiceMenuUpdate(update models.MenuUpdate) (result models.MenuDetail, err error) {
	return service.UpdateMenu(service.DB, update)
}

func (service *MenuService) ServiceMenuDelete(ids []string) (err error) {
	return service.DeleteMenu(service.DB, ids)
}
func (service *MenuService) ServiceGetAllMenu(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.MenuList, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAllMenu(service.DB, page, size, order, query, queryArgs)
	return
}
