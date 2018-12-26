// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"{{.ProjectPath}}/pkg/conf"
	"{{.ProjectPath}}/pkg/models"
)

type I{{.ModelName}}Service interface {
	ServiceGet{{.ModelName}}ByID(id int) (result models.{{.ModelName}}Detail, err error)
	Service{{.ModelName}}Create(model models.{{.ModelName}}Create) (result models.{{.ModelName}}Detail, err error)
	Service{{.ModelName}}Update(update models.{{.ModelName}}Update) (result models.{{.ModelName}}Detail, err error)
	Service{{.ModelName}}Delete(ids []int) (err error)
	ServiceGetAll{{.ModelName}}(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.{{.ModelName}}List, pagination conf.Pagination, err error)
}
