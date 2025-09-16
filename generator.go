package ginroute

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

func generateRouterFile(cfg *GenCfg, structs []*StructInfo) {
	t := template.Must(template.ParseFS(templateFS, "templates/route.tmpl"))
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
	genPath := cfg.Path + "/register_gen.go"
	err := os.WriteFile(genPath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("路由注册文件已生成: " + genPath)
}
