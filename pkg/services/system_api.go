// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type SystemAPIService struct {
	DB *gorm.DB
	repositories.ISystemAPIRepository
}

func (service *SystemAPIService) ServiceGetAllSystemAPI(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.SystemApiList, count int, err error) {
	return service.FindAllSystemAPI(service.DB, page, size, order, query, queryArgs)
}
func (service *SystemAPIService) ServiceGetSystemAPIByID(id string) (result models.SystemApiDetail, err error) {
	return service.FindSystemAPIByID(service.DB, id)
}
func (service *SystemAPIService) ServiceSystemAPIDelete(ids []string) (err error) {
	return service.DeleteSystemAPI(service.DB, ids)
}
func (service *SystemAPIService) ServiceSystemAPIUpdate(model models.SystemApiUpdate) (result models.SystemApiDetail, err error) {
	return service.UpdateSystemAPI(service.DB, model)
}
func (service *SystemAPIService) ServiceSystemAPICreate(model models.SystemApiCreate) (result models.SystemApiDetail, err error) {
	return service.InsertSystemAPI(service.DB, model)
}
func (service *SystemAPIService) ServiceActiveSystemAPI(ids []string, active bool) (err error) {
	return service.ActiveSystemAPI(service.DB, ids, active)

}
