package log

import (
	`github.com/storezhang/glog`
	`github.com/storezhang/pangu`
)

func newLogger(config *pangu.Config) (logger glog.Logger, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}
	log := panguConfig.Log

	opts := glog.NewOptions(glog.Skip(log.Skip))
	switch log.Type {
	case glog.TypeSystem:
		opts = append(opts, glog.System())
	case glog.TypeZap:
		opts = append(opts, glog.Zap())
	case glog.TypeLogrus:
		opts = append(opts, glog.Logrus())
	}
	logger, err = glog.New(opts...)

	return
}
