package plugin

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/pangum/logging/internal/core"
	"github.com/pangum/pangu"
)

type Creator struct {
	// 纯方法封装
}

func (c *Creator) New(config *pangu.Config) (logger simaqian.Logger, err error) {
	wrapper := new(Wrapper)
	if le := config.Load(wrapper); nil != le {
		err = le
	} else {
		logger, err = c.new(&wrapper.Logging)
	}

	return
}

func (c *Creator) new(config *Config) (logger simaqian.Logger, err error) {
	builder := simaqian.New().Level(simaqian.ParseLevel(config.Level))
	if nil != config.Stacktrace {
		builder.Stacktrace(*config.Stacktrace)
	}

	switch config.Type {
	case core.TypeZap:
		logger, err = builder.Zap().Build()
	case core.TypeLogrus:
		logger = builder.Logrus().Build()
	case core.TypeBuiltin:
		logger = builder.Builtin().Build()
	default:
		err = exc.NewField("不支持的日志收集器", field.New("type", config.Type))
	}

	return
}
