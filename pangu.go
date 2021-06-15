package log

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newLogger); nil != err {
		panic(err)
	}
}
