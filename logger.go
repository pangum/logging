package logging

import (
	`github.com/pangum/pangu`
	`github.com/storezhang/simaqian`
)

// Logger 日志简单包装，方便调用
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

	opts := simaqian.NewOptions(simaqian.Skip(log.Skip))
	switch log.Type {
	case simaqian.TypeSystem:
		opts = append(opts, simaqian.System())
	case simaqian.TypeZap:
		opts = append(opts, simaqian.Zap())
	case simaqian.TypeLogrus:
		opts = append(opts, simaqian.Logrus())
	}
	logger.Logger, err = simaqian.New(opts...)

	return
}
