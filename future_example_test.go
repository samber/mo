package mo

import "fmt"

func ExampleNewFuture_resolve() {
	value, err := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Collect()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	// foobar
	// <nil>
}

func ExampleNewFuture_reject() {
	value, err := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Collect()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	//
	// failure
}

func ExampleFuture_Collect_resolve() {
	value, err := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Collect()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	// foobar
	// <nil>
}

func ExampleFuture_Collect_reject() {
	value, err := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Collect()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	//
	// failure
}

func ExampleFuture_Result_resolve() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// foobar
	// <nil>
}

func ExampleFuture_Result_reject() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	//
	// failure
}

func ExampleFuture_Then_resolve() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Then(func(s string) (string, error) {
		return "baz", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// baz
	// <nil>
}

func ExampleFuture_Then_reject() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Then(func(s string) (string, error) {
		return "foobar", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	//
	// failure
}

func ExampleFuture_Catch_resolve() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Catch(func(err error) (string, error) {
		return "baz", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// foobar
	// <nil>
}

func ExampleFuture_Catch_reject() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Catch(func(err error) (string, error) {
		return "foobar", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// foobar
	// <nil>
}

func ExampleFuture_Finally_resolve() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Finally(func(value string, err error) (string, error) {
		return "baz", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// baz
	// <nil>
}

func ExampleFuture_Finally_reject() {
	result := NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Finally(func(value string, err error) (string, error) {
		return "foobar", nil
	}).Result()

	fmt.Println(result.OrEmpty())
	fmt.Println(result.Error())
	// Output:
	// foobar
	// <nil>
}

func ExampleFuture_Cancel_resolve() {
	NewFuture(func(resolve func(string), reject func(error)) {
		resolve("foobar")
	}).Cancel()
}

func ExampleFuture_Cancel_reject() {
	NewFuture(func(resolve func(string), reject func(error)) {
		reject(fmt.Errorf("failure"))
	}).Cancel()
}
