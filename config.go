package log

import (
	`github.com/storezhang/glog`
)

type config struct {
	// 日志类型
	Type glog.Type `default:"zap" json:"type" yaml:"type" validate:"required,oneof=zap system logrus zero"`
	// 日志调用方法过滤层级
	Skip int `default:"1" json:"skip" yaml:"skip"`
}
