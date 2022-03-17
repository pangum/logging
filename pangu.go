package logging

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Musts(newLogger)
}
