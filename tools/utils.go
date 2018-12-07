// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package tools

import (
	"crypto/sha256"
	"fmt"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"

	"github.com/theckman/go-flock"
	"math/rand"
	"os"
)

func LockOrDie(dir string) *flock.Flock {
	f := flock.New(dir)
	success, err := f.TryLock()
	if err != nil {
		logs.Logger.Error("Locking base-framework", zap.Error(err))
	}

	if !success {
		logs.Logger.Error("base-framework is locked", zap.Error(err))
	}

	return f
}

func ShuffleSlice(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// MakeDirectory makes directory if is not exists
func MakeDirectory(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dir, 0775)
		}
		return err
	}
	return nil
}

func StringInSlice(slice []string, v string) bool {
	for _, item := range slice {
		if v == item {
			return true
		}
	}
	return false
}
func SHA256(str string) (result string) {
	h := sha256.New()
	h.Write([]byte(str))
	result = fmt.Sprintf("%x", h.Sum(nil))
	return
}
