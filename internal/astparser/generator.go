package astparser

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateRouterFile(cfg *GenCfg, structs []*StructInfo) {
	t := template.Must(template.ParseFiles("template/route.tmpl"))
	var params struct {
		Pkg     string
		Structs []*StructInfo
	}
	params.Structs = structs
	params.Pkg = filepath.Base(cfg.Path)

	var buf bytes.Buffer
	if err := t.Execute(&buf, params); err != nil {
		log.Fatal(err)
	}

	// 写入文件
	err := os.WriteFile("routes/register_gen.go", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("路由注册文件已生成: routes/register_gen.go")
}
