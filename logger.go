package logging

import (
	"github.com/goexl/gox"
	"github.com/goexl/simaqian"
	"github.com/pangum/pangu"
)

// Logger 日志简单包装，方便调用
type Logger struct {
	simaqian.Logger

	// 限制只能使用指针
	_ gox.CannotCopy
}

func newLogger(config *pangu.Config) (logger *Logger, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	_config := _panguConfig.Logging

	builder := simaqian.New().Skip(_config.Skip).Level(simaqian.ParseLevel(_config.Level))
	if nil != _config.Stacktrace && _config.Stacktrace.Enable() {
		builder.Stacktrace(_config.Stacktrace.Skip, _config.Stacktrace.Stack)
	}

	logger = new(Logger)
	switch _config.Type {
	case "zap":
		logger.Logger, err = builder.Zap().Build()
	case "logrus":
		logger.Logger = builder.Logrus().Build()
	case "builtin":
		logger.Logger = builder.Builtin().Build()
	}

	return
}
