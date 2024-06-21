package mo

import (
	"fmt"
	"time"
)

var mockTimeNowYear = func() int {
	return time.Now().Year()
}

func ExampleIO() {
	originalTimeNowYear := mockTimeNowYear
	defer func() { mockTimeNowYear = originalTimeNowYear }() // Restore the original function after the test

	mockTimeNowYear = func() int {
		return 2023
	}

	io := NewIO(func() int {
		return mockTimeNowYear()
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
