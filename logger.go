package logging

import (
	`github.com/pangum/pangu`
	`github.com/storezhang/glog`
)

func newLogger(config *pangu.Config) (logger glog.Logger, err error) {
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
	logger, err = glog.New(opts...)

	return
}
