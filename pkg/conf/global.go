// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/casbin/casbin"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path"
)
type ResponseJson struct {
	Code    uint        `yaml:"code" json:"code"`       // response code
	Data    interface{} `yaml:"data" json:"data"`       // response data
	Message string      `yaml:"message" json:"message"` // response message
}
var (
	// 数据库连接对象
	DBConn *gorm.DB
	SqlxDB *sqlx.DB

	// Config global parameters
	Config GlobalConfig

	// casbin
	Enforcer *casbin.Enforcer
	// 全局语言对象
	I18nBundles = make(map[string]*i18n.Bundle)
)

//BundleI18nLanguages 加载所有可支持的语言资源
func BundleI18nLanguages() {

	if len(Config.I18ns) < 1 {
		Config.I18ns = append(Config.I18ns, "en")
	}
	var langTag language.Tag
	for _, lang := range Config.I18ns {
		switch lang {
		case "zh-CN":
		case "zh":
			langTag = language.SimplifiedChinese

		case "en":
			langTag = language.English
		default:
			langTag = language.English

		}
		bundle := &i18n.Bundle{DefaultLanguage: langTag}
		// i18n文件类型
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		//加载文件
		cwd, _ := os.Getwd()
		i18nPath := path.Join(cwd, "pkg", "i18n", tools.StringsJoin(lang, ".toml"))
		bundle.MustLoadMessageFile(i18nPath)
		I18nBundles[lang] = bundle

	}
}
