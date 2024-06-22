package mo

import (
	"fmt"
	"time"
)

func ExampleIO() {
	t := time.Date(2024, 6, 22, 0, 0, 0, 0, time.Local)

	io := NewIO(func() int {
		return t.Year()
	})

	result1 := io.Run()
	result2 := io.Run()
	result3 := io.Run()

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	// Output:
	// 2024
	// 2024
	// 2024
}
