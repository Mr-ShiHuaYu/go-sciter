package panicutil

type TryPanicCallback func(interface{})

func TryPanic(method TryPanicCallback) {
	if err := recover(); err != nil {
		method(err)
	}
}
