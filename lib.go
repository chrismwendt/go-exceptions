package exceptions

import (
	"fmt"
	"runtime"
	"strings"
)

func Catch(err *error) {
	if v := recover(); v != nil {
		if e, ok := v.(error); ok {
			*err = e
		} else {
			*err = fmt.Errorf("%+v", v)
		}
	}
}

func Throw(e error, msgs ...string) {
	if e != nil {
		panic(wrap(e, msgs...))
	}
}

func Throw1[T any](t T, e error, msgs ...string) T {
	if e != nil {
		panic(wrap(e, msgs...))
	}
	return t
}

func Throw2[T1, T2 any](t1 T1, t2 T2, e error, msgs ...string) (T1, T2) {
	if e != nil {
		panic(wrap(e, msgs...))
	}
	return t1, t2
}

func wrap(e error, msgs ...string) error {
	for _, msg := range msgs {
		e = fmt.Errorf("%s: %w", msg, e)
	}

	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return e
	}

	details := runtime.FuncForPC(pc)
	if details == nil {
		return e
	}

	name := details.Name()
	name = strings.ReplaceAll(name, "[...]", "")
	components := strings.Split(name, ".")
	if len(components) < 1 {
		return e
	}
	e = fmt.Errorf("%s: %w", components[len(components)-1]+"()", e)
	return e
}
