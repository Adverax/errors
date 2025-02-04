package errors

func Check(errs ...error) error {
	es := NewErrors()
	return es.Check(errs...)
}

func Must[T any](e T, err error) T {
	if err != nil {
		panic(err)
	}
	return e
}

type Option[T any] func(*Options[T])

type Options[T any] struct {
	items []*wantItem[T]
}

type wantItem[T any] struct {
	value T
	err   error
}

func WithItem[T any](value T, err error) Option[T] {
	return func(opts *Options[T]) {
		opts.items = append(opts.items, &wantItem[T]{value: value, err: err})
	}
}

func CheckList[T any](options ...Option[T]) ([]T, error) {
	opts := &Options[T]{}
	for _, opt := range options {
		opt(opts)
	}

	var errs *Errors
	res := make([]T, 0, len(opts.items))
	for _, item := range opts.items {
		if item.err != nil {
			if errs == nil {
				errs = NewErrors()
			}
			errs.Add(item.err)
		} else {
			res = append(res, item.value)
		}
	}

	if errs == nil {
		return res, nil
	}

	return nil, errs.ResError()
}
