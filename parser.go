package ginroute

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"regexp"
	"strings"
)

var (
	regController = regexp.MustCompile(`@Controller\("([^"]+)"\)`)
	regRoute      = regexp.MustCompile(`@(\w+)(?:\(["'](.*?)["']\))?`)
)

func processFile(fset *token.FileSet, filename string, structsMap map[string]*StructInfo) {
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Printf("解析文件 %s 出错: %v", filename, err)
		return
	}

	// 第一次遍历：收集所有结构体
	ast.Inspect(file, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					// 检查 TypeSpec 的类型是否是 StructType
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						structInfo := &StructInfo{
							GroupName: typeSpec.Name.Name,
						}
						isCtrl := false
						// 注释是关联在 GenDecl 节点上的
						if genDecl.Doc != nil {
							for _, comment := range genDecl.Doc.List {
								matches := regController.FindStringSubmatch(comment.Text)
								if len(matches) > 1 {
									isCtrl = true
									structInfo.GroupPath = matches[1]
								}
							}
						}
						if isCtrl {
							structsMap[typeSpec.Name.Name] = structInfo
						}
					}
				}
			}
		}
		return true
	})

	// 第二次遍历：收集结构体字段和方法
	ast.Inspect(file, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil { // 是方法而不是函数
				recvType := exprToString(t.Recv.List[0].Type)
				if structInfo, exists := structsMap[strings.TrimPrefix(recvType, "*")]; exists {
					methodInfo := MethodInfo{
						Name: t.Name.Name,
					}
					isRoute := false
					// 获取方法注释
					if t.Doc != nil {
						for _, comment := range t.Doc.List {
							matches := regRoute.FindStringSubmatch(comment.Text)
							if len(matches) > 1 {
								isRoute = true
								methodInfo.Method = matches[1]
							}
							if len(matches) > 2 {
								methodInfo.Path = matches[2]
							}
						}
					}
					if isRoute {
						structInfo.Methods = append(structInfo.Methods, methodInfo)
					}
				}
			}
		}
		return true
	})

}
