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

type RoleService struct {
	DB *gorm.DB
	repositories.IRoleRepository
}

func (service *RoleService) ServiceGetRoleByID(id int) (result models.RoleDetail, err error) {
	return service.FindRoleByID(service.DB, id)
}

func (service *RoleService) ServiceRoleCreate(userCreate models.RoleCreate) (result models.RoleDetail, err error) {
	userCreate.ID = 0
	return service.InsertRole(service.DB, userCreate)

}
func (service *RoleService) ServiceRoleUpdate(update models.RoleUpdate) (result models.RoleDetail, err error) {
	return service.UpdateRole(service.DB, update)
}

func (service *RoleService) ServiceRoleDelete(ids []int) (err error) {
	return service.DeleteRole(service.DB, ids)
}
func (service *RoleService) ServiceGetAllRole(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.RoleList, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAllRole(service.DB, page, size, order, query, queryArgs)
	return
}
