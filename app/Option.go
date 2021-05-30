package app

import (
	"github.com/Jack1c/toyms/log"
	"github.com/Jack1c/toyms/registry"
	"github.com/Jack1c/toyms/server"
)

type appOption struct {
	name     string
	version  string
	server   server.ServerInterface
	logger   log.Logger
	register *registry.Register
}

type AppOptions interface {
	apply(*appOption)
}

type funcAppOption struct {
	f func(options *appOption)
}

func (fdo *funcAppOption) apply(do *appOption) {
	fdo.f(do)
}

func newFuncAppOption(f func(*appOption)) *funcAppOption {
	return &funcAppOption{
		f: f,
	}
}

func Name(name string) AppOptions {
	return newFuncAppOption(func(options *appOption) {
		options.name = name
	})
}

func Versopn(v string) AppOptions {
	return newFuncAppOption(func(op *appOption) {
		op.version = v
	})
}

func Server(server server.ServerInterface) AppOptions {
	return newFuncAppOption(func(options *appOption) {
		options.server = server
	})
}

func Logger(logger log.Logger) AppOptions {
	return newFuncAppOption(func(option *appOption) {
		if logger != nil {
			option.logger = logger
		}
	})
}

func Register(register *registry.Register) AppOptions {
	return newFuncAppOption(func(option *appOption) {
		if register != nil {
			option.register = register
		}
	})
}
