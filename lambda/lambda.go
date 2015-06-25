package lambda

import "github.com/jcla1/gisp/core"

type List struct {
	Head core.Any
	Tail *List
}

func Cons(hd core.Any, _tl core.Any) core.Any {
	var tl *List
	if _tl != nil {
		tmp := _tl.(List)
		tl = &tmp
	}
	return List{hd, tl}
}

func Car(l core.Any) core.Any {
	return l.(List).Head
}

func Cdr(l core.Any) core.Any {
	return *(l.(List).Tail)
}

func IsNull(l core.Any) core.Any {
	return l == nil
}
