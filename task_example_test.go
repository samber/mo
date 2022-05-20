package mo

import (
	"fmt"
	"time"
)

func ExampleTask() {
	task := NewTask(func() *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			resolve(time.Now().Year())
		})
	})

	// returns a future
	future := task.Run()

	// a Task never fail
	result, _ := future.Collect()

	fmt.Println(result)
	// Output:
	// 2022
}
