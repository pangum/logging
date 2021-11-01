package logging

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newLogger); nil != err {
		panic(err)
	}
}
