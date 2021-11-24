package logging

import (
	`sync`

	`github.com/pangum/pangu`
	`github.com/storezhang/simaqian`
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

	logger.Logger, err = simaqian.New(
		simaqian.Skip(logging.Skip),
		simaqian.Levels(logging.Level),
		simaqian.Types(logging.Type),
	)

	return
}
