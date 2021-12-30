package logging

type stacktrace struct {
	// 是否开启
	Enabled *bool `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 跳过层级
	Skip int `json:"skip" yaml:"skip" xml:"skip" toml:"skip" validate:"min=0"`
	// 需要打印的堆栈数
	Stack int `json:"stack" yaml:"stack" xml:"stack" toml:"stack" validate:"min=0"`
}

func (s *stacktrace) Enable() bool {
	return nil == s.Enabled || *s.Enabled
}
