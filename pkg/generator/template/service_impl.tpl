// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"{{.ProjectPath}}/pkg/conf"
	"{{.ProjectPath}}/pkg/interfaces/repositories"
	"{{.ProjectPath}}/pkg/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type {{.ModelName}}Service struct {
	DB *gorm.DB
	repositories.I{{.ModelName}}Repository
}

func (service *{{.ModelName}}Service) ServiceGet{{.ModelName}}ByID(id int) (result models.{{.ModelName}}Detail, err error) {
	return service.Find{{.ModelName}}ByID(service.DB, id)
}
func (service *{{.ModelName}}Service) Service{{.ModelName}}Delete(ids []int) (err error) {
	return service.Delete{{.ModelName}}(service.DB, ids)
}
func (service *{{.ModelName}}Service) Service{{.ModelName}}Update(model models.{{.ModelName}}Update) (result models.{{.ModelName}}Detail, err error) {
	return service.Update{{.ModelName}}(service.DB, model)
}
func (service *{{.ModelName}}Service) Service{{.ModelName}}Create(model models.{{.ModelName}}Create) (result models.{{.ModelName}}Detail, err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(model); err != nil {
		return
	}
	return service.Insert{{.ModelName}}(service.DB, model)
}

func (service *{{.ModelName}}Service) ServiceGetAll{{.ModelName}}(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.{{.ModelName}}List, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAll{{.ModelName}}(service.DB, page, size, order, query, queryArgs)
	return
}
