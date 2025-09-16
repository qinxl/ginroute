package astparser

type StructInfo struct {
	GroupName string
	GroupPath string
	Methods   []MethodInfo
}

type MethodInfo struct {
	Name   string
	Method string
	Path   string
}

type GenCfg struct {
	Path string
}
