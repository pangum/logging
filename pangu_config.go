package logging

type panguConfig struct {
	// 日志配置
	Logging config `json:"logging" yaml:"logging" xml:"logging" toml:"logging" validate:"required"`
}
