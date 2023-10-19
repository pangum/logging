package logging

import (
	"github.com/pangum/logging/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependencies().Build().Provide(new(plugin.Creator).New)
}
