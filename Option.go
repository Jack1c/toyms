package toyms

type appOptions struct {
	name    string
	version string
}

type OptionInterface interface {
	apply(*appOptions)
}
