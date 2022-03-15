package logging

import (
	`sync`

	`github.com/goexl/simaqian`
	`github.com/pangum/pangu`
)

// Logger 日志简单包装，方便调用
type Logger struct {
	simaqian.Logger

	// 限制只能使用指针
	_ sync.Mutex
}

func newLogger(config *pangu.Config) (logger *Logger, err error) {
	logger = new(Logger)
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	logging := _panguConfig.Logging

	options := simaqian.NewOptions(
		simaqian.Skip(logging.Skip),
		simaqian.Levels(logging.Level),
		simaqian.Types(logging.Type),
	)
	if nil != logging.Stacktrace && logging.Stacktrace.Enable() {
		options = append(options, simaqian.Stacktrace(logging.Stacktrace.Skip, logging.Stacktrace.Stack))
	}
	logger.Logger, err = simaqian.New(options...)

	return
}
