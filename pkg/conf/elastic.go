// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package conf

import (
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
)

func GetElasticClient(config ElasticConfig) {
	logs.Logger.Info("ready connect elastic", zap.String("version", elastic.Version), zap.Strings("urls", config.URLs))
	var (
		err     error
		options []elastic.ClientOptionFunc
	)
	options = append(options, elastic.SetURL(config.URLs...))
	if config.AuthUserName != "" {
		logs.Logger.Info("elastic connect info: need auth")
		options = append(options, elastic.SetBasicAuth(config.AuthUserName, config.AuthPassword))
	} else {
		logs.Logger.Info("elastic connect info: no auth")
	}
	if config.Sniff {
		logs.Logger.Info("elastic connect info: enable sniff")
		options = append(options, elastic.SetSniff(config.Sniff), elastic.SetSnifferInterval(config.Sniffer))
	} else {
		logs.Logger.Info("elastic connect info: disable sniff")
	}
	if config.HealthCheck {
		logs.Logger.Info("elastic connect info: enable health check")
		options = append(options, elastic.SetHealthcheck(config.HealthCheck), elastic.SetHealthcheckInterval(config.HealthChecker))
	} else {
		logs.Logger.Info("elastic connect info: disable health check")
	}
	if config.Gzip {
		logs.Logger.Info("elastic connect info: enable gzip")
		options = append(options, elastic.SetGzip(config.Gzip))
	} else {
		logs.Logger.Info("elastic connect info: disable gzip")
	}

	ElasticClient, err = elastic.NewClient(elastic.SetURL(config.URLs...))
	if err != nil {
		logs.Logger.Fatal("connect elastic filed", zap.Error(err), zap.String("version", elastic.Version), zap.Strings("urls", config.URLs))
	} else {
		logs.Logger.Info("connect elastic success", zap.String("version", elastic.Version), zap.Strings("urls", config.URLs))
		logs.Logger.Info("elastic server information", zap.String("information", ElasticClient.String()))
	}

}
