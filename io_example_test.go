package mo

import (
	"fmt"
	"time"
)

func ExampleIO() {
	io := NewIO(func() int {
		return time.Now().Year()
	})

	result1 := io.Run()
	result2 := io.Run()
	result3 := io.Run()

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	// Output:
	// 2023
	// 2023
	// 2023
}
