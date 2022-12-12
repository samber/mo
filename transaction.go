package mo

type transactionStep[T any] struct {
	exec       func(T) (T, error)
	onRollback func(T) T
}

// NewTransaction instanciate a new transaction.
func NewTransaction[T any]() *Transaction[T] {
	return &Transaction[T]{
		steps: []transactionStep[T]{},
	}
}

// Transaction implements a Saga pattern
type Transaction[T any] struct {
	steps []transactionStep[T]
}

// Then adds a step to the chain of callbacks. It returns the same Transaction.
func (t *Transaction[T]) Then(exec func(T) (T, error), onRollback func(T) T) *Transaction[T] {
	t.steps = append(t.steps, transactionStep[T]{
		exec:       exec,
		onRollback: onRollback,
	})

	return t
}

// Process runs the Transaction steps and rollbacks in case of errors.
func (t *Transaction[T]) Process(state T) (T, error) {
	var i int
	var err error

	for i < len(t.steps) {
		state, err = t.steps[i].exec(state)
		if err != nil {
			break
		}

		i++
	}

	if err == nil {
		return state, nil
	}

	for i > 0 {
		i--
		state = t.steps[i].onRollback(state)
	}

	return state, err
}
