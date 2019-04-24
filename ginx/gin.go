package ginx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Ginx struct {
	Engine     *gin.Engine
	ListenPort int `json:"listen_port"`
	init       bool
}

func (ginx *Ginx) Init() {
	if !ginx.init {
		ginx.New()
		ginx.init = true
	}
}

func (ginx *Ginx) New() {
	ginx.Engine = gin.New()
	// 添加一些中间件
	// ginx.Engine.Use(...)
	// TODO
}

// addr = [:port] e.g. :1500
func (ginx *Ginx) Run() error {
	if !ginx.init {
		panic("gin engine not init.")
	}
	if ginx.ListenPort == 0 {
		return errors.New("server listen port error")
	}

	return ginx.Engine.Run(fmt.Sprintf(":%d", ginx.ListenPort))
}
