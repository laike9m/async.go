package async

import "reflect"

// Each iterates over all elements.
// http://caolan.github.io/async/docs.html#each
func Each(coll interface{}, iteratee interface{},
	callback ...func(e error)) {
	fn := reflect.ValueOf(iteratee)
	fnType := fn.Type()
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 {
		panic("Expected a unary function returning a single value")
	}
	if reflect.ValueOf(coll).Kind() != reflect.Slice {
		panic("Should pass slice as first argument")
	}
	iterable := reflect.ValueOf(coll)
	for i := 0; i < iterable.Len(); i++ {
		e := iterable.Index(i)
		fn.Call([]reflect.Value{e})
	}
	if len(callback) > 1 {
		panic("Expect one callback function!")
	}
	// TODO: catch above panic & other errors.
	if len(callback) == 1 {
		callback[0](nil)
	}
}
