// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package publisher

import (
	"github.com/centrifugal/gocent"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"

	"sync"
	"time"
)

type ClientsChannels struct {
	storage map[int64]string
	sync.RWMutex
}

func (cn *ClientsChannels) Set(id int64, s string) {
	cn.Lock()
	defer cn.Unlock()
	cn.storage[id] = s
}

func (cn *ClientsChannels) Get(id int64) string {
	cn.RLock()
	defer cn.RUnlock()
	return cn.storage[id]
}

var (
	clientsChannels   = ClientsChannels{storage: make(map[int64]string)}
	centrifugoTimeout = time.Second * 5
	publisher         *gocent.Client
	config            conf.CentrifugoConfig
)

// InitCentrifugo client
func InitCentrifugo(cfg conf.CentrifugoConfig) {
	logs.Logger.Error("Init Centrifugo client, connect information:" + cfg.String())
	config = cfg
	publisher = gocent.New(gocent.Config{Addr: cfg.URL, Key: cfg.Secret})
}

// Write is publishing data to server
/*func Write(userID int64, data string) (bool, error) {
	return publisher.Publish("client"+strconv.FormatInt(userID, 10), []byte(data))
}
*/
