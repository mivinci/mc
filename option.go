package mc

type Option struct {
	cache Cache
}

type OptionFunc func(*Option)

func WithCache(c Cache) OptionFunc {
	return func(o *Option) {
		o.cache = c
	}
}
