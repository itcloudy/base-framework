// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package init

import (
	"errors"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories/common"
	"github.com/itcloudy/base-framework/pkg/services"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type systemMenus struct {
	Menus []models.MenuDetail `yaml:"menus" json:"menus"`
}

func initMenus() {
	filePath := conf.Config.Init.Menu
	var systemMs *systemMenus
	if filePath == "" {
		logs.Logger.Error("menu init file path is empty", zap.String("path", filePath))
		os.Exit(-1)
	}
	//读取菜单文件
	menuData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Logger.Error("menu init file read filed", zap.String("path", filePath))
		os.Exit(-1)
	}
	//映射菜单
	err = yaml.Unmarshal(menuData, &systemMs)
	if err != nil {
		logs.Logger.Error("menu init file parser filed", zap.String("path", filePath))
		os.Exit(-1)
	}
	menuService := services.MenuService{}
	menuService.IMenuRepository = &common.MenuRepository{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		menuService.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))

	}
	//如果系统菜单数据不为空则不能导入
	menus, _ := menuService.ServiceGetAllMenu()
	if len(menus) > 0 {
		logs.Logger.Error("system menu not empty can't init")
		os.Exit(-1)
	}
	//导入数据
	for _, m := range systemMs.Menus {

		insertMenus(&m,&menuService)
	}
}
func insertMenus(men *models.MenuDetail, service *services.MenuService) {
	var (
		sysMenu models.MenuCreate
		m       models.MenuDetail
		err     error
	)
	sysMenu.Name = men.Name
	sysMenu.Component = men.Component
	sysMenu.Icon = men.Icon
	sysMenu.Route = men.Route
	sysMenu.ParentID = men.ParentID
	sysMenu.Sequence = men.Sequence
	sysMenu.UniqueTag = men.UniqueTag
	m, err = service.InsertMenu(service.DB,sysMenu)
	if err != nil {
		logs.Logger.Error("system menu create failed", zap.Error(err), zap.String("unique tag", m.UniqueTag))
		logs.Logger.Sync()
		os.Exit(-1)
	}

	for _, me := range men.Children {
		me.ParentID = m.ID
		insertMenus(me, service)
	}

}
