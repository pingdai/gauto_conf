package confx

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
)

// 解析结构体
func ConfP(in interface{}) {

	tpe := reflect.TypeOf(in)
	if tpe.Kind() != reflect.Ptr {
		panic(fmt.Errorf("ConfP pass ptr for setting value"))
	}

	bts := getConfigContent()
	fmt.Printf("获取配置文件信息:%s\n", string(bts))
	if err := json.Unmarshal(bts, in); err != nil {
		panic(fmt.Sprintf("json.Unmarshal conf err:%v", err))
	}

	UnmarshalConf(in)
}

func getConfigContent() []byte {
	var localCfgPath string
	var bts []byte
	var err error

	flag.StringVar(&localCfgPath, "c", "", "local config file path")
	flag.Parse()

	if localCfgPath != "" {
		// 进行本地文件解析
		bts, err = ioutil.ReadFile(localCfgPath)
		if err != nil {
			panic(fmt.Sprintf("Read local file[%s] err:%v", localCfgPath, err))
		}
	} else {
		// todo 从远端获取配置服务器获取相应配置，比如zk、etcd等
	}

	return bts
}
