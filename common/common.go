// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"bytes"
	"fmt"
)

// StringsJoin string array join
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}
func Join2String(split string, strs ...interface{}) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		var str interface{}
		switch str.(type) {
		case string:
			b.WriteString(str.(string))
		case int:
			b.WriteString(fmt.Sprintf("%d", str.(int)))
		case int64:
			b.WriteString(fmt.Sprintf("%d", str.(int)))

		}
	}
	str = b.String()
	return str

}
