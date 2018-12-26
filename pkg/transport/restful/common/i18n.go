// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func parseTags(langs ...string) []language.Tag {
	tags := []language.Tag{}
	for _, lang := range langs {
		t, _, err := language.ParseAcceptLanguage(lang)
		if err != nil {
			continue
		}
		tags = append(tags, t...)
	}
	return tags
}
func messageI18n(c *gin.Context, code int) (message string) {
	//lang:= c.Request.FormValue("lang")
	//accept:=c.GetHeader("Accept-Language")
	// todo 可考虑在token中增加语言，从中取出语言，如果客户端更新语言，需要刷新token
	return
}
