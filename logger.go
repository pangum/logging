package logging

import (
	`github.com/pangum/pangu`
	`github.com/storezhang/glog`
	`github.com/storezhang/simaqian`
)

// Logger 日志简单包装
type Logger struct {
	simaqian.Logger
}

func newLogger(config *pangu.Config) (logger *Logger, err error) {
	logger = new(Logger)
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	log := _panguConfig.Logging

	opts := glog.NewOptions(glog.Skip(log.Skip))
	switch log.Type {
	case glog.TypeSystem:
		opts = append(opts, glog.System())
	case glog.TypeZap:
		opts = append(opts, glog.Zap())
	case glog.TypeLogrus:
		opts = append(opts, glog.Logrus())
	}
	logger.Logger, err = glog.New(opts...)

	return
}
