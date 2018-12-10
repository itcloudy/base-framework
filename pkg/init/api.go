// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package init

import (
	"errors"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/services"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type systemApis struct {
	APIs []models.SystemApiCreate `yml:"apis" yml:"apis"`
}

func initApi() {
	var systemapis *systemApis
	filePath := conf.Config.Init.API
	if filePath == "" {
		logs.Logger.Error("api init file path is empty", zap.String("path", filePath))
		panic("api init file path is empty")
	}
	apiData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Logger.Error("api init file read filed", zap.String("path", filePath))

		panic("api init file read filed, file path: " + filePath)
	}
	err = yaml.Unmarshal(apiData, &systemapis)
	if err != nil {
		logs.Logger.Error("api init file parser filed", zap.String("path", filePath))
		panic("api init file parser filed, err information: ")
	}
	apiService := services.SystemAPIService{}
	dbType := conf.Config.DB.DbType
	switch dbType {
	case "mysql":
	case "postgres":
		apiService.DB = conf.DBConn
		break
	default:
		panic(errors.New("un support sql type:" + dbType))
	}
	apis, _ := apiService.ServiceGetAllSystemAPI()
	if len(apis) > 0 {
		logs.Logger.Error("system api not empty can't init", zap.String("path", filePath))
		panic("system api not empty can't init")
	} else {
		insertApis(systemapis, &apiService)
	}

}
func insertApis(apis *systemApis, service *services.SystemAPIService) {
	var (
		err error
	)
	for _, a := range apis.APIs {
		_, err = service.InsertSystemAPI(service.DB, a)
		if err != nil {
			logs.Logger.Error("system api create failed", zap.Error(err), zap.String("api", a.Display))
			logs.Logger.Sync()
			os.Exit(-1)
		}
	}
}
