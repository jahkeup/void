package void

// Pointer creates a pointer to the value.
func Pointer[T voidable](value T) *T {
	return &value
}

// PointerSlice creates a slice of pointers to some values.
func PointerSlice[T voidable](value []T) []*T {
	if len(value) == 0 {
		return nil
	}

	ps := make([]*T, len(value))
	for i := range ps {
		ps[i] = Pointer(value[i])
	}

	return ps
}

// Value resolves a pointer to its value, or its zero value if unset. Callers
// are recommended to check if the pointer they have is nil before calling this
// function. Although you might consider some "style" to be more readable in
// this way:
//
//     var maybeString *string
//     maybeString, err := someOperation()
//     if err != nil {
//         return err
//     }
//
//     if void.Value(maybeString) != "" {
//         /* handle not-empty */
//     }
//
func Value[T voidable](p *T) T {
	if p != nil {
		return *p
	}

	// no value, make one.
	zero := new(T)

	// TODO: thoughts folks? runtime reflection here? ( T.NewVoid() ), eg: github.com/jahkeup/foo.Foo.NewVoid()?

	return *zero
}

// SliceValues resolves a slice of pointers to the respective value, or its zero value if
// unset (NOTE: this means nils will become an allocated zero value object in
// some cases!). Callers are recommended to check if the slice and element
// pointers they have are nil before calling this function.
func SliceValues[T voidable](value []*T) []T {
	if len(value) == 0 {
		return nil
	}

	fs := []T{}
	for i := range value {
		fs = append(fs, Value(value[i]))
	}

	return fs
}

// SliceValues resolves a slice of pointers to the respective value, or if nil,
// omitted. Would-be callers are recommended to consume the slice themselves if
// they want to handle nils instead - see SliceValues if you want nils to become
// their zero values.
func SliceValuesCompact[T voidable](value []*T) []T {
	if len(value) == 0 {
		return nil
	}

	fs := []T{}
	for i := range value {
		fs = append(fs, Value(value[i]))
	}

	return fs
}
