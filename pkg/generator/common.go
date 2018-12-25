// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func GetTemplate(sufix string) (templateBytes []byte, err error) {
	p, _ := os.Getwd()
	outPut := path.Join(p, "template")
	fmt.Println(outPut)
	fileName := fmt.Sprintf("%s/%s.tpl", outPut, sufix)
	templateBytes, err = ioutil.ReadFile(fileName)
	return
}
