package main

import (
	"github.com/pingdai/gauto_conf/confx"
	"github.com/pingdai/gauto_conf/ginx"
)

var Config Cfg

func init() {
	confx.ConfP(&Config)
}

type Cfg struct {
	ConnTimeout int        `json:"conn_timeout"`
	Ginx        *ginx.Ginx `json:"ginx"`
	// other config
}

func main() {
	// todo something
	Config.Ginx.Run()
}
