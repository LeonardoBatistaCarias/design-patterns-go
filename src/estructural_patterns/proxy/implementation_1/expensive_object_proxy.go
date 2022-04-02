package main

type ExpensiveObjectProxy struct {
	object ExpensiveObject
}

func (e *ExpensiveObjectProxy) Process() {
	if e.object == nil {
		e.object = NewExpensiveObjectImpl()
	}
	e.object.Process()
}

func NewExpensiveObjectProxy() ExpensiveObject {
	return &ExpensiveObjectProxy{}
}
