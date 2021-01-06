package main

import (
	"fmt"
	"sync"
)

const EnumOptionNil = "Not Implemented"

type factory struct{}

var (
	f    *factory
	once sync.Once

	listOptions []IFactoryOption
)

func setFactory() {
	once.Do(func() {
		f = &factory{}
	})
}

func addOption(option IFactoryOption) {
	listOptions = append(listOptions, option)
}

func GetFactory() IFactory {
	setFactory()
	return f
}

func (f *factory) getOption(params ...interface{}) (IFactoryOption, error) {
	for _, option := range listOptions {
		if option.validate(params...) {
			return option.getOption(), nil
		}
	}
	return nil, fmt.Errorf("Error Factory. Option %s not found", params)
}
