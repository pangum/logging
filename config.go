package logging

type config struct {
	// 日志级别
	Level string `default:"debug" json:"level" yaml:"level" xml:"level" toml:"level"`
	// 日志类型
	Type string `default:"zap" json:"type" yaml:"type" toml:"type" validate:"required,oneof=zap system logrus zero"`
	// 日志调用方法过滤层级
	Skip int `default:"1" json:"skip" yaml:"skip" toml:"skip"`
}
