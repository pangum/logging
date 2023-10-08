package plugin

import (
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
		self := config.Loki
		loki := builder.Loki().Url(self.Url)
		if "" != self.Username || "" != self.Password {
			loki.Username(self.Username)
			loki.Password(self.Password)
		}
		if nil != self.Batch {
			loki.Batch(self.Batch.Size, self.Batch.Wait)
		}
		if 0 != len(self.Labels) {
			loki.Labels(self.Labels)
		}
		logger, err = loki.Build()
	}

	return
}
