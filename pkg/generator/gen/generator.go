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
	generatorRepository()
	generatorController()
	generatorService()
	generatorMockRepository()
	generatorMockService()
	generatorRepositoryInterface()
	generatorServiceInterface()

}
func generatorServiceInterface() {
	templateBytes, err := generator.GetTemplate("service_interface")
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
		outPut = path.Join(outPut, "interfaces")
		tools.MakeDirectory(outPut)
		outPut = path.Join(outPut, "services")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "interfaces", "services")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("service_interface").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator service interface %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}
}
func generatorRepositoryInterface() {
	templateBytes, err := generator.GetTemplate("repository_interface")
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
		outPut = path.Join(outPut, "interfaces")
		tools.MakeDirectory(outPut)
		outPut = path.Join(outPut, "repositories")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "interfaces", "repositories")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("service_interface").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator repository interface %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}
}
func generatorMockService() {
	templateBytes, err := generator.GetTemplate("mock_service")
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
		outPut = path.Join(outPut, "mocks")
		tools.MakeDirectory(outPut)
		outPut = path.Join(outPut, "services")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "mocks", "services")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("mock_service").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator mock service %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}
}
func generatorMockRepository() {
	templateBytes, err := generator.GetTemplate("mock_repository")
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
		outPut = path.Join(outPut, "mocks")
		tools.MakeDirectory(outPut)
		outPut = path.Join(outPut, "repositories")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "mocks", "repositories")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("mock_repository").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator mock repository %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}
}
func generatorController() {
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
		outPut = path.Join(outPut, "controllers")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "transport", "restful", "controllers")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("controller").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator controller %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}

}

func generatorService() {
	templateBytes, err := generator.GetTemplate("service_impl")
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
		outPut = path.Join(outPut, "services")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "services")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("service").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator service %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}

}

func generatorRepository() {
	templateBytes, err := generator.GetTemplate("repository_impl")
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
		outPut = path.Join(outPut, "repositories")
		tools.MakeDirectory(outPut)
	} else {
		outPut = path.Join(p, "..", "repositories", "common")
	}

	for _, model := range generator.AllModels {
		modelName := reflect.ValueOf(model).Type().String()
		splitRe := strings.Split(modelName, ".")
		modelName = splitRe[len(splitRe)-1]
		snakeModelName := tools.SnakeString(modelName)
		fileName := path.Join(outPut, tools.StringsJoin(snakeModelName, ".go"))
		tmpl, err := template.New("repository").Parse(string(templateBytes))
		if err != nil {
			fmt.Printf("generator repository %s failed, error info: %s ", modelName, err.Error())
			continue
		}
		if _, err := os.Stat(fileName); err == nil {
			if toTar {
				continue
			}
		}
		fmt.Println("generate file to : ", fileName)

		f, _ := os.Create(fileName)
		defer f.Close()
		defer tools.FormatSourceCode(fileName)
		Params := make(map[string]interface{})
		Params["ProjectPath"] = projectPath
		Params["ModelName"] = modelName
		tmpl.Execute(f, Params)

	}

}
