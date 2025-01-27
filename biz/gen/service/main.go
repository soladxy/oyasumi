package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

func main() {
	var err error
	inputName := flag.String("name", "example", "service name")
	inputOutputDir := flag.String("outdir", "./biz/service", "output dir, default: ./biz/service")
	flag.Parse()

	input := &inputData{
		name:      *inputName,
		outputDir: *inputOutputDir,
	}

	err = input.CheckInput()
	if err != nil {
		errAndOsExit(err)
		return
	}

	err = gen(input)
	if err != nil {
		errAndOsExit(err)
		return
	}
}

var genFiles = []genFile{
	{"interface", interfaceTmpl, "%s.go"},
	{"impl", implTmpl, "%s.impl.go"},
}

func gen(input *inputData) error {
	d := input.ToData()
	if _, err := os.Stat(d.OutputDir); os.IsNotExist(err) {
		// 如果文件夹不存在，则创建文件夹
		err := os.MkdirAll(d.OutputDir, 0755) // 使用 0755 创建多级文件夹
		if err != nil {
			return fmt.Errorf("创建文件夹失败: %w", err)
		}
		log.Printf("create output: %s", d.OutputDir)
	}

	for _, gf := range genFiles {
		parse, err := template.New(gf.name).Parse(gf.tmpl)
		if err != nil {
			return fmt.Errorf("create %s parse err: %w", gf.name, err)
		}
		sb := &strings.Builder{}
		err = parse.Execute(sb, d)
		if err != nil {
			return fmt.Errorf("parse %s err: %w", gf.name, err)
		}

		filepath := path.Join(d.OutputDir, fmt.Sprintf(gf.fileNameFmt, d.LowerName))
		// 检查文件夹是否存在

		f, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("%s os.Create err: %w", gf.name, err)
		}

		_, err = f.WriteString(sb.String())
		if err != nil {
			f.Close()
			return fmt.Errorf("%s WriteString err: %w", gf.name, err)
		}

		f.Close()
		log.Printf("create file: %s", filepath)
	}
	return nil
}

func errAndOsExit(err error) {
	log.Printf("[Error]: %v\n", err)
	os.Exit(1)
}
