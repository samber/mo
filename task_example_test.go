package mo

import (
	"fmt"
	"time"
)

func ExampleTask() {
	t := time.Date(2024, 6, 22, 0, 0, 0, 0, time.Local)

	task := NewTask(func() *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			resolve(t.Year())
		})
	})

	// returns a future
	future := task.Run()

	// a Task never fail
	result, _ := future.Collect()

	fmt.Println(result)
	// Output:
	// 2024
}
