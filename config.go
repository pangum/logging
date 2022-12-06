package logging

type config struct {
	// 日志级别
	Level string `default:"debug" json:"level" yaml:"level" xml:"level" toml:"level" validate:"oneof=debug info warn error fatal"`
	// 日志类型
	Type string `default:"zap" json:"type" yaml:"type" toml:"type" validate:"required,oneof=zap builtin logrus zero"`
	// 日志调用方法过滤层级
	Skip int `default:"2" json:"skip" yaml:"skip" toml:"skip"`
	// 调用堆栈层级
	Stacktrace *stacktrace `json:"stacktrace" yaml:"stacktrace" toml:"stacktrace"`
}
