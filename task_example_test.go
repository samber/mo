package mo

import (
	"fmt"
)

func ExampleTask() {
	originalTimeNowYear := mockTimeNowYear
	defer func() { mockTimeNowYear = originalTimeNowYear }() // Restore the original function after the test

	mockTimeNowYear = func() int {
		return 2023
	}

	task := NewTask(func() *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			resolve(mockTimeNowYear())
		})
	})

	// returns a future
	future := task.Run()

	// a Task never fail
	result, _ := future.Collect()

	fmt.Println(result)
	// Output:
	// 2023
}
