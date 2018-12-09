// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
)

type MenuService struct {
	repositories.IMenuRepository
}

func (service *MenuService) GetSelfMenu(roles []string) (menus []models.Menu, err error) {
	return
}
