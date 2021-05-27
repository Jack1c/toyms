package app

import "github.com/Jack1c/toyms/server"

type appOptions struct {
	name    string
	version string
	server  server.ServerInterface
}

type AppOptions interface {
	apply(*appOptions)
}

type funcAppOption struct {
	f func(options *appOptions)
}

func (fdo *funcAppOption) apply(do *appOptions) {
	fdo.f(do)
}

func newFuncAppOption(f func(*appOptions)) *funcAppOption {
	return &funcAppOption{
		f: f,
	}
}

func Name(name string) AppOptions {
	return newFuncAppOption(func(options *appOptions) {
		options.name = name
	})
}

func Versopn(v string) AppOptions {
	return newFuncAppOption(func(op *appOptions) {
		op.version = v
	})
}

func Server(server server.ServerInterface) AppOptions {
	return newFuncAppOption(func(options *appOptions) {
		options.server = server
	})
}
