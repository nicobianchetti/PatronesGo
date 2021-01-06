package main

type IFactoryOption interface {
	validate(...interface{}) bool
	getOption() IFactoryOption
	Do()
}

type IFactory interface {
	getOption(...interface{}) (IFactoryOption, error)
}
