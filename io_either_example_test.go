package mo

import (
	"errors"
	"fmt"
	"os"
)

func ExampleIOEither1() {
	io := NewIOEither1(func(path string) (bool, error) {
		_, err := os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else if err != nil {
			// other errors
			return false, err
		}

		return true, nil
	})

	either1 := io.Run("./io_either.go")
	either2 := io.Run("./foo_bar.go")

	exist1, _ := either1.Right()
	exist2, _ := either2.Right()

	fmt.Println(exist1)
	fmt.Println(exist2)
	// Output:
	// true
	// false
}
