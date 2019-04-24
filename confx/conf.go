package confx

import (
	"reflect"
)

func UnmarshalConf(c interface{}) {
	rv := Indirect(reflect.ValueOf(c))

	if !rv.CanSet() || rv.Type().Kind() != reflect.Struct {
		panic("UnmarshalConf need an variable which can set")
	}

	InitialRoot(rv)
}

type ICanInit interface {
	Init()
}

func InitialRoot(rv reflect.Value) {
	tpe := rv.Type()
	for i := 0; i < tpe.NumField(); i++ {
		value := rv.Field(i)
		if conf, ok := value.Interface().(ICanInit); ok {
			conf.Init()
		}
	}
}

func Indirect(v reflect.Value) reflect.Value {
	for {
		if v.Kind() == reflect.Interface {
			e := v.Elem()
			if e.Kind() == reflect.Ptr {
				v = e
				continue
			}
		}
		if v.Kind() != reflect.Ptr {
			break
		}
		v = v.Elem()
	}
	return v
}
