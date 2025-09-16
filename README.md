# ginroute
自动生成gin路由注册文件

```go
// main.go
package main

import "github.com/qinxl/ginroute"

func main() {
	cfg := &ginroute.GenCfg{
		Path: "internal/routes", // 默认routes
	}
	ginroute.Generate(cfg)
}
```

```go
// internal/routes/test.go
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SysUserCtrl @Controller("/api/user")
type SysUserCtrl struct {
}

// Create @POST("/")
func (u *SysUserCtrl) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// List @GET("/")
func (u *SysUserCtrl) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hi",
	})
}

// SysRole @Controller("/role", group="/admin")
type SysRole struct {
}

// Create @POST("/")
func (u *SysRole) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// List @GET("/")
func (u *SysRole) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

```


