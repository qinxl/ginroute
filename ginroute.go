package ginroute

import (
	"github.com/qinxl/ginroute/internal/astparser"
	"go/token"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func Generate(cfg *astparser.GenCfg) {
	if cfg == nil {
		cfg = &astparser.GenCfg{
			Path: "routes",
		}
	}
	fset := token.NewFileSet()
	structsMap := make(map[string]*astparser.StructInfo)
	err := filepath.Walk(cfg.Path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".go" && !strings.HasSuffix(info.Name(), "_gen.go") {
			astparser.ProcessFile(fset, path, structsMap)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	slice := make([]*astparser.StructInfo, 0, len(structsMap))
	for _, v := range structsMap {
		slice = append(slice, v)
	}
	astparser.GenerateRouterFile(cfg, slice)
}
