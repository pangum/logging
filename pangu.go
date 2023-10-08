package logging

import (
	"github.com/pangum/logging/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(new(plugin.Creator).New)
}
