package logging

import (
	"github.com/goexl/simaqian"
	"github.com/pangum/pangu"
)

func newLogger(config *pangu.Config) (logger Logger, err error) {
	pc := new(panguConfig)
	if err = config.Load(pc); nil != err {
		return
	}
	conf := pc.Logging

	builder := simaqian.New().Skip(conf.Skip).Level(simaqian.ParseLevel(conf.Level))
	if nil != conf.Stacktrace && conf.Stacktrace.Enable() {
		builder.Stacktrace(conf.Stacktrace.Skip, conf.Stacktrace.Stack)
	}

	switch conf.Type {
	case typeZap:
		logger, err = builder.Zap().Build()
	case typeLogrus:
		logger = builder.Logrus().Build()
	case typeBuiltin:
		logger = builder.Builtin().Build()
	}

	return
}
