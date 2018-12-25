// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"github.com/itcloudy/base-framework/pkg/generator"
	"github.com/itcloudy/base-framework/tools"
	"os"
	"path"
	"reflect"
	"strings"
	"text/template"
)

var (
	projectPath string
	toTar       bool
)

func init() {
	flag.StringVar(&projectPath, "p", "", "project path")
	flag.BoolVar(&toTar, "t", false, "save to project location")
}
func main() {
	flag.Parse()
	if len(projectPath) == 0 {
		fmt.Println("project path is empty")
		os.Exit(-1)
	}
	generatorControllers()
}
func generatorControllers() {
	templateBytes, err := generator.GetTemplate("controller")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	//判断文件存放位置
	p, _ := os.Getwd()
	var outPut string
	if !toTar {
		outPut = path.Join(p, "out")
		 tools.MakeDirectory(outPut)
		outPut = path.Join(outPut, "controller")
		 tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "transport", "restful", "controllers")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		controllerName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("controller").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator controller %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(controllerName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", controllerName)

		f, _ := os.Create(controllerName)
		defer f.Close()
		defer tools.FormatSourceCode(controllerName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}

}
