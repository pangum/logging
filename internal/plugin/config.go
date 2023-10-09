package plugin

import (
	"github.com/pangum/logging/internal/config"
	"github.com/pangum/logging/internal/core"
)

type Config struct {
	// 日志级别
	// nolint: lll
	Level string `default:"debug" json:"level" yaml:"level" xml:"level" toml:"level" validate:"oneof=debug info warn error fatal"`
	// 类型
	// nolint: lll
	Type core.Type `default:"zap" json:"type" yaml:"type" xml:"type" toml:"type" validate:"required,oneof=zap builtin logrus zero loki"`
	// 日志调用方法过滤层级
	Skip int `default:"2" json:"skip" yaml:"skip" xml:"skip" toml:"skip"`
	// 调用堆栈层级
	Stacktrace *config.Stacktrace `json:"stacktrace" yaml:"stacktrace" xml:"stacktrace" toml:"stacktrace"`
	// Loki日志配置
	Loki *config.Loki `json:"loki" yaml:"loki" xml:"loki" toml:"loki" validate:"required_if=Type loki"`
}