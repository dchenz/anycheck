package anycheck

type optionFunc func(*config) error

func (f optionFunc) apply(cfg *config) error {
	return f(cfg)
}

type config struct {
	allowInterface bool
	allowAny       bool
}

type Option interface {
	apply(*config) error
}

func SetAllowInterface(v bool) Option {
	return optionFunc(func(cfg *config) error {
		cfg.allowInterface = v
		return nil
	})
}

func SetAllowAny(v bool) Option {
	return optionFunc(func(cfg *config) error {
		cfg.allowAny = v
		return nil
	})
}
