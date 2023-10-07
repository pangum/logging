package plugin

import (
	"github.com/pangum/logging"
)

type Wrapper struct {
	// 日志配置
	Logging logging.config `json:"logging" yaml:"logging" xml:"logging" toml:"logging" validate:"required"`
}
