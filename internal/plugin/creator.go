package plugin

import (
	"github.com/goexl/simaqian"
	"github.com/pangum/logging"
	"github.com/pangum/logging/internal/core"
	"github.com/pangum/pangu"
)

type Creator struct{}

func (c *Creator) New(config *pangu.Config) (logger logging.Logger, err error) {
	wrapper := new(Wrapper)
	if le := config.Load(wrapper); nil != le {
		err = le
	} else {
		logger, err = c.new(wrapper.Logging)
	}

	return
}

func (c *Creator) new(config *Config) (logger logging.Logger, err error) {
	builder := simaqian.New().Skip(config.Skip).Level(simaqian.ParseLevel(config.Level))
	if nil != config.Stacktrace && config.Stacktrace.Enable() {
		builder.Stacktrace(config.Stacktrace.Skip, config.Stacktrace.Stack)
	}

	switch config.Type {
	case core.TypeZap:
		logger, err = builder.Zap().Build()
	case core.TypeLogrus:
		logger = builder.Logrus().Build()
	case core.TypeBuiltin:
		logger = builder.Builtin().Build()
	case core.TypeLoki:
		logger = builder
	}

	return
}
