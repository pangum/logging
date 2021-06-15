package log

type panguConfig struct {
	// 日志配置
	Log config `json:"log" yaml:"log" validate:"required"`
}
