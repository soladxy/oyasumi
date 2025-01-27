package main

import (
	"fmt"
	"github.com/soladxy/oyasumi/biz/util"
	"path"
	"strings"
)

type inputData struct {
	name      string
	outputDir string
}

func (i *inputData) CheckInput() error {
	if !util.IsServiceNameValid(i.name) {
		return fmt.Errorf("invalid service name: %s", i.name)
	}
	return nil
}

func (i *inputData) ToData() *data {
	name := strings.Trim(strings.ToLower(i.name), " \t\n")
	return &data{
		FirstCharUpperName: util.CapitalizeFirstLetter(name),
		LowerName:          name,
		OutputDir:          path.Join(i.outputDir, name),
	}
}

type data struct {
	FirstCharUpperName string // 代码内的名称，首字母大写
	LowerName          string // 文件系统名称，全小写
	OutputDir          string // 生成目录，文件生成的具体目录
}

type genFile struct {
	name        string // file name
	tmpl        string // 模板
	fileNameFmt string // 文件名模板
}
