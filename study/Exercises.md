# Exercises

Practical coding exercises for every block in the Go Roadmap.

## Block 1: Language Basics and Syntax

### 1.1 Basic Syntax

1. Create a Go program with `package main` that prints "Hello, Go!" to stdout.
1. Create a program that imports `fmt` and `math` and prints the square root of 144.
1. Write a program with both single-line and multi-line comments explaining each line.
1. Demonstrate that Go inserts semicolons automatically: write two statements on one line separated by `;` and explain why putting `{` on a new line after `func main()` causes a compile error.
1. Create a package `greeting` with an exported function `Hello` and an unexported helper `buildMessage`. Import and use `Hello` from `main`.

### 1.2 Variables and Constants

1. Declare variables using `var` for `name string`, `age int`, and `active bool`. Print their zero values.
1. Use short declaration `:=` to create variables `x`, `y`, `z` with values 10, 20, 30 in one line.
1. Swap two variables `a` and `b` using multiple assignment without a temp variable.
1. Print the zero values of `int`, `float64`, `bool`, `string`, and a pointer `*int`.
1. Show that a variable declared inside an `if` block is not accessible outside it.
1. Declare a `const` block with `StatusOK = 200`, `StatusNotFound = 404`, `StatusError = 500`.
1. Use `iota` to create an enum for weekdays (Sunday=0 through Saturday=6).
1. Show the difference between a typed constant `const x int = 5` and an untyped constant `const y = 5` by assigning them to different numeric types.

### 1.3 Basic Data Types

1. Declare a `bool` variable, negate it, and print both values.
1. Declare variables of types `int8`, `int16`, `int32`, `int64` and print their max values using `math.MaxInt8`, etc.
1. Declare `uint8`, `uint16`, `uint32`, `uint64` variables and print their max values.
1. Show the difference between `int` (platform-dependent) and `int64` by printing their sizes with `unsafe.Sizeof`.
1. Assign a character `'A'` to a `byte` variable and print it as both a number and a character.
1. Assign a Unicode character (e.g. `'Я'`) to a `rune` variable. Print its numeric value and character representation.
1. Declare `float32` and `float64` variables. Show precision loss by computing `1.0/3.0` with each.
1. Create two `complex128` numbers and print their sum, product, real part (`real()`), and imaginary part (`imag()`).

### 1.4 Strings

1. Declare a string `s := "Hello, World!"` and print its length.
1. Show that strings are immutable: try `s[0] = 'h'` and observe the compile error.
1. Print the same string using an interpreted literal (`"line1\nline2"`) and a raw literal (backticks).
1. Print the byte at index 0 and the length of the string `"Hello"`.
1. Extract the substring `"World"` from `"Hello, World!"` using a slice expression.
1. Show that `len("Привет")` returns the byte count, not the rune count. Use `utf8.RuneCountInString` to get the rune count.
1. Iterate over `"Hello, Мир!"` with `for i, r := range s` and print each index and rune.

### 1.5 Operators

1. Write a function that takes two `int` values and prints the result of `+`, `-`, `*`, `/`, `%`.
1. Compare two strings `"abc"` and `"abd"` using all comparison operators and print results.
1. Write an expression using `&&`, `||`, `!` that checks if a number is between 10 and 100 (inclusive) or equals 0.
1. Use bitwise operators to: (a) set bit 3 of a number, (b) clear bit 2, (c) toggle bit 1, (d) check if bit 0 is set.
1. Use `<<=` and `>>=` to multiply and divide a number by 2.
1. Show that `i++` is a statement, not an expression (cannot be used in `fmt.Println(i++)`).
1. Declare a variable `x := 42`, take its address with `&x`, and dereference it with `*`.

### 1.6 Type Conversions

1. Convert an `int` to `float64` and back. Show that `int(3.9)` truncates to `3`.
1. Convert between `int32` and `int64` explicitly. Show that implicit conversion causes a compile error.
1. Convert a string to `[]byte` and `[]rune`. Print the lengths of each for `"Привет"`.
1. Convert an integer to a string using `strconv.Itoa` and parse it back with `strconv.Atoi`. Handle the error from `Atoi`.

## Block 2: Control Flow

### 2.1 Conditional Operators

1. Write a function `abs(x int) int` that returns the absolute value using `if`.
1. Write an `if` with initialization: `if v := compute(); v > 0 { ... }` that checks the return value of a function.
1. Write a function `classify(n int) string` that returns "negative", "zero", or "positive" using `if-else if-else`.
1. Write a function `max(a, b int) int` without using a ternary operator (Go doesn't have one).

### 2.2 Switch

1. Write a function `dayType(day string) string` that returns "weekday" or "weekend" using switch.
1. Write a switch with expressions in case clauses to classify a score: "A" (90-100), "B" (80-89), "C" (70-79), "F" (below 70).
1. Write a switch without an expression (tagless switch) as an alternative to an `if-else` chain for temperature classification: "cold" (<10), "mild" (10-25), "hot" (>25).
1. Use multiple values in a single case: classify a month number into a season.
1. Demonstrate `fallthrough`: write a switch where matching case 1 also executes case 2.
1. Write a function that accepts `interface{}` and uses a type switch to print the type and value.

### 2.3 Loops

1. Print numbers from 1 to 20 using a classic `for i := 1; i <= 20; i++`.
1. Read user input in a loop using `for` as `while`: keep reading until the user types "quit".
1. Write an infinite loop that breaks when a random number equals 7.
1. Iterate over an array `[5]int{10, 20, 30, 40, 50}` using `for range` and print index and value.
1. Iterate over a slice `[]string{"go", "rust", "python"}` and print each element.
1. Iterate over the string `"Hёllo"` with `range` and print each rune with its byte offset.
1. Iterate over a `map[string]int{"a": 1, "b": 2, "c": 3}` and print key-value pairs.
1. Create a channel, send 3 values, close it, and iterate with `for range`.
1. Use `break` to exit a loop early and `continue` to skip even numbers while printing 1-20.
1. Use a label to break out of a nested loop: find the first pair `(i, j)` where `i*j == 12` with `i, j` in `[1, 10]`.

### 2.4 Flow Control

1. Write a program that uses `goto` to skip a section of code (understand why this is rarely used).
1. Write a function with multiple `return` points based on input validation.
1. Write a function that opens a file, defers `file.Close()`, and reads the content.
1. Write a function that panics with a message and observe the stack trace.
1. Write a function that uses `recover` inside a deferred function to catch a panic and return an error instead.

## Block 3: Data Structures

### 3.1 Arrays

1. Declare an array `var a [5]int` and print all elements (zero values).
1. Initialize an array with values: `b := [3]string{"go", "rust", "c"}`.
1. Access and modify the element at index 2.
1. Print the length of an array using `len`.
1. Create a 3x3 matrix as `[3][3]int` and fill it with values 1-9.
1. Show that assigning one array to another copies all values (modify the copy, original unchanged).
1. Compare two `[3]int` arrays using `==`.
1. Iterate over an array using `for range` and compute the sum of elements.

### 3.2 Slices

1. Declare a nil slice `var s []int` and print its length, capacity, and whether it's nil.
1. Create a slice with `make([]int, 5, 10)` and print len and cap.
1. Create a slice from an array: `arr := [5]int{1,2,3,4,5}; s := arr[1:4]`. Modify `s[0]` and check `arr[1]`.
1. Show how `append` grows a slice: start with `make([]int, 0, 2)`, append 5 elements, and print cap after each append.
1. Copy one slice into another using `copy`. Show that the slices are independent.
1. Create a 2D slice (matrix) and fill it dynamically.
1. Compare a nil slice (`var s []int`) with an empty slice (`s := []int{}`): both have len 0, but nil check differs.
1. Write a function `remove(s []int, i int) []int` that removes element at index `i` without preserving order (swap with last).
1. Write a function `filter(s []int, fn func(int) bool) []int` that returns a new slice with only elements where `fn` returns true.

### 3.3 Map

1. Declare a nil map and show that writing to it panics. Then create one with `make`.
1. Create a map `ages := map[string]int{"Alice": 30, "Bob": 25}` and add `"Charlie": 35`.
1. Check if a key exists using the two-value form: `v, ok := m[key]`.
1. Delete a key with `delete(m, key)` and verify it's gone.
1. Count word frequencies in a sentence using `map[string]int`.
1. Iterate over a map and print keys sorted alphabetically (collect keys into a slice, sort, then iterate).
1. Show that iterating over a map twice may yield different orders.
1. Demonstrate that writing to a map from multiple goroutines without synchronization causes a panic (use `go run -race`).

### 3.4 Structs

1. Define a struct `Person` with fields `Name string`, `Age int`. Create an instance and print it.
1. Initialize a struct using field names: `p := Person{Name: "Alice", Age: 30}`.
1. Create an anonymous struct inline: `p := struct{ X, Y int }{10, 20}`.
1. Create a pointer to a struct with `&Person{...}` and access fields through the pointer.
1. Define `Address` and `Person` structs where `Person` has an `Address` field (nested).
1. Use embedding: define `Employee` that embeds `Person` and has an additional `Salary` field.
1. Add a JSON tag to struct fields: `json:"name,omitempty"` and marshal to JSON.
1. Show that unexported fields (lowercase) are not accessible outside the package.
1. Compare two structs of the same type using `==`.

## Block 4: Functions

### 4.1 Function Basics

1. Write a function `add(a, b int) int` that returns the sum.
1. Write a function `divide(a, b float64) (float64, error)` that returns an error if `b == 0`.
1. Write a function with named return values: `func split(sum int) (x, y int)`.
1. Use `_` to ignore one of the return values from a multi-return function.
1. Write a recursive function `factorial(n int) int`.
1. Write a recursive function `fibonacci(n int) int` and observe the performance for large `n`.

### 4.2 Variadic Functions

1. Write a function `sum(nums ...int) int` that sums all arguments.
1. Call a variadic function by unpacking a slice: `nums := []int{1,2,3}; sum(nums...)`.

### 4.3 Higher-Order Functions

1. Assign a function to a variable: `add := func(a, b int) int { return a + b }`.
1. Write a function `apply(f func(int, int) int, a, b int) int` that calls `f` with `a` and `b`.
1. Write a function `makeMultiplier(factor int) func(int) int` that returns a closure.
1. Write a function `makeCounter() func() int` that returns a closure incrementing an internal counter.
1. Use a closure to create an accumulator that sums values passed to it across calls.

### 4.4 Methods

1. Define a type `Rect struct { Width, Height float64 }` with a method `Area() float64`.
1. Add a method `Scale(factor float64)` with a pointer receiver that modifies the rect.
1. Show the difference: call a value receiver method and a pointer receiver method, checking if the original is modified.
1. Define a custom type `type Celsius float64` and add a method `ToFahrenheit() float64`.

### 4.5 Defer, Panic, Recover

1. Write a function with three deferred calls and print the order of execution (LIFO).
1. Write a function with named return values where a deferred function modifies the return value.
1. Write a function that opens a file, defers `Close()`, writes data, and returns.
1. Write a function `mustParseInt(s string) int` that panics if parsing fails.
1. Write a `safeDiv(a, b int) (result int, err error)` that recovers from a divide-by-zero panic.

## Block 5: Interfaces and Polymorphism

### 5.1 Interface Basics

1. Define an interface `Shape` with `Area() float64` and `Perimeter() float64`. Implement it for `Circle` and `Rectangle`.
1. Show that implementation is implicit: a struct satisfies an interface without declaring it.
1. Write a function that accepts `interface{}` (or `any`) and prints the value and its type using `fmt.Sprintf("%T")`.
1. Use type assertion: given an `interface{}` value, assert it's a `string` and print it.
1. Use type assertion with the ok idiom: `s, ok := val.(string)`.
1. Write a function that accepts `interface{}` and uses a type switch to handle `int`, `string`, `bool`.

### 5.2 Working with Interfaces

1. Write a function `printAreas(shapes []Shape)` that prints the area of each shape.
1. Write a function that returns a `Shape` interface (return either `Circle` or `Rectangle` based on input).
1. Define `Reader` and `Writer` interfaces, then compose them into a `ReadWriter` interface.
1. Show that two interfaces with the same method set are interchangeable.

### 5.3 Standard Interfaces

1. Implement `io.Reader` for a custom type `InfiniteZeros` that always reads zero bytes.
1. Implement `io.Writer` for a type `UpperWriter` that converts all bytes to uppercase before writing to an underlying writer.
1. Implement `fmt.Stringer` for the `Person` struct to customize its printed output.
1. Implement the `error` interface for a custom `ValidationError` struct with `Field` and `Message`.
1. Implement `sort.Interface` for a `ByAge` type (`[]Person`) to sort people by age.

### 5.4 Internals

1. Show that a nil interface (`var s Shape`) is different from an interface holding a nil pointer (`var c *Circle; var s Shape = c`). Print both and compare to nil.
1. Demonstrate that calling a method on a nil concrete value inside an interface does not panic (if the method handles nil).
1. Print the dynamic type of an interface value using `fmt.Sprintf("%T", val)`.

## Block 6: Pointers and Memory

### 6.1 Pointers

1. Declare a pointer `var p *int`, print it (nil), then assign `p = &x` and dereference with `*p`.
1. Write a function `increment(p *int)` that increments the value at the pointer.
1. Show automatic dereferencing: given `p := &Person{Name: "Alice"}`, access `p.Name` without explicit `(*p).Name`.
1. Write a function `swap(a, b *int)` that swaps two values using pointers.
1. Show the difference between pass-by-value and pass-by-pointer: modify a struct in a function both ways.

### 6.2 Memory Allocation

1. Use `new(int)` to allocate an int, set its value through the pointer, and print it.
1. Use `make([]int, 5)` and `make(map[string]int)`. Show that `new` returns a pointer while `make` returns an initialized value.
1. Create a struct using a composite literal with `&`: `p := &Point{X: 1, Y: 2}`.

### 6.3 Memory Management

1. Write a function that returns a pointer to a local variable. Show that it works in Go (variable escapes to heap).
1. Run `go build -gcflags="-m"` on a small program and identify which variables escape to the heap.
1. Write a benchmark comparing allocating many small objects vs reusing them from a `sync.Pool`.

## Block 7: Packages and Modules

### 7.1 Packages

1. Create a package `mathutil` with exported functions `Add` and `Multiply`, and an unexported helper `validate`.
1. Import a package with an alias: `import m "myapp/mathutil"` and use `m.Add(1, 2)`.
1. Use a blank import `import _ "image/png"` and explain when this is useful.
1. Write an `init()` function in a package that prints "package initialized" to demonstrate init order.
1. Create an `internal` package and show that it cannot be imported from outside its parent tree.

### 7.2 Go Modules

1. Initialize a new module with `go mod init myapp`. Inspect the generated `go.mod`.
1. Add a dependency (e.g., `github.com/fatih/color`) with `go get` and examine `go.sum`.
1. Run `go mod tidy` and observe how unused dependencies are removed.
1. Use a `replace` directive in `go.mod` to point a dependency to a local path.

### 7.3 Project Organization

1. Create a project with `cmd/server/main.go` and `cmd/cli/main.go` producing two separate binaries.
1. Organize shared code into `internal/config` and `internal/database` packages.
1. Create a `pkg/validator` package intended for external use.

## Block 8: Error Handling

### 8.1 Basic Error Handling

1. Write a function `divide(a, b float64) (float64, error)` that returns `errors.New("division by zero")` if `b == 0`.
1. Write a function `parseAge(s string) (int, error)` that wraps `strconv.Atoi` errors with `fmt.Errorf("invalid age %q: %w", s, err)`.
1. Show the anti-pattern of ignoring errors: `result, _ := divide(10, 0)`.

### 8.2 Custom Errors

1. Define a `NotFoundError` struct with `Resource string` and `ID int` fields. Implement the `error` interface.
1. Define a `ValidationError` struct with `Field string` and `Message string`. Return it from a validation function.

### 8.3 Advanced Error Handling

1. Use `errors.Is` to check if an error in a chain matches `os.ErrNotExist`.
1. Use `errors.As` to extract a custom `*NotFoundError` from a wrapped error.
1. Wrap errors at each layer: `repository -> service -> handler` with `fmt.Errorf("...: %w", err)`.
1. Define sentinel errors `var ErrNotFound = errors.New("not found")` and `var ErrForbidden = errors.New("forbidden")` and check them with `errors.Is`.

## Block 9: Concurrency

### 9.1 Concurrency Basics

1. Write a program that prints "hello" from a goroutine and "world" from main. Use `time.Sleep` to observe interleaving.

### 9.2 Goroutines

1. Launch 5 goroutines, each printing its index. Use `sync.WaitGroup` to wait for all to finish.
1. Print `runtime.NumCPU()` and `runtime.NumGoroutine()` before and after launching goroutines.
1. Write a program that demonstrates a goroutine leak: launch a goroutine that blocks on a channel forever. Print goroutine count.
1. Fix the goroutine leak by using `context.WithCancel` to signal the goroutine to exit.

### 9.3 Channels

1. Create an unbuffered channel, send a value from a goroutine, receive it in main.
1. Create a buffered channel of size 3, send 3 values without blocking, then receive them.
1. Close a channel and show that receiving from a closed channel returns the zero value.
1. Use the two-value receive `v, ok := <-ch` to detect a closed channel.
1. Use `for range ch` to receive all values until the channel is closed.
1. Write a function that takes a `<-chan int` (receive-only) and a function that takes `chan<- int` (send-only).

### 9.4 Select

1. Use `select` to receive from whichever of two channels has data first.
1. Implement a timeout using `select` with `time.After(2 * time.Second)`.
1. Use `select` with a `default` case to attempt a non-blocking send.

### 9.5 Concurrency Patterns

1. Implement a worker pool: N workers read jobs from a channel and send results to another channel.
1. Implement fan-out/fan-in: split work across multiple goroutines, merge results into one channel.
1. Build a pipeline: `generator -> square -> filter(odd) -> print`.
1. Implement a rate limiter using `time.Ticker` that allows max 5 operations per second.
1. Implement a semaphore using a buffered channel to limit concurrency to 3 goroutines.

### 9.6 Sync Package

1. Use `sync.Mutex` to protect a shared counter incremented by 10 goroutines.
1. Use `sync.RWMutex`: multiple goroutines read a map concurrently, one goroutine writes.
1. Use `sync.WaitGroup` to wait for N goroutines to complete.
1. Use `sync.Once` to ensure an expensive initialization runs exactly once across goroutines.
1. Use `sync.Pool` to reuse `bytes.Buffer` objects in a concurrent program.
1. Use `sync.Map` instead of a regular map with mutex for concurrent access.

### 9.7 Atomic Operations

1. Use `atomic.AddInt64` to increment a counter from multiple goroutines without a mutex.
1. Use `atomic.LoadInt64` and `atomic.StoreInt64` for a shared flag variable.
1. Use `atomic.CompareAndSwapInt32` to implement a simple spinlock.

### 9.8 Context

1. Use `context.WithCancel` to cancel a long-running goroutine.
1. Use `context.WithTimeout` to set a 2-second timeout on a simulated database query.
1. Use `context.WithValue` to pass a request ID through a chain of function calls.
1. Write an HTTP handler that respects `r.Context()` cancellation.

### 9.9 Advanced Topics

1. Write a program with a data race (two goroutines writing to the same variable). Run with `go run -race` and observe the report.
1. Fix the data race using a mutex or channel.
1. Implement graceful shutdown: catch `SIGINT`/`SIGTERM` and wait for in-flight requests to finish.

## Block 10: Standard Library

### 10.1 fmt Package

1. Print a struct with `%v`, `%+v`, and `%#v` to see different levels of detail.
1. Use `fmt.Sprintf` to build a formatted string without printing it.
1. Use `fmt.Fprintf(os.Stderr, ...)` to write an error message to stderr.
1. Implement the `Stringer` interface on a custom type and print it with `%s` and `%v`.

### 10.2 strings Package

1. Check if a string contains a substring using `strings.Contains`.
1. Split a CSV line `"a,b,c,d"` using `strings.Split` and join it back with `strings.Join` using `" | "`.
1. Trim whitespace from `"  hello  "` using `strings.TrimSpace`.
1. Use `strings.NewReplacer` to replace multiple substrings at once.
1. Build a large string efficiently using `strings.Builder` in a loop (compare with `+=` concatenation).

### 10.3 strconv Package

1. Convert `"42"` to an int with `strconv.Atoi` and back with `strconv.Itoa`.
1. Parse `"3.14"` as `float64` with `strconv.ParseFloat`.
1. Format a float with `strconv.FormatFloat(3.14159, 'f', 2, 64)` to get `"3.14"`.

### 10.4 math Package

1. Compute the hypotenuse of a right triangle using `math.Sqrt(a*a + b*b)`.
1. Generate 10 random integers between 1 and 100 using `math/rand`.
1. Round `3.7` up (`math.Ceil`), down (`math.Floor`), and to nearest (`math.Round`).

### 10.5 time Package

1. Print the current time with `time.Now()` and format it as `"2006-01-02 15:04:05"`.
1. Parse `"2024-12-25"` into a `time.Time` using `time.Parse`.
1. Calculate the duration between two dates.
1. Write a function that measures execution time of another function using `time.Since`.
1. Use `time.NewTicker` to print a message every 500ms for 3 seconds.

### 10.6 sort Package

1. Sort a slice of ints: `sort.Ints([]int{5, 3, 1, 4, 2})`.
1. Sort a slice of structs by a field using `sort.Slice`.
1. Implement `sort.Interface` for a custom type and sort it.
1. Use `sort.Search` to find the insertion point for a value in a sorted slice.

### 10.7 regexp Package

1. Check if a string matches an email pattern using `regexp.MatchString`.
1. Extract all numbers from a string using `regexp.MustCompile(`\d+`).FindAllString`.
1. Replace all whitespace sequences with a single space using `regexp.MustCompile`.

## Block 11: Files and I/O

### 11.1 os Package

1. Create a file, write "Hello, File!" to it, close it, then read it back.
1. Check if a file exists using `os.Stat` and `os.IsNotExist`.
1. List all files in a directory using `os.ReadDir`.
1. Get and set an environment variable using `os.Getenv` and `os.Setenv`.
1. Create a nested directory structure with `os.MkdirAll`.

### 11.2 io Package

1. Copy the contents of one file to another using `io.Copy`.
1. Use `io.MultiWriter` to write the same data to both a file and stdout.
1. Use `io.TeeReader` to read from a file and simultaneously write to stdout.
1. Use `io.Pipe` to connect a writer and reader in separate goroutines.

### 11.3 bufio Package

1. Read a file line-by-line using `bufio.Scanner`.
1. Read user input word-by-word using `bufio.Scanner` with `bufio.ScanWords`.
1. Write multiple lines to a file using `bufio.Writer` and call `Flush()` at the end.

### 11.4 ioutil (deprecated alternatives)

1. Read an entire file with `os.ReadFile` (the modern replacement for `ioutil.ReadFile`).
1. Write data to a file with `os.WriteFile`.
1. Create a temporary file with `os.CreateTemp` and a temp directory with `os.MkdirTemp`.

### 11.5 os/exec Package

1. Run `ls -la` (or `dir` on Windows) using `exec.Command` and print the output.
1. Run a command with a timeout using `context.WithTimeout` and `exec.CommandContext`.
1. Pipe input to a command's stdin and read from its stdout.

### 11.6 path/filepath Package

1. Join path components using `filepath.Join("usr", "local", "bin")`.
1. Get the extension of a filename with `filepath.Ext`.
1. Walk a directory tree with `filepath.WalkDir` and print all `.go` files.

### 11.7 embed Package

1. Embed a single text file into a binary using `//go:embed` and print its content at runtime.
1. Embed an entire directory using `embed.FS` and serve its contents via `http.FileServer`.

## Block 12: Testing

### 12.1 testing Package

1. Write a function `Add(a, b int) int` and a test `TestAdd` that checks several cases.
1. Use `t.Errorf` to report a non-fatal error and `t.Fatalf` to report a fatal one.
1. Write subtests with `t.Run("case name", func(t *testing.T) { ... })`.
1. Write a table-driven test for a `FizzBuzz(n int) string` function with at least 10 test cases.
1. Run tests with `go test -v -run TestAdd` to filter and see verbose output.

### 12.2 Benchmarks

1. Write a benchmark `BenchmarkFibonacci` that benchmarks both recursive and iterative fibonacci.
1. Use `b.ResetTimer()` after setup code in a benchmark.
1. Run `go test -bench=. -benchmem` and interpret the output (ns/op, allocs/op).

### 12.3 Examples

1. Write an `ExampleAdd` function with an `// Output:` comment and run it as a test.

### 12.4 Coverage

1. Run `go test -cover` on a package and check the coverage percentage.
1. Generate a coverage profile with `-coverprofile=cover.out` and view it with `go tool cover -html=cover.out`.

### 12.5 Test Helpers

1. Write a helper function `assertEqual(t *testing.T, got, want int)` that calls `t.Helper()` for better error reporting.
1. Write a `TestMain(m *testing.M)` function for setup/teardown of the entire test suite.

### 12.6 Mocking

1. Define an interface `UserRepository` with `FindByID(id int) (*User, error)`. Create a mock implementation for testing.
1. Write a `UserService` that depends on `UserRepository` and test it with the mock.

### 12.7 Integration Testing

1. Write an integration test that starts an HTTP server, sends a request, and checks the response (use `httptest.NewServer`).

## Block 13: Advanced Language Features

### 13.1 Generics

1. Write a generic function `Max[T int | float64 | string](a, b T) T`.
1. Write a generic function `Contains[T comparable](slice []T, item T) bool`.
1. Write a generic `Stack[T any]` type with `Push`, `Pop`, and `IsEmpty` methods.
1. Define a custom constraint `Number` that includes all integer and float types.
1. Show type inference: call `Max(3, 5)` without explicitly specifying the type parameter.

### 13.2 Reflection

1. Print the type and kind of various values using `reflect.TypeOf` and `reflect.ValueOf(...).Kind()`.
1. Write a function that takes any struct and prints all field names and values using reflection.
1. Write a function that sets a struct field by name using reflection (field must be exported and addressable).
1. Write a simple struct-to-map converter using reflection.

### 13.3 Unsafe

1. Use `unsafe.Sizeof` to print the size of various types.
1. Use `unsafe.Pointer` to convert between `*int` and `*float64` (understand this is dangerous).
1. Use `unsafe.Offsetof` to find the byte offset of a field in a struct.

### 13.4 Build Tags

1. Create two files `greet_linux.go` and `greet_darwin.go` with `//go:build` tags that return different greetings.
1. Create a file with a custom build tag `//go:build debug` and build with `go build -tags debug`.

### 13.5 Code Generation

1. Use `go generate` with `stringer` to generate `String()` methods for an enum type.
1. Write a simple Go program that generates Go source code using `text/template`.

## Block 14: Network Programming

### 14.1 net Package

1. Resolve a hostname to IP addresses using `net.LookupHost("google.com")`.
1. Create a `net.TCPAddr` and print it.

### 14.2 TCP

1. Write a TCP echo server that reads a line from the client and sends it back.
1. Write a TCP client that connects to the echo server and sends a message.
1. Modify the server to handle multiple connections concurrently using goroutines.

### 14.3 UDP

1. Write a simple UDP server and client that exchange a single message.

### 14.4 HTTP Client

1. Make a GET request to a public API (e.g., `https://httpbin.org/get`) and print the response body.
1. Make a POST request with a JSON body using `http.NewRequest`.
1. Set custom headers (e.g., `Authorization`, `Content-Type`) on a request.
1. Create an `http.Client` with a 5-second timeout.

### 14.5 HTTP Server

1. Write an HTTP server with two endpoints: `GET /hello` returns "Hello, World!" and `GET /time` returns the current time.
1. Write an HTTP handler that reads a JSON body and responds with JSON.
1. Write a logging middleware that logs the method, path, and duration of each request.
1. Implement graceful shutdown using `server.Shutdown(ctx)`.

### 14.6 Advanced HTTP

1. Serve a static directory of files over HTTP.
1. Implement a file upload endpoint that saves uploaded files to disk.
1. Write an SSE (Server-Sent Events) endpoint that sends a timestamp every second.

### 14.7 WebSocket

1. Write a WebSocket echo server using `gorilla/websocket`.
1. Write a client that connects to the WebSocket server and sends/receives messages.

### 14.8 DNS

1. Look up the MX records for a domain using `net.LookupMX`.
1. Look up the IP addresses for a hostname and print the results.

## Block 15: REST API Development

### 15.1 REST Basics

1. Write a REST API with CRUD endpoints for a `Book` resource (in-memory storage):
   - `GET /books` - list all books
   - `GET /books/:id` - get a book by ID
   - `POST /books` - create a book
   - `PUT /books/:id` - update a book
   - `DELETE /books/:id` - delete a book

### 15.2 JSON

1. Marshal a struct to JSON with `json.Marshal` and pretty-print with `json.MarshalIndent`.
1. Unmarshal a JSON string into a struct.
1. Use `json.Decoder` to decode JSON from an `io.Reader` (e.g., `http.Request.Body`).
1. Use `omitempty` to skip zero-value fields during marshaling.
1. Implement `json.Marshaler` for a custom type that serializes a `time.Time` as a Unix timestamp.
1. Use `json.RawMessage` to delay parsing of a nested JSON field.

### 15.3 HTTP Routers

1. Build the Book CRUD API using the standard `http.ServeMux`.
1. Rebuild it using `chi` router with URL parameters (`chi.URLParam`).

### 15.4 Gin Framework

1. Build the Book CRUD API using Gin.
1. Add request validation using Gin's binding tags.
1. Add a Gin middleware that checks for an `X-API-Key` header.

### 15.5 go-chi Router

1. Build the Book CRUD API using chi with subrouters for `/api/v1/books`.
1. Add chi middleware for logging and request ID.

### 15.6 API Best Practices

1. Add pagination to the `GET /books` endpoint (query params `page` and `limit`).
1. Add sorting to the `GET /books` endpoint (query param `sort=title` or `sort=-title` for desc).
1. Implement a simple rate limiter middleware that limits to 10 requests per minute per IP.
1. Add CORS middleware to allow requests from any origin.

### 15.7 OpenAPI Specification

1. Write an OpenAPI 3.0 YAML spec for the Book CRUD API.

### 15.8 Validation

1. Use `go-playground/validator` to validate a struct: `Title` required, `Pages` min 1, `ISBN` matches a regex.
1. Return structured validation error responses with field names and messages.

## Block 16: gRPC

### 16.1 gRPC Basics

1. Write a comparison table (as comments in code) of REST vs gRPC for different use cases.

### 16.2 Protocol Buffers

1. Write a `.proto` file defining a `UserService` with `GetUser`, `CreateUser`, and `ListUsers` RPCs.
1. Generate Go code from the `.proto` file using `protoc`.

### 16.3 gRPC Server

1. Implement the `UserService` gRPC server with in-memory storage.
1. Add a server streaming RPC: `ListUsers` that streams users one by one.

### 16.4 gRPC Client

1. Write a gRPC client that calls `CreateUser` and `GetUser`.
1. Write a client that consumes a server stream from `ListUsers`.

### 16.5 Advanced gRPC

1. Write a unary interceptor that logs every RPC call with its duration.
1. Add metadata (like an auth token) to a gRPC request from the client.
1. Implement deadline propagation: set a deadline on the client and check it on the server.

## Block 17: Databases

### 17.1 SQL Basics

1. Write SQL queries for a `users` table: SELECT with JOIN, GROUP BY with HAVING, and a subquery.

### 17.2 database/sql Package

1. Connect to a SQLite database (or PostgreSQL) using `database/sql` and `sql.Open`.
1. Execute a `CREATE TABLE` statement with `db.Exec`.
1. Insert a row with parameterized query (`db.Exec` with `?` or `$1` placeholders).
1. Query rows with `db.Query` and scan results into struct fields.
1. Use `db.QueryRow` for a single-row query.

### 17.3 PostgreSQL

1. Connect to PostgreSQL using `pgx` driver and perform CRUD operations.
1. Write a transaction: transfer money between two accounts atomically (BEGIN, UPDATE, COMMIT/ROLLBACK).
1. Use prepared statements for repeated queries.
1. Handle NULL values with `sql.NullString`, `sql.NullInt64`.

### 17.4 ORM - GORM

1. Define GORM models for `User` and `Post` (one-to-many relationship).
1. Use `AutoMigrate` to create tables.
1. Perform CRUD operations using GORM methods.
1. Use `Preload` to eager-load related records.
1. Wrap multiple operations in a GORM transaction.

### 17.5 sqlx

1. Use `sqlx.Get` to query a single row into a struct.
1. Use `sqlx.Select` to query multiple rows into a slice of structs.
1. Use named queries with `sqlx.NamedExec`.

### 17.6 Migrations

1. Create up/down migration files for a `users` table using `golang-migrate`.
1. Write a migration that adds a column and a migration that creates an index.

### 17.7 Redis

1. Connect to Redis using `go-redis` and perform `Set`/`Get` operations with TTL.
1. Use Redis as a cache: check cache before querying the database.
1. Implement a simple pub/sub: one goroutine publishes, another subscribes and prints messages.

### 17.8 MongoDB

1. Connect to MongoDB using the official driver and insert a document.
1. Query documents with filters and projections.
1. Update and delete documents.

## Block 18: Message Queues

### 18.1 Apache Kafka

1. Write a Kafka producer that sends 10 messages to a topic using `kafka-go`.
1. Write a Kafka consumer that reads messages from the topic and prints them.
1. Write a consumer group with two consumers that share partitions.

### 18.2 RabbitMQ

1. Write a RabbitMQ producer that publishes messages to a fanout exchange.
1. Write a consumer that declares a queue, binds it to the exchange, and consumes messages.
1. Use a direct exchange with routing keys to route messages to different queues.

## Block 19: Architecture and Design Patterns

### 19.1 SOLID Principles

1. Refactor a `UserService` that handles validation, storage, and email sending into separate responsibilities (SRP).
1. Write a notification system using interfaces so that adding a new channel (SMS) doesn't modify existing code (OCP).
1. Define small interfaces: `Saver`, `Loader`, `Deleter` instead of one large `Repository` interface (ISP).
1. Write a `Service` that depends on an interface, not a concrete type. Inject the dependency via constructor (DIP).

### 19.2 KISS, DRY, YAGNI

1. Refactor duplicated error-handling code in three HTTP handlers into a helper function (DRY).
1. Take an over-engineered function with unused parameters and simplify it (KISS).

### 19.3 Design Patterns

1. Implement a Singleton for a database connection pool using `sync.Once`.
1. Implement a Factory function `NewShape(kind string) Shape` that returns `Circle` or `Rectangle`.
1. Implement a Builder for constructing a complex `Config` struct with optional fields.
1. Implement the Strategy pattern: a `Sorter` that accepts different sorting strategies.
1. Implement the Observer pattern: an `EventBus` where listeners subscribe to events.
1. Implement the Decorator pattern: wrap an `http.Handler` with logging and auth decorators.

### 19.4 Architectural Patterns

1. Structure a project following Clean Architecture: `domain`, `usecase`, `repository`, `delivery` layers.
1. Implement a simple CQRS example: separate read and write models for a `Product`.

### 19.5 Dependency Injection

1. Write constructor injection: `NewUserService(repo UserRepository, logger Logger) *UserService`.
1. Write a simple DI container that resolves dependencies by type.

### 19.6 Microservices

1. Design two microservices (`OrderService` and `InventoryService`) that communicate via HTTP.
1. Implement a circuit breaker that opens after 3 consecutive failures and retries after a cooldown.

## Block 20: Logging and Monitoring

### 20.1 Logging

1. Use the standard `log` package to log messages with timestamps.
1. Set up `zap` logger with JSON output and log at different levels (Info, Warn, Error).
1. Add context fields to log entries (e.g., request ID, user ID) using zap's `With`.
1. Set up `zerolog` with pretty console output for development and JSON for production.

### 20.2 ELK Stack

1. Configure a Go app to output JSON logs compatible with Elasticsearch ingestion.

### 20.3 Metrics

1. Expose Prometheus metrics from a Go HTTP server using `prometheus/client_golang`.
1. Create a Counter for total HTTP requests and a Histogram for request duration.
1. Add labels to metrics (method, path, status code).

### 20.4 Tracing

1. Add OpenTelemetry tracing to an HTTP server: create spans for each handler.
1. Propagate trace context between two services via HTTP headers.

### 20.5 Healthchecks

1. Add `/healthz` (liveness) and `/readyz` (readiness) endpoints to an HTTP server.
1. Make the readiness probe check database connectivity before returning 200.

## Block 21: Docker and Containerization

### 21.1 Docker Basics

1. Write a basic Dockerfile for a Go application using `golang:latest` as the base image.

### 21.2 Dockerfile for Go

1. Write a multi-stage Dockerfile: build stage with `golang:1.22` and runtime stage with `alpine:3.19`.
1. Add a `.dockerignore` file that excludes `.git`, `vendor`, and test files.
1. Optimize layer caching by copying `go.mod` and `go.sum` before copying source code.

### 21.3 Docker Compose

1. Write a `docker-compose.yml` that runs your Go app with PostgreSQL and Redis.
1. Add health checks and dependency ordering (`depends_on` with `condition: service_healthy`).
1. Use environment variables for database credentials.

### 21.4 Docker Best Practices

1. Modify the Dockerfile to run as a non-root user.
1. Add a `HEALTHCHECK` instruction to the Dockerfile.

## Block 22: Kubernetes

### 22.1 Kubernetes Basics

1. Write a Pod manifest YAML for your Go application.

### 22.2 Kubernetes Resources

1. Write a Deployment manifest with 3 replicas.
1. Write a Service manifest (ClusterIP) to expose the deployment.
1. Write ConfigMap and Secret manifests for application configuration.

### 22.3 kubectl

1. List commands to deploy an app: `kubectl apply`, `kubectl get pods`, `kubectl logs`, `kubectl port-forward`.

### 22.4 YAML Manifests

1. Write a complete set of manifests (Deployment + Service + ConfigMap + Secret) for your Go app.
1. Add liveness and readiness probes to the Deployment.

### 22.5 Helm

1. Create a basic Helm chart structure for your Go application with `values.yaml`.

### 22.6 Kubernetes for Go Apps

1. Add resource limits and requests to the Deployment manifest.
1. Add a HorizontalPodAutoscaler manifest targeting 70% CPU utilization.

## Block 23: CI/CD

### 23.1 Git

1. Write a `.gitignore` file for a Go project.
1. Practice creating a feature branch, making commits, and opening a PR.

### 23.2 GitLab CI/CD

1. Write a `.gitlab-ci.yml` with stages: `lint`, `test`, `build`, `deploy`.
1. Add caching for Go modules.

### 23.3 CI/CD for Go

1. Write a CI job that runs `golangci-lint run`.
1. Write a CI job that runs tests with coverage and fails if coverage is below 80%.
1. Write a CI job that builds a Docker image and pushes it to a registry.

### 23.4 GitHub Actions

1. Write a `.github/workflows/ci.yml` that lints, tests, and builds on push to main and on PRs.
1. Add a step to upload test coverage as an artifact.

## Block 24: Profiling and Optimization

### 24.1 Benchmarking

1. Write benchmarks comparing `fmt.Sprintf` vs `strconv.Itoa` vs `strings.Builder` for int-to-string conversion.
1. Use `benchstat` to compare benchmark results before and after an optimization.

### 24.2 Profiling

1. Add `net/http/pprof` to a running server and capture a CPU profile.
1. Capture a memory profile and identify the top memory-allocating functions.

### 24.3 pprof Tool

1. Analyze a CPU profile with `go tool pprof`: use `top`, `list`, and `web` commands.
1. Generate a flame graph from a CPU profile.

### 24.4 Tracing

1. Use `runtime/trace` to trace a program and analyze it with `go tool trace`.

### 24.5 Optimization

1. Optimize a function that concatenates strings in a loop (replace `+=` with `strings.Builder`).
1. Optimize a function that allocates slices in a loop (pre-allocate with `make([]T, 0, n)`).
1. Use `sync.Pool` to reuse buffers and benchmark the improvement.

### 24.6 Escape Analysis

1. Run `go build -gcflags="-m"` on a function and identify heap allocations.
1. Refactor a function to reduce heap allocations (return value instead of pointer where possible).

## Block 25: Advanced Topics and Internals

### 25.1 Go Scheduler

1. Write a program that sets `GOMAXPROCS(1)` and launches CPU-bound goroutines. Observe that they run sequentially.
1. Increase `GOMAXPROCS` and observe parallel execution.

### 25.2 Garbage Collector

1. Write a program that allocates many objects, then use `runtime.ReadMemStats` to print GC stats.
1. Experiment with `GOGC` environment variable: set it to 50 and 200, observe GC frequency changes.

### 25.3 Memory Model

1. Write a program that demonstrates happens-before: use a channel to synchronize two goroutines writing and reading a shared variable.
1. Show that without synchronization, a reader goroutine may see a stale value.

### 25.4 Compiler

1. Disassemble a simple function with `go tool objdump` and identify the generated instructions.

### 25.5 Assembly

1. Use `go tool compile -S` to view the assembly output of a simple Go function.

### 25.6 Runtime

1. Use `runtime.ReadMemStats` to print heap size, number of GC cycles, and pause times.
1. Use `runtime.Stack` to print the stack trace of all goroutines.

## Block 26: Security

### 26.1 Cryptography

1. Hash a password using SHA-256 and print the hex-encoded result.
1. Generate an HMAC-SHA256 signature for a message and verify it.
1. Encrypt and decrypt a message using AES-GCM.
1. Generate an RSA key pair, sign a message, and verify the signature.

### 26.2 Authentication

1. Implement Basic Auth middleware that checks credentials from the `Authorization` header.
1. Generate and validate JWT tokens using a library like `golang-jwt/jwt`.
1. Implement token-based auth: `/login` returns a token, other endpoints require it.

### 26.3 Authorization

1. Implement a simple RBAC middleware: define roles (admin, user) and restrict endpoints by role.
1. Write a middleware that extracts the user role from a JWT and checks permissions.

### 26.4 Application Security

1. Show SQL injection vulnerability with string concatenation, then fix it with parameterized queries.
1. Sanitize HTML input to prevent XSS using `html.EscapeString` or a sanitization library.
1. Implement CSRF protection using a token stored in a cookie and validated in form submissions.
1. Write a rate limiter middleware using `golang.org/x/time/rate`.

## Block 27: Linux and System Programming

### 27.1 Linux Basics

1. Write a Go program that lists all files in a directory with their permissions.
1. Read environment variables and print them formatted as `KEY=VALUE`.

### 27.2 Unix Tools

1. Write a Go program that mimics `grep`: search for a pattern in a file and print matching lines.
1. Write a Go program that mimics `wc`: count lines, words, and bytes in a file.

### 27.3 System Programming in Go

1. Handle `SIGINT` and `SIGTERM` signals using `os/signal.Notify` and exit gracefully.
1. Write a program that launches a child process, waits for it, and prints its exit code.

### 27.4 Make

1. Write a `Makefile` with targets: `build`, `test`, `lint`, `run`, `clean`, and `docker-build`.
1. Add a `help` target that prints all available targets with descriptions.

## Block 28: Development Tools

### 28.1 Go Toolchain

1. Use `go build -o myapp` to build a binary with a custom name.
1. Use `go vet` to find suspicious constructs in your code.
1. Use `go doc fmt.Println` to read documentation from the terminal.
1. Use `go list -m all` to list all dependencies.

### 28.2 Linters

1. Install and run `golangci-lint run` on your project.
1. Create a `.golangci.yml` configuration that enables `govet`, `staticcheck`, `errcheck`, and `gosimple`.
1. Fix all lint issues reported by `golangci-lint`.

### 28.3 Code Formatting

1. Run `gofmt -d .` to see formatting differences, then `gofmt -w .` to fix them.
1. Run `goimports -w .` to fix imports and formatting simultaneously.

### 28.4 Documentation

1. Write godoc-style comments for a package and its exported functions.
1. Run `go doc` on your package and verify the documentation output.

### 28.5 IDE and Editors

1. Configure VS Code with the Go extension: enable format-on-save, lint-on-save, and test-on-save.
1. Set up debugging in VS Code: create a `launch.json` and set breakpoints.

## Block 29: Go Best Practices and Idioms

### 29.1 Effective Go

1. Refactor a piece of code to follow Go naming conventions: exported types in PascalCase, unexported in camelCase.
1. Replace a complex `if-else` chain with a `switch` or early returns.
1. Refactor a function with 5+ parameters to accept an options struct.

### 29.2 Code Review Comments

1. Review a code snippet and fix issues: missing error checks, wrong receiver type, non-idiomatic naming.
1. Refactor a function to follow "indent error flow" (handle errors first, keep the happy path unindented).
1. Fix a test that gives unhelpful failure messages: add `got` and `want` to assertion messages.

### 29.3 Go Proverbs

1. Refactor shared-memory communication to use channels ("share memory by communicating").
1. Replace a large interface with several small ones ("the bigger the interface, the weaker the abstraction").
1. Show an example where copying a small function is better than importing a dependency ("a little copying is better than a little dependency").

### 29.4 Common Mistakes

1. Fix a bug caused by range loop variable capture in a goroutine (pre-Go 1.22 behavior).
1. Fix a bug caused by modifying a slice while iterating over it.
1. Fix a deferred call in a loop that leaks resources (move the body to a separate function).
1. Fix a goroutine that forgets to unlock a mutex (use `defer mu.Unlock()` right after `mu.Lock()`).
1. Fix a program that closes a channel from the receiver side (move close to the sender).

## Block 30: Interview Preparation

### 30.1 Algorithms and Data Structures

1. Implement a singly linked list with `Insert`, `Delete`, `Find`, and `Print` methods.
1. Implement a stack using a slice with `Push`, `Pop`, `Peek`, and `IsEmpty`.
1. Implement a queue using a slice (or linked list) with `Enqueue`, `Dequeue`, and `IsEmpty`.
1. Implement a hash table with separate chaining for collision resolution.
1. Implement a binary search tree with `Insert`, `Search`, `Delete`, and in-order traversal.
1. Implement BFS and DFS for a graph represented as an adjacency list.
1. Implement a min-heap with `Insert`, `ExtractMin`, and `Peek`.
1. Implement a Trie with `Insert`, `Search`, and `StartsWith`.
1. Implement quicksort and mergesort for a slice of ints.
1. Implement binary search on a sorted slice.
1. Solve the classic dynamic programming problems: Fibonacci (memoized), knapsack (0/1), longest common subsequence.
1. Solve a backtracking problem: generate all permutations of a slice.
1. Implement Dijkstra's shortest path algorithm.

### 30.2 System Design

1. Design a URL shortener: describe the components (API, storage, encoding), write the Go code for the API layer.
1. Design a rate limiter: implement a token bucket algorithm in Go.
1. Design a key-value store with TTL support: implement `Set(key, value, ttl)` and `Get(key)`.
1. Design a simple load balancer: implement round-robin and random selection strategies.
1. Design a basic message queue: implement in-memory pub/sub with topics in Go.

## Block 31: Production Readiness and Resilience

### 31.1 Graceful Lifecycle

1. Write an HTTP server that drains in-flight requests on `SIGTERM`: stop accepting new connections, wait up to 30 seconds for active handlers to finish, then exit.
1. Implement a worker that processes jobs from a channel and shuts down cleanly when a context is cancelled, finishing the current job before exiting.
1. Write a service that registers multiple shutdown hooks (close DB, flush logs, close message broker) and executes them in reverse order on termination.
1. Implement a startup sequence that checks dependencies (database, cache, external API) before marking the service as ready via a `/readyz` endpoint.

### 31.2 Retry and Backoff

1. Write a function `Retry(ctx context.Context, maxAttempts int, fn func() error) error` with exponential backoff between attempts.
1. Add jitter to the backoff: implement both full jitter and decorrelated jitter strategies and compare them.
1. Write a retry function that distinguishes between retryable errors (network timeout) and permanent errors (400 Bad Request) using a custom `RetryableError` interface.
1. Implement a retry with a maximum total duration using `context.WithTimeout` so that retries stop after a deadline regardless of remaining attempts.

### 31.3 Circuit Breaker

1. Implement a circuit breaker with three states (Closed, Open, HalfOpen) that opens after N consecutive failures and attempts recovery after a cooldown.
1. Add a success threshold to the HalfOpen state: require M consecutive successes before transitioning back to Closed.
1. Write an HTTP client wrapper that uses the circuit breaker for outgoing requests and returns a `ErrCircuitOpen` immediately when open.
1. Add metrics to the circuit breaker: track total calls, failures, state transitions, and expose them via a `Stats()` method.

### 31.4 Bulkhead and Load Shedding

1. Implement a bulkhead pattern using a buffered channel as a semaphore to limit concurrent access to a downstream service.
1. Write a middleware that rejects requests with `503 Service Unavailable` when the in-flight request count exceeds a threshold.
1. Implement adaptive load shedding: measure p99 latency with a sliding window and start rejecting a percentage of requests when latency exceeds a target.

### 31.5 Singleflight and Deduplication

1. Use `golang.org/x/sync/singleflight` to deduplicate concurrent requests for the same cache key to a database.
1. Write a benchmark comparing a cache lookup with and without `singleflight` under concurrent access to demonstrate thundering herd prevention.
1. Implement request deduplication for an API endpoint using an idempotency key stored in a `sync.Map` with TTL-based expiration.

### 31.6 Timeouts and Deadlines

1. Write an HTTP handler chain where each layer (handler, service, repository) propagates the request context and respects its deadline.
1. Implement a function that makes two parallel HTTP calls using `errgroup` and returns the first successful result, cancelling the other via context.
1. Write a test that verifies a slow dependency causes a timeout error, not a hang, by using `context.WithTimeout` and asserting the error is `context.DeadlineExceeded`.

## Block 32: Distributed Systems Patterns

### 32.1 Distributed Locking

1. Implement a distributed lock using Redis `SET NX EX` with `go-redis`. Acquire the lock, perform work, release it.
1. Add lock renewal: write a goroutine that extends the lock TTL periodically while the work is in progress.
1. Handle the case where a lock expires before work completes: use a fencing token to prevent stale writes.
1. Implement a distributed lock using etcd with lease and `KeepAlive`.

### 32.2 Leader Election

1. Implement leader election using etcd: multiple instances compete for a key, the winner becomes the leader.
1. Write a service that performs periodic work (cron-like) only on the leader node and gracefully hands off leadership on shutdown.

### 32.3 Saga Pattern

1. Implement a saga for an order flow: `CreateOrder -> ReserveInventory -> ChargePayment`. Each step has a compensating action.
1. Write compensating transactions that undo previous steps when a later step fails (e.g., payment fails - release inventory, cancel order).
1. Add a saga log that records each step's status so that a crashed saga can be recovered and continued or compensated.

### 32.4 Idempotency

1. Implement an idempotent `POST /orders` endpoint: accept an `Idempotency-Key` header, store the result, return the cached response on duplicate requests.
1. Write a message consumer that processes each message exactly once using a deduplication table keyed by message ID.
1. Implement an idempotent database operation using `INSERT ... ON CONFLICT DO NOTHING` and return the existing row if it already exists.

### 32.5 Event Sourcing

1. Implement an event store: append events (`OrderCreated`, `ItemAdded`, `OrderPaid`) to a slice and rebuild the current state by replaying them.
1. Write a projection that listens to events and maintains a read-optimized view (e.g., `OrderSummary` with total price and item count).
1. Implement snapshotting: save the aggregate state every N events and replay only events after the snapshot.

### 32.6 CQRS

1. Separate write and read models for a `Product`: the write model validates and stores events, the read model is a denormalized struct for fast queries.
1. Write a command handler that accepts `CreateProduct` and `UpdatePrice` commands, validates them, and emits events.
1. Write a query handler that reads from the projection and supports filtering and pagination.

### 32.7 Consistency Patterns

1. Implement the Outbox pattern: write a domain event to an `outbox` table in the same transaction as the business data, then publish it asynchronously.
1. Write a poller that reads unpublished events from the outbox table and publishes them to a message broker, marking them as published.
1. Implement optimistic concurrency control: add a `version` field to a struct, increment it on update, and reject updates with a stale version.

## Block 33: Advanced Testing

### 33.1 Fuzz Testing

1. Write a fuzz test `FuzzParseEmail(f *testing.F)` for a function that parses email addresses. Add seed corpus entries and run with `go test -fuzz`.
1. Write a fuzz test for a JSON unmarshaling function: ensure it never panics on arbitrary input.
1. Write a fuzz test for a URL parser that checks round-trip consistency: `Parse(url).String() == url`.

### 33.2 Testcontainers

1. Write an integration test that starts a PostgreSQL container using `testcontainers-go`, runs migrations, and tests CRUD operations.
1. Write an integration test that starts a Redis container and tests cache `Set`/`Get` with TTL expiration.
1. Write a test suite that starts multiple containers (PostgreSQL + Redis) using a shared `TestMain` setup and tears them down after all tests.

### 33.3 Golden Files

1. Write a test that compares JSON output of a function against a golden file in `testdata/`. Use `-update` flag to regenerate golden files.
1. Write a golden file test for an HTML template renderer that detects unintended changes in output.

### 33.4 Contract Testing

1. Define a `UserRepository` interface and write a contract test suite that can be run against both a mock and a real implementation.
1. Write an HTTP API contract test: start the server, send requests, and verify responses match a predefined schema.

### 33.5 Load Testing

1. Write a load test using `vegeta` library: send 100 RPS to an endpoint for 30 seconds and print latency percentiles.
1. Write a custom load test in pure Go: spawn N goroutines, each sending requests in a loop, and collect latency histograms using `sync.Mutex`.
1. Write a load test that ramps up from 10 to 500 RPS over 60 seconds and detects the saturation point where p99 latency exceeds 1 second.

### 33.6 Test Patterns

1. Write a test helper `testDB(t *testing.T) *sql.DB` that creates a temporary test database, runs migrations, and registers a cleanup function with `t.Cleanup`.
1. Write a table-driven test with subtests where each case uses a different mock behavior (success, not found, internal error).
1. Write a test that uses `t.Parallel()` and verify that shared state is properly isolated between parallel subtests.
1. Write a test that verifies goroutine cleanup: check `runtime.NumGoroutine()` before and after the test to detect leaks.

## Block 34: CLI Tools and Configuration

### 34.1 CLI with Cobra

1. Build a CLI application using `cobra` with a root command and two subcommands: `serve` (starts an HTTP server) and `migrate` (runs DB migrations).
1. Add persistent flags (`--config`, `--verbose`) to the root command and local flags (`--port`) to the `serve` subcommand.
1. Add shell autocompletion generation: `myapp completion bash/zsh/fish`.
1. Write a `version` subcommand that prints the build version, commit hash, and build date injected via `-ldflags`.

### 34.2 Configuration Management

1. Use `viper` to load configuration from a YAML file, environment variables, and command-line flags with priority: flags > env > file > defaults.
1. Define a `Config` struct with nested fields and unmarshal the configuration into it using `viper.Unmarshal`.
1. Implement configuration validation: check that required fields are set and values are within expected ranges, return a descriptive error.
1. Implement hot-reload of configuration: watch the config file with `viper.WatchConfig` and update the running application without restart.

### 34.3 Feature Flags

1. Implement a simple in-memory feature flag system: `IsEnabled(flag string, userID string) bool` with percentage-based rollout.
1. Add a feature flag middleware that injects enabled features into the request context.
1. Write a feature flag evaluation function that supports targeting rules: by user ID, by user attribute (country, plan), and by percentage.

### 34.4 Go Workspaces

1. Create a multi-module workspace with `go work init`. Add two modules: `api` and `shared`. Import `shared` from `api` without publishing.
1. Add a third module `cli` to the workspace that also depends on `shared`. Build all modules with `go build ./...`.
1. Show how to replace a dependency with a local module using `go.work` instead of `replace` directives in `go.mod`.

### 34.5 Build and Release

1. Use `-ldflags` to inject version, commit hash, and build time into a Go binary at compile time. Print them from the `version` command.
1. Cross-compile a Go binary for `linux/amd64`, `linux/arm64`, and `darwin/arm64` using `GOOS` and `GOARCH` environment variables.
1. Write a `goreleaser` configuration (`.goreleaser.yaml`) that builds binaries for multiple platforms and creates a GitHub release with changelogs.

## Block 35: Performance Engineering

### 35.1 GC Tuning

1. Write a program that allocates millions of small objects, print GC stats with `runtime.ReadMemStats`, then set `GOGC=50` and `GOGC=200` and compare GC frequency and pause times.
1. Use `GOMEMLIMIT` (Go 1.19+) to set a soft memory limit. Write a program that allocates memory up to the limit and observe how the GC becomes more aggressive.
1. Write a program that triggers frequent GC pauses. Use `GODEBUG=gctrace=1` to trace GC behavior and identify the cause.

### 35.2 Zero-Allocation Patterns

1. Write a benchmark for a function that formats log lines. Optimize it from `fmt.Sprintf` to `strconv.AppendInt` + `[]byte` buffer to achieve zero allocations.
1. Rewrite a function that returns `[]byte` to accept a caller-provided buffer parameter, eliminating internal allocations. Benchmark both versions.
1. Use `sync.Pool` to reuse `bytes.Buffer` instances in a high-throughput HTTP handler. Benchmark with and without the pool under concurrent load.
1. Write a struct method that avoids interface boxing: compare `fmt.Stringer` dispatch vs a direct method call in a benchmark.

### 35.3 Memory Layout and Cache Efficiency

1. Compare two struct layouts: one with fields ordered `bool, int64, bool, int64` and another reordered to minimize padding. Print sizes with `unsafe.Sizeof`.
1. Benchmark struct-of-arrays vs array-of-structs for iterating over a million records and summing a single field.
1. Write a benchmark comparing iteration over a `[]int` (contiguous memory) vs a `[]*int` (pointer chasing) to demonstrate cache locality effects.

### 35.4 Profiling Workflows

1. Write an HTTP server with `net/http/pprof`, generate load, capture a 30-second CPU profile, and identify the hottest functions with `go tool pprof`.
1. Capture a heap profile under load, use `go tool pprof -alloc_space` to find functions with the most allocations, and optimize the top one.
1. Use `go tool trace` to capture an execution trace of a concurrent program. Identify goroutine scheduling delays and blocking events.
1. Write a benchmark with `b.ReportAllocs()` and iteratively optimize until allocations reach zero. Document each optimization step.

### 35.5 Singleflight and Caching

1. Implement a cache with `singleflight` that prevents thundering herd on cache miss: only one goroutine fetches the value, others wait for the result.
1. Implement a two-level cache (in-memory L1 + Redis L2) with `singleflight` on both levels. Benchmark cache hit rates under concurrent access.
1. Implement cache stampede prevention using probabilistic early expiration: refresh the cache before TTL expires with a probability that increases as expiration approaches.

### 35.6 Concurrency Performance

1. Benchmark a shared counter implemented with `sync.Mutex` vs `atomic.AddInt64` vs per-goroutine counters with final aggregation. Compare throughput.
1. Implement a sharded map (N maps with N mutexes, key hashed to a shard) and benchmark it against `sync.Map` and a single `map` with `sync.RWMutex` under mixed read/write load.
1. Write a pipeline that processes items through 3 stages. Benchmark different channel buffer sizes (0, 1, 10, 100) and find the optimal throughput.
1. Use `errgroup` with `SetLimit(N)` to process 1000 tasks with bounded concurrency. Benchmark different values of N to find the optimal parallelism.

## Block 36: Advanced API Design

### 36.1 API Versioning

1. Implement URL-based API versioning: serve `/api/v1/users` and `/api/v2/users` with different response formats from the same server.
1. Implement header-based versioning: read `Accept: application/vnd.myapi.v2+json` and route to the correct handler.
1. Write a backward-compatible API change: add a new field to the response without breaking existing clients. Write tests that verify both old and new clients work.

### 36.2 Cursor-Based Pagination

1. Implement cursor-based pagination for a `GET /items` endpoint: return `next_cursor` in the response, accept `cursor` as a query parameter.
1. Encode the cursor as a base64-encoded JSON object containing the last seen ID and sort field value.
1. Write tests verifying pagination correctness: no duplicates, no skipped items, stable ordering under concurrent inserts.

### 36.3 Advanced Rate Limiting

1. Implement a token bucket rate limiter using `golang.org/x/time/rate` with per-user limits identified by API key.
1. Implement a sliding window rate limiter using Redis sorted sets: `ZADD` with timestamp, `ZREMRANGEBYSCORE` to trim, `ZCARD` to count.
1. Add rate limit headers to responses: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`.

### 36.4 Structured API Errors

1. Implement RFC 7807 Problem Details responses: return errors as JSON with `type`, `title`, `status`, `detail`, and `instance` fields.
1. Write a central error handler middleware that converts domain errors (`NotFoundError`, `ValidationError`) into structured API error responses.
1. Write a validation error response that includes per-field errors: `{"errors": [{"field": "email", "message": "invalid format"}]}`.

### 36.5 GraphQL

1. Build a GraphQL API using `gqlgen`: define a schema with `Query` and `Mutation` for a `User` type, generate resolvers, and implement them.
1. Add a `posts` field to `User` that demonstrates the N+1 problem. Fix it using a dataloader.
1. Implement a GraphQL subscription that streams new messages using WebSocket.

### 36.6 OpenAPI and Code Generation

1. Write an OpenAPI 3.0 spec for a `User` API and generate server stubs using `oapi-codegen`.
1. Generate a Go client from the same OpenAPI spec and use it in an integration test.
1. Add custom validation to the generated server using middleware that validates request bodies against the OpenAPI schema.

## Block 37: Observability in Depth

### 37.1 Structured Logging

1. Set up `slog` (Go 1.21+) with JSON handler for production and text handler for development. Log at `Info`, `Warn`, `Error` levels with structured fields.
1. Write a middleware that adds a request ID to the logger via `slog.With` and stores it in the request context. All subsequent logs include the request ID.
1. Write a log enrichment middleware that extracts user ID, trace ID, and method from the request and adds them to every log line automatically.

### 37.2 Distributed Tracing

1. Instrument an HTTP server with OpenTelemetry: create a span for each handler, add attributes (method, path, status code), and record errors.
1. Write an HTTP client that propagates trace context via `W3C Traceparent` headers. Verify that spans from two services appear in the same trace.
1. Add database query tracing: wrap `database/sql` calls with spans that include the query text (sanitized) and duration.
1. Write a gRPC interceptor that creates spans for each RPC call and propagates trace context via metadata.

### 37.3 Metrics in Depth

1. Write a custom Prometheus collector that exposes application-specific metrics (e.g., queue depth, active connections) by implementing the `prometheus.Collector` interface.
1. Create a RED metrics middleware (Rate, Errors, Duration) for an HTTP server: `http_requests_total`, `http_request_duration_seconds`, `http_requests_errors_total` with method and path labels.
1. Implement a business metrics exporter: track domain events (orders created, payments processed) as Prometheus counters.
1. Write a Grafana dashboard JSON model (as code) that visualizes request rate, error rate, and latency percentiles.

### 37.4 SLOs and Alerting

1. Define SLOs for an API service: 99.9% availability, p99 latency < 500ms. Write Prometheus recording rules that compute the error budget.
1. Implement an error budget tracker in Go that reads metrics from Prometheus and calculates remaining error budget as a percentage.
1. Write a health aggregator that checks multiple dependencies (database, cache, external API) and returns a structured health response with per-dependency status.

## Block 38: Advanced Security

### 38.1 OAuth2 and OIDC

1. Implement an OAuth2 authorization code flow: redirect to provider, handle callback, exchange code for token, extract user info.
1. Write middleware that validates an OIDC ID token (JWT) by fetching the provider's JWKS and verifying the signature.
1. Implement token refresh: detect expired access tokens and use the refresh token to obtain a new one transparently.

### 38.2 Secret Management

1. Write a configuration loader that reads secrets from environment variables in development and from HashiCorp Vault in production.
1. Implement secret rotation: write a service that periodically fetches a new database password from Vault and updates the connection pool.
1. Write a function that redacts sensitive fields (`password`, `token`, `secret`) from log output by implementing a custom `slog.Handler`.

### 38.3 TLS and mTLS

1. Write an HTTPS server that loads a TLS certificate and key. Generate self-signed certs with `crypto/x509` for testing.
1. Implement mutual TLS: configure both server and client to present certificates and verify each other's identity.
1. Write a test that starts an HTTPS server with a test certificate using `httptest.NewTLSServer` and makes requests with a custom `http.Transport`.

### 38.4 Security Hardening

1. Write a middleware that sets security headers: `Content-Security-Policy`, `X-Content-Type-Options`, `Strict-Transport-Security`, `X-Frame-Options`.
1. Implement request body size limiting middleware that rejects payloads larger than a configurable maximum with `413 Payload Too Large`.
1. Write a middleware that validates `Content-Type` header and rejects requests with unexpected content types.
1. Implement IP allowlisting middleware that reads allowed CIDRs from configuration and rejects requests from unauthorized IPs.

## Block 39: Advanced Concurrency Patterns

### 39.1 errgroup

1. Use `golang.org/x/sync/errgroup` to fetch data from 3 URLs concurrently. If any request fails, cancel the others and return the first error.
1. Use `errgroup` with `SetLimit(5)` to process 100 files concurrently with bounded parallelism.
1. Write a pipeline using `errgroup` where each stage runs in its own goroutine: read lines from a file, parse JSON, validate, and write to database.

### 39.2 Advanced Channel Patterns

1. Implement a fan-out/fan-in pattern where a slow consumer is protected by a dropping channel: if the buffer is full, new messages are discarded.
1. Write an `OrDone` channel combinator: given a `done` channel and a data channel, return a channel that closes when either closes.
1. Implement a priority queue using two channels: high-priority items are always processed before low-priority ones using `select` with `default`.
1. Write a `Merge` function that takes N channels and returns a single channel that receives values from all of them.

### 39.3 Worker Pool Patterns

1. Implement a dynamic worker pool that scales up to `maxWorkers` under load and scales down after an idle timeout.
1. Write a worker pool with per-worker local state (e.g., a database connection) that is initialized on first use and cleaned up on shutdown.
1. Implement a work-stealing scheduler: multiple workers each have their own queue, and idle workers steal from busy workers' queues.

### 39.4 Synchronization Patterns

1. Implement a barrier that blocks N goroutines until all of them have reached the barrier point, then releases all of them simultaneously.
1. Write a concurrent-safe lazy initializer using `sync.Once` that handles initialization errors: if init fails, allow retrying on the next call.
1. Implement a read-write lock with writer priority: pending writers block new readers to prevent writer starvation.

## Block 40: Go Internals and Runtime

### 40.1 Compiler and Linker

1. Use `go build -gcflags="-m -m"` to see detailed escape analysis decisions. Identify why specific variables escape to the heap.
1. Use `-ldflags="-s -w"` to strip debug information from a binary and compare file sizes.
1. Use `go build -gcflags="-N -l"` to disable optimizations and inlining for debugging. Compare the generated assembly with and without optimizations.

### 40.2 Runtime Internals

1. Write a program that prints the goroutine stack size using `runtime.Stack` and observe how it grows with deep recursion.
1. Use `runtime.SetFinalizer` to attach a cleanup function to an object and observe when it runs relative to GC.
1. Write a program that uses `runtime.LockOSThread` and explain when this is necessary (e.g., graphics, CGo with thread-local state).

### 40.3 CGo

1. Write a Go program that calls a C function (`strlen`) using `import "C"` and `// #include <string.h>`.
1. Pass a Go string to a C function: convert it with `C.CString`, call the function, and free the memory with `C.free`.
1. Write a Go wrapper around a simple C library: define C functions in a `.c` file, include the header in Go, and call them.
1. Benchmark a pure Go implementation vs a CGo call for the same computation to measure the CGo overhead.

### 40.4 Go Plugins

1. Build a Go plugin (`.so`) with `go build -buildmode=plugin` that exports a `Greet(name string) string` function. Load and call it from a host program using `plugin.Open` and `plugin.Lookup`.
1. Define a plugin interface in a shared package. Write two plugins that implement it differently and load one based on a config flag.

## Block 41: Production Debugging

### 41.1 Delve Debugger

1. Install Delve and debug a running Go program: set breakpoints, inspect variables, and step through code with `dlv debug`.
1. Attach Delve to an already running process using `dlv attach <pid>`. Inspect goroutines and their stack traces.
1. Use Delve's conditional breakpoints: break only when a variable equals a specific value (e.g., `break main.go:42 if userId == 5`).
1. Start a Delve headless server (`dlv debug --headless --listen=:2345`) and connect to it from VS Code for remote debugging.

### 41.2 Core Dumps and Crash Analysis

1. Set `GOTRACEBACK=crash` and trigger a panic. Inspect the resulting core dump with `dlv core`.
1. Write a program that crashes with a nil pointer dereference. Analyze the core dump to identify the exact line and goroutine that caused it.
1. Use `GODEBUG=tracebackancestors=10` to include ancestor goroutine stacks in panic output. Launch nested goroutines and observe the full chain.

### 41.3 Goroutine Leak Detection

1. Write a program with a goroutine leak (goroutine blocked on a channel with no sender). Use `runtime.NumGoroutine()` to detect the leak.
1. Capture a goroutine profile via `net/http/pprof` (`/debug/pprof/goroutine?debug=2`) and identify leaked goroutines by their stack trace.
1. Write a test using `go.uber.org/goleak` that fails if any unexpected goroutines remain after the test completes.
1. Implement a goroutine leak detector helper: record `runtime.NumGoroutine()` before and after a test, fail if the count increased.

### 41.4 Scheduler and GC Diagnostics

1. Run a program with `GODEBUG=schedtrace=1000` and interpret the output: P count, runnable goroutines, idle threads.
1. Run a program with `GODEBUG=gctrace=1` and interpret the output: heap size before/after GC, pause time, GC percentage.
1. Write a program with excessive GC pressure (allocating many short-lived objects). Use `GODEBUG=gctrace=1` to observe frequent GC cycles, then optimize with `sync.Pool` and compare.
1. Use `runtime/metrics` (Go 1.16+) to programmatically read GC pause times and scheduling latencies without parsing debug output.

### 41.5 Memory Leak Detection

1. Write a program with a memory leak (e.g., appending to a global slice in a loop). Capture heap profiles at intervals and compare them with `go tool pprof -diff_base`.
1. Use `go tool pprof -inuse_space` and `-alloc_space` to distinguish between currently held memory and total allocations over time.
1. Write an HTTP handler that leaks `bytes.Buffer` objects by storing them in a package-level map. Detect the leak with repeated heap profiles under load.

## Block 42: Advanced Database Patterns

### 42.1 Connection Pool Tuning

1. Write a program that connects to PostgreSQL and configures `SetMaxOpenConns`, `SetMaxIdleConns`, and `SetConnMaxLifetime`. Print pool stats with `db.Stats()`.
1. Benchmark different pool sizes (5, 20, 50, 100) under concurrent load and find the optimal configuration by measuring p99 latency and error rate.
1. Demonstrate connection exhaustion: set `MaxOpenConns` to 2, launch 10 concurrent queries, and observe timeouts. Fix with appropriate pool sizing and `context.WithTimeout`.

### 42.2 Query Optimization from Go

1. Write a function that runs `EXPLAIN ANALYZE` on a query and prints the execution plan from Go code.
1. Implement a slow query logger: wrap `database/sql` calls with timing and log queries exceeding a configurable threshold.
1. Write a query builder that constructs parameterized SQL dynamically based on optional filter parameters (avoid `WHERE 1=1` anti-pattern).

### 42.3 Repository and Unit of Work Patterns

1. Implement a `Repository` interface with `FindByID`, `FindAll`, `Create`, `Update`, `Delete`. Write both a PostgreSQL and an in-memory implementation.
1. Implement a Unit of Work pattern: begin a transaction, perform multiple repository operations, and commit or rollback atomically.
1. Write a `TxManager` that wraps `sql.Tx` lifecycle and passes the transaction through context: `func (m *TxManager) Do(ctx context.Context, fn func(ctx context.Context) error) error`.

### 42.4 Read Replicas and Sharding

1. Write a database router that sends `SELECT` queries to a read replica and all other queries to the primary. Use a custom `DB` wrapper with `Primary()` and `Replica()` methods.
1. Implement key-based sharding: given a `userID`, hash it to determine which of N database shards to query. Write `ShardFor(userID int) *sql.DB`.
1. Write a `ScatterGather` function that queries all shards in parallel, merges results, and sorts them.

### 42.5 Zero-Downtime Database Migrations

1. Write an expand-contract migration: add a new column as nullable, backfill data, then make it non-nullable in a separate migration.
1. Write a migration that renames a column safely: add new column, dual-write in application code, backfill, drop old column.
1. Implement a migration locking mechanism using a database advisory lock (`pg_advisory_lock`) to prevent concurrent migration runs.

## Block 43: Custom Static Analysis

### 43.1 Go AST Basics

1. Write a program that parses a Go source file with `go/parser` and prints all function declarations with their names and parameter counts.
1. Walk the AST with `ast.Inspect` and find all calls to `fmt.Println`. Print the file, line number, and arguments for each call.
1. Write a program that parses a Go file and lists all struct types with their field names and tags.

### 43.2 Writing Custom Linters

1. Write a linter using `go/analysis` that reports an error when a function has more than 5 parameters.
1. Write a linter that detects `context.Background()` usage inside HTTP handlers (should use `r.Context()` instead).
1. Write a linter that checks all exported functions in a package have a doc comment.
1. Write a linter that detects `time.Sleep` calls in non-test code and suggests using `time.Ticker` or `context.WithTimeout` instead.

### 43.3 Code Transformation

1. Write a tool using `go/ast` and `go/printer` that adds `//nolint` comments to specific function calls.
1. Write a tool that reads a Go file, finds all string literals, and replaces them with constants defined at the top of the file.
1. Write a refactoring tool that renames a function across a package using `golang.org/x/tools/go/packages` and `golang.org/x/tools/refactor/rename`.

## Block 44: Multi-Tenancy

### 44.1 Tenant Context Propagation

1. Write a middleware that extracts a tenant ID from the `X-Tenant-ID` header and stores it in the request context.
1. Write a `TenantFromContext(ctx context.Context) string` helper and use it in service and repository layers.
1. Write a logging middleware that automatically includes the tenant ID in every log line using `slog.With`.

### 44.2 Tenant Isolation Strategies

1. Implement row-level tenant isolation: add a `tenant_id` column to all tables and automatically filter queries by the tenant from context.
1. Implement schema-per-tenant isolation: write a connection wrapper that sets `search_path` to the tenant's schema before each query.
1. Implement database-per-tenant isolation: maintain a `map[string]*sql.DB` of tenant connections and route queries based on context.

### 44.3 Per-Tenant Resource Limits

1. Implement per-tenant rate limiting: maintain a `map[string]*rate.Limiter` and apply the correct limiter based on tenant ID from context.
1. Write a middleware that enforces per-tenant storage quotas: check current usage before allowing file uploads.
1. Implement per-tenant connection pool limits: allocate a separate `sql.DB` per tenant with individually configured `MaxOpenConns`.

## Block 45: Service Mesh and Service Discovery

### 45.1 Service Discovery

1. Write a service that registers itself with Consul on startup and deregisters on shutdown using the Consul HTTP API.
1. Write a client that discovers service instances from Consul and load-balances requests across them using round-robin.
1. Implement health check integration: register a Consul health check that calls your service's `/healthz` endpoint.

### 45.2 etcd for Configuration and Discovery

1. Connect to etcd using `clientv3` and perform `Put`/`Get`/`Delete` operations for key-value configuration.
1. Watch a configuration key in etcd and dynamically update application behavior when the value changes.
1. Implement leader election using etcd leases: multiple instances compete for a key, only the leader runs periodic tasks.

### 45.3 Service Mesh Awareness

1. Write an HTTP server that detects it's running behind an Envoy sidecar by checking for `x-envoy-` headers and logs them.
1. Implement sidecar-aware health checks: separate liveness (app is running) from readiness (app can accept traffic, sidecar is connected).
1. Write a gRPC service that propagates tracing headers injected by a service mesh (Istio/Linkerd) through the request context.

## Block 46: Custom Protocol Design

### 46.1 Binary Protocols over TCP

1. Design a simple binary protocol with a header (4-byte magic, 4-byte length, 1-byte type) and a body. Write `Encode` and `Decode` functions using `encoding/binary`.
1. Write a TCP server and client that communicate using the binary protocol: client sends a request, server parses it and sends a response.
1. Implement length-prefixed framing: write a `ReadFrame` function that first reads 4 bytes for length, then reads that many bytes for the payload.

### 46.2 Custom Connection Management

1. Write a connection pool for TCP connections: `Get() net.Conn`, `Put(net.Conn)`, `Close()` with max size and idle timeout.
1. Implement connection health checking: before returning a connection from the pool, send a ping and replace dead connections.
1. Write a `net.Conn` wrapper that adds read/write deadlines, metrics (bytes sent/received), and logging.

### 46.3 Multiplexing

1. Implement stream multiplexing over a single TCP connection: multiple logical streams share one connection, each frame tagged with a stream ID.
1. Write a multiplexed RPC system: client sends requests with unique IDs, server processes them concurrently, client matches responses to requests by ID.

## Block 47: Data Pipelines and Stream Processing

### 47.1 ETL Patterns

1. Write an ETL pipeline that reads CSV rows from a file, transforms each row (parse dates, normalize strings), and writes the results to a database.
1. Implement batch processing: accumulate rows into batches of 1000 and insert them with a single `INSERT ... VALUES` statement using `strings.Builder`.
1. Add error handling to the pipeline: write failed rows to a dead-letter file with the error message, continue processing the rest.

### 47.2 Backpressure

1. Implement a pipeline with backpressure using buffered channels: a slow consumer causes the producer to block when the buffer is full.
1. Write a dropping pipeline: if the consumer is too slow, drop the oldest messages (ring buffer semantics) and log the drop count.
1. Implement a rate-controlled pipeline using `time.Ticker`: process at most N items per second regardless of input rate.

### 47.3 Windowed Aggregation

1. Implement a time-based sliding window: aggregate events over the last 60 seconds, emitting a summary every 10 seconds.
1. Implement a count-based window: accumulate N events, process the batch, then start the next window.
1. Write a tumbling window aggregator that groups events into non-overlapping 1-minute buckets and computes min, max, avg for each bucket.

### 47.4 Exactly-Once Processing

1. Implement at-least-once processing with idempotent operations: write a consumer that tracks processed message IDs in a `sync.Map`.
1. Implement transactional outbox for exactly-once: consume a message, update state, and write an outbox record in the same database transaction.

## Block 48: Advanced HTTP and Networking

### 48.1 HTTP/2

1. Write an HTTP/2 server using `net/http` with TLS (HTTP/2 is enabled automatically with TLS). Verify with `curl --http2`.
1. Implement server push: use `http.Pusher` to push a CSS file alongside an HTML response.
1. Write a client that uses HTTP/2 connection multiplexing: send 100 concurrent requests over a single TCP connection and measure latency.

### 48.2 Custom http.RoundTripper

1. Write a custom `http.RoundTripper` that adds an `Authorization` header to every outgoing request.
1. Write a `RetryRoundTripper` that retries failed requests with exponential backoff for 5xx responses and network errors.
1. Write a `TracingRoundTripper` that creates an OpenTelemetry span for each outgoing HTTP request with method, URL, and status code attributes.
1. Compose multiple `RoundTripper` wrappers: `logging -> tracing -> retry -> http.DefaultTransport`.

### 48.3 HTTP Transport Tuning

1. Configure `http.Transport` with custom `MaxIdleConns`, `MaxIdleConnsPerHost`, `IdleConnTimeout`, and `TLSHandshakeTimeout`. Benchmark the difference under load.
1. Write a benchmark comparing HTTP/1.1 (one connection per request) vs HTTP/2 (multiplexed) for 100 concurrent requests to the same server.
1. Implement connection warm-up: pre-establish connections to downstream services on startup before accepting traffic.

### 48.4 Server-Sent Events and Long Polling

1. Implement an SSE endpoint with reconnection support: send a `Last-Event-ID` header from the client and resume from the correct position.
1. Implement long polling: client sends a request, server holds it open until new data is available or a timeout expires, then responds.
1. Write a broadcast hub for SSE: multiple clients subscribe, the server sends events to all connected clients concurrently using a fan-out pattern.

## Block 49: Go Module Governance

### 49.1 Vendoring

1. Vendor all dependencies with `go mod vendor`. Build with `go build -mod=vendor` and verify the binary works without network access.
1. Add a CI step that checks vendor consistency: run `go mod vendor` and fail if the diff is non-empty.
1. Show how to patch a vendored dependency: modify a file in `vendor/`, build, and document why vendoring was necessary.

### 49.2 Private Modules and Proxies

1. Configure `GOPRIVATE` to bypass the public proxy for private modules: `GOPRIVATE=github.com/myorg/*`.
1. Set up `GOPROXY` with a fallback chain: `GOPROXY=https://proxy.myorg.com,https://proxy.golang.org,direct`.
1. Configure `GONOSUMCHECK` for private modules and explain the security trade-off.

### 49.3 Module Lifecycle

1. Use `//Deprecated:` comment in a package doc to deprecate a module version. Verify it shows in `go doc`.
1. Use `retract` directive in `go.mod` to retract a broken version. Publish a new version and verify `go list` warns about the retracted version.
1. Manage a monorepo with multiple modules: create `go.mod` files in `libs/auth` and `libs/logging`, reference them from the main module using `replace` directives.

### 49.4 Go Workspaces for Monorepos

1. Create a workspace with `go work init`. Add three modules: `services/api`, `services/worker`, and `libs/shared`. Build all with `go build ./...`.
1. Demonstrate that `go.work` replaces are local-only: verify that CI without `go.work` still builds using published module versions.

## Block 50: Dependency Injection Frameworks

### 50.1 Google Wire

1. Install Wire and define a simple injector: a `NewServer` function that depends on `NewHandler`, which depends on `NewService`, which depends on `NewRepository`.
1. Write a `wire.go` file with `wire.Build` sets. Run `wire ./...` to generate the `wire_gen.go` file. Inspect the generated code.
1. Use `wire.Bind` to bind an interface to a concrete implementation in the provider set.
1. Use `wire.Struct` to auto-inject all fields of a struct.

### 50.2 Uber Fx

1. Write an application using `fx.New` with `fx.Provide` for constructors and `fx.Invoke` for startup logic.
1. Use `fx.Lifecycle` to register `OnStart` and `OnStop` hooks for an HTTP server.
1. Use `fx.Supply` to inject configuration values and `fx.Annotate` with `name` tags to distinguish multiple instances of the same type.
1. Write a test using `fxtest.New` that replaces a real database with a mock by overriding a provider.

### 50.3 Manual vs Framework DI Comparison

1. Implement the same application three ways: manual constructor injection, Google Wire, and Uber Fx. Compare boilerplate, error messages, and testability.

## Block 51: Chaos Engineering

### 51.1 Fault Injection

1. Write a middleware that randomly returns `500 Internal Server Error` for a configurable percentage of requests to simulate downstream failures.
1. Implement a `ChaosTransport` (`http.RoundTripper`) that randomly adds latency (100ms-2s) or returns errors on outgoing HTTP calls.
1. Write a test that injects a network partition (close a database connection mid-transaction) and verifies the application handles it gracefully.

### 51.2 Toxiproxy

1. Write an integration test that uses `toxiproxy-go` to add 500ms latency to a database connection and verifies the application returns a timeout error.
1. Simulate a connection reset using Toxiproxy's `reset_peer` toxic and verify the application retries and recovers.
1. Simulate bandwidth limitation using Toxiproxy's `bandwidth` toxic and measure the impact on request throughput.

### 51.3 Resilience Validation

1. Write a test suite that validates all resilience patterns from Block 31: inject failures and verify that retries, circuit breakers, and bulkheads behave correctly.
1. Implement a chaos monkey goroutine that randomly kills and restarts worker goroutines in a worker pool and verifies the pool self-heals.
1. Write a load test combined with fault injection: run `vegeta` while randomly killing downstream services and verify that error rate stays within SLO bounds.

## Block 52: WebAssembly

### 52.1 Go to WASM

1. Compile a Go program to WebAssembly using `GOOS=js GOARCH=wasm go build -o main.wasm`. Serve it with the provided `wasm_exec.js` and load it in a browser.
1. Write a Go function exported to JavaScript via `js.Global().Set` that accepts arguments from JS and returns a result.
1. Call JavaScript functions from Go using `js.Global().Get("document").Call("getElementById", "myDiv")`.

### 52.2 WASI

1. Compile a Go program targeting WASI with `GOOS=wasip1 GOARCH=wasm`. Run it with `wasmtime` or `wazero`.
1. Write a Go program compiled to WASI that reads from stdin and writes to stdout, demonstrating WASI I/O capabilities.

### 52.3 WASM Plugins

1. Write a host application using `wazero` that loads a WASM module compiled from Go and calls an exported function.
1. Implement a plugin system: define a plugin interface, compile plugins to WASM, and load them dynamically at runtime using `wazero`.

## Block 53: Profile-Guided Optimization

### 53.1 PGO Basics

1. Build a Go HTTP server without PGO. Generate load with `hey` or `vegeta` while capturing a CPU profile to `default.pgo`.
1. Rebuild the binary with PGO: place `default.pgo` in the main package directory and run `go build`. Compare binary sizes.
1. Benchmark the PGO-optimized binary against the non-PGO binary under the same load. Measure and compare p50, p95, p99 latencies.

### 53.2 PGO Workflow

1. Implement a PGO workflow in CI: download the production CPU profile, build with PGO, run benchmarks, and deploy if performance improves.
1. Use `go build -pgo=path/to/profile.pgo` to explicitly specify the profile path instead of relying on auto-detection.
1. Write a script that merges multiple CPU profiles from different production instances into a single representative profile using `go tool pprof -proto`.

## Block 54: Tech Lead Practices

### 54.1 Architecture Decision Records

1. Write an ADR template with sections: Title, Status, Context, Decision, Consequences. Create an ADR for choosing PostgreSQL over MongoDB for a new service.
1. Create an `docs/adr/` directory structure with a numbered ADR format (`0001-use-postgresql.md`). Write a Makefile target `adr-new` that generates the next numbered file from a template.
1. Write an ADR for choosing between `chi` and `gin` as the HTTP router, documenting trade-offs and the final decision.

### 54.2 RFC Process

1. Write an RFC template with sections: Summary, Motivation, Detailed Design, Drawbacks, Alternatives, Unresolved Questions.
1. Write an RFC proposing a migration from monolith to microservices: describe the extraction strategy, communication patterns, and rollback plan.

### 54.3 Code Review Guidelines

1. Write a code review checklist for Go projects covering: error handling, concurrency safety, test coverage, naming conventions, and performance considerations.
1. Write a `CONTRIBUTING.md` for a Go project with sections on code style, commit message format, PR process, and required checks.
1. Create a `.golangci.yml` configuration with custom rules enforcing project-specific conventions (e.g., no `panic` in library code, mandatory error wrapping).

### 54.4 Technical Debt Management

1. Write a Go program that scans a codebase for `TODO`, `FIXME`, and `HACK` comments and generates a technical debt report with file, line, author (from `git blame`), and age.
1. Define a technical debt scoring system: categorize each item by severity (critical, high, medium, low) and effort (hours). Write a Go struct and markdown report generator.

### 54.5 API Design Review

1. Write an API design review checklist covering: naming consistency, HTTP method correctness, error response format, pagination, versioning, and idempotency.
1. Review a provided API spec and identify issues: missing error codes, inconsistent naming, lack of pagination, missing rate limit headers. Write a report with recommendations.

## Block 55: Event-Driven Architecture

### 55.1 Event Schema Evolution

1. Define a Protobuf message `OrderCreated` with 5 fields. Add a new optional field in a backward-compatible way. Write a test that deserializes old messages with the new schema.
1. Remove a field from a Protobuf message by reserving its field number with `reserved`. Write a test proving that old serialized data still deserializes without error.
1. Implement a schema registry client: register a Protobuf schema by subject, retrieve it by ID, and check compatibility before producing messages.
1. Write a consumer that handles multiple versions of an event: detect the schema version from a header and deserialize accordingly.

### 55.2 Change Data Capture (CDC)

1. Write a program that reads PostgreSQL WAL changes using logical replication (`pglogrepl` or `pgx` logical replication support) and prints INSERT/UPDATE/DELETE events.
1. Implement a CDC pipeline: capture row changes from PostgreSQL, transform them into domain events, and publish to a Kafka topic.
1. Write a CDC consumer that maintains a read-optimized materialized view in Redis by applying database change events in order.
1. Handle CDC ordering guarantees: partition events by primary key so that updates to the same row are always processed in order.

### 55.3 Dead Letter Queues

1. Implement a Kafka consumer that retries failed messages 3 times with backoff, then publishes them to a dead letter topic with the original message, error, and retry count in headers.
1. Write a dead letter queue processor that reads messages from the DLQ, displays them in a CLI tool, and allows manual replay to the original topic.
1. Implement poison pill detection: if the same message fails N times across consumer restarts, move it to the DLQ and continue processing the partition.

### 55.4 Ordering and Partitioning

1. Write a Kafka producer that uses a consistent partition key (e.g., `userID`) to guarantee ordering of events for the same entity.
1. Implement an event resequencer: consume out-of-order events (each carrying a sequence number), buffer them, and emit in order once gaps are filled.
1. Write a test that verifies ordering guarantees: produce 100 events for the same key and assert the consumer receives them in the exact same order.

## Block 56: Advanced Caching

### 56.1 Cache Invalidation Strategies

1. Implement write-through caching: on every database write, update the cache synchronously before returning to the caller.
1. Implement write-behind (write-back) caching: buffer writes in the cache and flush them to the database asynchronously in batches using a background goroutine.
1. Implement event-based cache invalidation: listen to a Kafka topic of data change events and delete or update the corresponding cache entries.
1. Write a test that verifies cache consistency: perform a write, then immediately read from a different instance, and assert the cache is up to date.

### 56.2 Distributed Cache Patterns

1. Implement a two-level cache: L1 is an in-memory LRU cache (`map` + doubly linked list), L2 is Redis. Check L1 first, then L2, then the database.
1. Write a cache-aside pattern with `singleflight`: on cache miss, only one goroutine fetches from the database, populates both L1 and L2, and returns the result to all waiters.
1. Implement cache warming on startup: preload the top N most frequently accessed keys from the database into the cache before the service marks itself as ready.
1. Implement probabilistic early expiration: refresh a cache entry before its TTL expires with a probability that increases as the expiration approaches. Benchmark thundering herd reduction.

### 56.3 LRU Cache Implementation

1. Implement an LRU cache with `Get(key) (value, bool)` and `Put(key, value)` using a `map` and a doubly linked list. Ensure O(1) for both operations.
1. Make the LRU cache thread-safe using `sync.RWMutex`. Benchmark read-heavy workloads with `RWMutex` vs `sync.Mutex`.
1. Add TTL support to the LRU cache: entries expire after a configurable duration. Use a lazy expiration strategy (check on access) and a background goroutine for periodic cleanup.
1. Implement cache size limiting by memory (not just entry count): track approximate memory usage using `unsafe.Sizeof` and evict entries when the limit is reached.

## Block 57: Async Job Processing

### 57.1 Background Jobs with Asynq

1. Set up `asynq` with Redis: define a task type `email:send` with a JSON payload, enqueue it from an HTTP handler, and process it in a worker.
1. Configure task retries with exponential backoff: set `MaxRetry: 5` and a custom `RetryDelayFunc`. Write a test that verifies retry behavior.
1. Implement task priority queues: define `critical`, `default`, and `low` queues with different processing weights.
1. Add a unique task constraint: ensure that only one `report:generate` task per `userID` is enqueued at a time using `asynq.Unique`.

### 57.2 Scheduled and Delayed Jobs

1. Implement a delayed job: enqueue a `reminder:send` task that executes 24 hours after creation using `asynq.ProcessIn(24 * time.Hour)`.
1. Implement a cron-like scheduler: register a periodic task that runs every 5 minutes using `asynq.NewScheduler` with a cron expression.
1. Write a distributed cron job: ensure that only one instance of a periodic task runs across multiple worker nodes using Redis-based locking.

### 57.3 Job Processing Patterns

1. Implement a job chain: `CreateOrder -> SendConfirmationEmail -> UpdateAnalytics`. Each job enqueues the next one on success.
1. Implement a fan-out job: a single `batch:process` job spawns N child jobs (one per item) and tracks their completion in Redis.
1. Write a dead letter handler for failed jobs: after all retries are exhausted, write the task payload and error to a PostgreSQL `failed_jobs` table for manual inspection.
1. Implement graceful shutdown for a worker: on `SIGTERM`, stop accepting new jobs, wait for in-progress jobs to finish (up to 30 seconds), then exit.

## Block 58: Advanced Database Patterns

### 58.1 Locking Strategies

1. Implement optimistic locking: add a `version` column to a table, read the version with the row, and use `UPDATE ... WHERE version = $1` to detect conflicts. Return a `ErrConflict` on mismatch.
1. Implement pessimistic locking: use `SELECT ... FOR UPDATE` to lock a row during a transaction. Write a test with two concurrent goroutines to verify mutual exclusion.
1. Implement a distributed lock using PostgreSQL advisory locks: `pg_advisory_lock(key)` and `pg_advisory_unlock(key)`. Write a wrapper `func WithAdvisoryLock(ctx context.Context, db *sql.DB, key int64, fn func() error) error`.
1. Compare optimistic vs pessimistic locking under high contention: write a benchmark where 50 goroutines update the same row and measure throughput and conflict rate for each strategy.

### 58.2 Partitioning and Materialized Views

1. Write SQL migrations to create a range-partitioned `events` table by month in PostgreSQL. Write a Go function that inserts events and queries across partitions transparently.
1. Implement automatic partition management: write a Go function that creates future partitions (next 3 months) and drops old partitions (older than 12 months).
1. Create a materialized view for an analytics dashboard: `CREATE MATERIALIZED VIEW daily_stats AS ...`. Write a Go function that refreshes it concurrently with `REFRESH MATERIALIZED VIEW CONCURRENTLY`.
1. Implement an incremental materialized view in Go: maintain an in-memory aggregate that updates on each INSERT event instead of recomputing from scratch.

### 58.3 Audit Logging

1. Write a PostgreSQL trigger (via migration) that logs every INSERT, UPDATE, DELETE on a `users` table to an `audit_log` table with `operation`, `old_data`, `new_data`, `changed_by`, and `changed_at`.
1. Implement application-level audit logging: write a repository wrapper that records every mutation to an `audit_events` table in the same transaction as the business operation.
1. Write a query function that reconstructs the state of a row at a given point in time by replaying audit events from the creation event up to the target timestamp.

## Block 59: Modern Go Standard Library

### 59.1 log/slog Package (Go 1.21+)

1. Set up `slog` with a `JSONHandler` for production and a `TextHandler` for development. Log at `Info`, `Warn`, `Error` levels with structured key-value fields.
1. Write a middleware that creates a child logger with `slog.With("request_id", id)`, stores it in the request context, and retrieves it in handlers with a `LoggerFromContext` helper.
1. Implement a custom `slog.Handler` that redacts sensitive fields: replace values for keys matching `password`, `token`, `secret`, `authorization` with `"[REDACTED]"`.
1. Write a `slog.Handler` that adds log level colorization for terminal output while keeping JSON format for production.

### 59.2 maps and slices Packages (Go 1.21+)

1. Use `slices.SortFunc` to sort a `[]Person` by age, then by name for equal ages. Compare verbosity with the old `sort.Slice` approach.
1. Use `slices.Contains`, `slices.Index`, and `slices.Compact` to process a slice of strings: check membership, find position, and remove consecutive duplicates.
1. Use `maps.Keys`, `maps.Values`, and `maps.Equal` to work with maps. Write a function that returns the symmetric difference of two maps using these helpers.
1. Use `slices.BinarySearch` on a sorted slice and compare it with `sort.Search`.

### 59.3 iter Package (Go 1.23+)

1. Write a custom iterator function `func All[T any](s []T) iter.Seq2[int, T]` and use it with `for i, v := range All(mySlice)`.
1. Write an iterator `Lines(r io.Reader) iter.Seq[string]` that yields lines from a reader one at a time.
1. Write a `Filter[T any](seq iter.Seq[T], fn func(T) bool) iter.Seq[T]` combinator that chains with other iterators.
1. Write a `Map[T, U any](seq iter.Seq[T], fn func(T) U) iter.Seq[U]` combinator. Chain `Filter` and `Map` to process a sequence in a pipeline.

### 59.4 io/fs and testing/fstest (Go 1.16+)

1. Write a function that accepts `fs.FS` instead of a file path. Test it with `fstest.MapFS` containing in-memory files, no real disk I/O needed.
1. Refactor an `embed.FS` usage to accept any `fs.FS`, making it testable with `fstest.MapFS` and swappable with `os.DirFS` at runtime.
1. Write a `WalkTemplates(fsys fs.FS) ([]string, error)` function that finds all `.tmpl` files. Test it with `fstest.MapFS` containing a nested directory structure.
1. Implement an `fs.FS` backed by an S3-compatible client: `Open` fetches the object, `ReadDir` lists the prefix. Use `fstest.TestFS` to validate the implementation.

### 59.5 net/netip Package (Go 1.18+)

1. Parse an IPv4 and IPv6 address with `netip.ParseAddr`. Print whether each is IPv4, IPv6, loopback, or link-local.
1. Parse a CIDR prefix with `netip.ParsePrefix` and write a function `Contains(prefix netip.Prefix, addr netip.Addr) bool` that checks if an address belongs to the prefix.
1. Implement an IP allowlist using `[]netip.Prefix` and write middleware that rejects requests from addresses not in any allowed prefix.
1. Benchmark `netip.Addr` comparisons vs `net.IP` comparisons in a hot path to demonstrate the performance advantage of the value-type design.

### 59.6 cmp Package (Go 1.21+)

1. Use `cmp.Compare` to write a multi-field comparator: sort `[]Product` by price ascending, then by name ascending for equal prices using `slices.SortFunc`.
1. Use `cmp.Or` to implement a fallback chain: return the first non-zero value from a list of configuration sources (flag, env, config file, default).

## Block 60: Advanced HTTP Server Tuning

### 60.1 Server Timeout Configuration

1. Configure an `http.Server` with `ReadTimeout`, `WriteTimeout`, `ReadHeaderTimeout`, and `IdleTimeout`. Write a test that verifies a slow client triggers a timeout.
1. Implement per-route timeouts using `http.TimeoutHandler`: set a 2-second timeout on a slow endpoint and verify it returns `503 Service Unavailable`.
1. Write a middleware that sets `MaxBytesReader` on the request body to limit payload size. Return `413 Payload Too Large` for oversized requests.
1. Configure `MaxHeaderBytes` on the server and write a test that sends a request with an oversized header and verifies it's rejected.

### 60.2 Connection Management

1. Write a load test that opens 1000 concurrent connections. Tune `http.Server` with `ConnState` callback to log connection state transitions (new, active, idle, closed).
1. Implement connection draining: on `SIGTERM`, stop accepting new connections, set a deadline for existing ones, and log connections that are forcefully closed.
1. Write a connection limiter middleware using a buffered channel semaphore: reject new connections with `503` when the limit is reached.

### 60.3 Custom http.RoundTripper Composition

1. Write a `LoggingRoundTripper` that logs method, URL, status code, and duration for every outgoing request.
1. Write a `CachingRoundTripper` that caches GET responses in memory with `Cache-Control` header respect. Return cached responses for repeated requests.
1. Compose a transport chain: `Logging -> Tracing -> Retry -> Caching -> http.DefaultTransport`. Write a test that verifies each layer is called in order.

## Block 61: Graceful Degradation and Resilience

### 61.1 Fallback Strategies

1. Implement a service that calls an external API with a fallback: if the API returns an error or times out, return a cached response from Redis with a `X-Fallback: true` header.
1. Write a handler that returns a partial response: if one of three downstream calls fails, return the successful results with a `warnings` field indicating the degraded data.
1. Implement a circuit breaker with a static fallback: when the circuit is open, return a default response instead of an error.
1. Write a health-aware load balancer: track success rates per backend and stop sending traffic to unhealthy backends. Periodically probe them and restore traffic when they recover.

### 61.2 Request Hedging

1. Implement request hedging: send a request to two backends simultaneously, return the first successful response, and cancel the other using `context.WithCancel`.
1. Implement delayed hedging: send a request to backend A, if no response in 100ms send the same request to backend B, return whichever responds first.
1. Write a benchmark comparing hedged requests vs single requests under simulated variable latency. Measure p50, p95, p99 improvements.

### 61.3 Backpressure at the HTTP Layer

1. Implement a middleware that tracks in-flight requests using `atomic.AddInt64`. When the count exceeds a threshold, return `429 Too Many Requests` with a `Retry-After` header.
1. Implement adaptive load shedding: measure p99 latency with an exponentially weighted moving average. Start rejecting a percentage of requests when latency exceeds a target.
1. Write a priority-based admission control middleware: requests with an `X-Priority: high` header bypass load shedding while normal requests are subject to it.

## Block 62: API Gateway and BFF

### 62.1 API Composition

1. Write an API gateway handler that aggregates data from three microservices (`UserService`, `OrderService`, `ProductService`) into a single response using `errgroup` for concurrent fetching.
1. Implement request fan-out with partial failure handling: if one service is down, return the available data with a `degraded: true` flag and a list of failed services.
1. Write a response transformer: the gateway accepts a `fields` query parameter and strips unwanted fields from the aggregated response before returning it to the client.

### 62.2 Backend for Frontend (BFF)

1. Build a BFF for a mobile client: aggregate user profile, recent orders, and notifications into a single `GET /mobile/home` endpoint optimized for the mobile screen.
1. Build a BFF for a web dashboard: aggregate analytics summary, recent activity, and alerts into a single `GET /web/dashboard` endpoint with different field selection than the mobile BFF.
1. Write a BFF middleware that transforms snake_case API responses to camelCase for JavaScript clients.

### 62.3 Gateway Middleware

1. Implement a gateway-level rate limiter that enforces per-API-key limits across all downstream services using Redis-backed token buckets.
1. Write a request routing middleware that reads an `X-Api-Version` header and proxies the request to the corresponding backend version.
1. Implement request/response logging at the gateway level: log sanitized request bodies (redact sensitive fields) and response status codes with latency.

## Block 63: Distributed Transactions

### 63.1 Two-Phase Commit (2PC)

1. Implement a simplified 2PC coordinator: send `Prepare` to two participants (simulated as goroutines), collect votes, then send `Commit` or `Abort` based on unanimity.
1. Handle participant failure: if a participant times out during the prepare phase, the coordinator sends `Abort` to all participants.
1. Implement a participant that writes a prepare log before voting `Yes` and uses it to recover after a crash (simulate crash with `panic` + `recover`).

### 63.2 TCC Pattern (Try-Confirm-Cancel)

1. Implement a TCC flow for a transfer: `Try` reserves funds in both accounts, `Confirm` commits the transfer, `Cancel` releases the reservations.
1. Write a TCC coordinator that calls `Try` on all participants, then `Confirm` if all succeed or `Cancel` if any fail. Handle partial `Cancel` failures with retries.
1. Add timeout-based auto-cancellation: if `Confirm` is not called within 5 minutes after `Try`, a background goroutine automatically calls `Cancel`.

### 63.3 Outbox Pattern in Depth

1. Implement the transactional outbox pattern: write a domain event to an `outbox` table in the same database transaction as the business data.
1. Write an outbox poller that reads unpublished events every second using `SELECT ... FOR UPDATE SKIP LOCKED`, publishes them to Kafka, and marks them as published.
1. Implement outbox cleanup: write a background job that deletes published outbox records older than 7 days.
1. Write a test that verifies atomicity: if the business operation fails, no outbox record is created; if it succeeds, the outbox record is always present.

## Block 64: Advanced Observability

### 64.1 Continuous Profiling

1. Integrate `pyroscope-go` into an HTTP server for continuous CPU and memory profiling. Tag profiles with service name and version labels.
1. Write a middleware that attaches request path and tenant ID as profiling labels so that profiles can be filtered by endpoint or tenant.
1. Implement on-demand profiling: expose an endpoint `POST /debug/profile` that captures a 30-second CPU profile and returns it as a downloadable file.

### 64.2 OpenTelemetry Collector

1. Configure an OpenTelemetry Collector pipeline: receive OTLP traces from a Go service, batch them, and export to Jaeger.
1. Add a metrics pipeline to the collector: receive OTLP metrics from a Go service and export to Prometheus.
1. Implement tail-based sampling in the collector: only export traces that contain an error span or exceed 1 second duration.

### 64.3 Exemplars and Metrics-to-Traces Correlation

1. Attach trace ID as an exemplar to a Prometheus histogram observation using `prometheus.ExemplarObserver`. Verify the exemplar appears in Prometheus.
1. Write a middleware that records HTTP request duration as a histogram with an exemplar containing the trace ID, enabling one-click navigation from a Grafana metric to the corresponding trace.

### 64.4 SLO Implementation

1. Define SLOs for an API service (99.9% availability, p99 < 500ms). Write Prometheus recording rules that compute error budget consumption over a 30-day rolling window.
1. Implement an error budget tracker in Go: query Prometheus, calculate remaining error budget as a percentage, and expose it as a Prometheus gauge.
1. Write a multi-dependency health aggregator: check database, cache, and message broker. Return structured JSON with per-dependency status (`ok`, `degraded`, `down`) and overall health.

## Block 65: Data Formats and Serialization

### 65.1 Protocol Buffers Advanced

1. Use Protobuf `oneof` to model a `PaymentMethod` that is either `CreditCard`, `BankTransfer`, or `Wallet`. Marshal and unmarshal in Go.
1. Use Protobuf `map` field type to model a `Metadata map<string, string>` in a message. Write tests for serialization and deserialization.
1. Use well-known types: `google.protobuf.Timestamp`, `google.protobuf.Duration`, `google.protobuf.Struct`. Convert between Go `time.Time` and Protobuf `Timestamp`.
1. Implement a Protobuf-to-JSON gateway: accept a Protobuf-encoded request body, unmarshal it, process it, and return a JSON response using `protojson.Marshal`.

### 65.2 Binary Formats

1. Serialize and deserialize a struct using MessagePack (`vmihailenco/msgpack`). Benchmark against `encoding/json` for a nested struct with 20 fields.
1. Serialize and deserialize a struct using CBOR (`fxamacker/cbor`). Compare payload sizes with JSON and MessagePack for the same data.
1. Write a benchmark comparing serialization formats (JSON, Protobuf, MessagePack, CBOR) for throughput (ops/sec), payload size, and allocation count.

### 65.3 Configuration Formats

1. Parse a YAML configuration file into a Go struct using `gopkg.in/yaml.v3`. Handle nested structs, slices, and maps.
1. Parse a TOML configuration file into a Go struct using `pelletier/go-toml`. Compare with YAML for readability and type safety.
1. Write a unified configuration loader that reads the same config from JSON, YAML, or TOML based on file extension, returning the same Go struct.

### 65.4 CSV Processing

1. Read a CSV file using `encoding/csv`, parse each row into a struct, and print invalid rows with line numbers and errors.
1. Write a CSV export function: given a `[]User`, write a CSV file with headers derived from struct field names using `reflect`.
1. Implement streaming CSV processing: read a 1GB CSV file line-by-line without loading it entirely into memory, transform each row, and write to an output CSV.

## Block 66: Incident Management and On-Call

### 66.1 Runbooks

1. Write a runbook template with sections: Alert Name, Severity, Symptoms, Investigation Steps, Mitigation Steps, Root Cause Checklist, Escalation Path.
1. Write a runbook for "High Error Rate on /api/orders": check recent deployments, inspect logs for new error patterns, check downstream service health, check database connection pool stats.
1. Write a runbook for "Database Connection Pool Exhausted": check `db.Stats()` output, identify long-running queries with `pg_stat_activity`, kill idle-in-transaction connections, increase pool size.

### 66.2 Incident Response

1. Write a Go CLI tool that creates an incident: generates a timestamped incident document from a template, creates a Slack channel name suggestion, and outputs initial status update text.
1. Implement a status page endpoint `GET /status` that returns the current operational status of all service components with severity levels (`operational`, `degraded`, `outage`).
1. Write a postmortem template generator: given incident start/end times, it queries Prometheus for error rates and latency during the window and generates a timeline section.

### 66.3 Alerting

1. Write Prometheus alerting rules for: high error rate (>1% of 5xx in 5 minutes), high latency (p99 > 1s for 5 minutes), and low availability (less than 2 healthy pods).
1. Implement an alert webhook handler in Go: receive PagerDuty/Alertmanager webhooks, parse the payload, and execute automated remediation (e.g., restart a pod, clear a cache).
1. Write an alert fatigue analyzer: query Alertmanager API for alert history, group by alert name, and report alerts that fire more than 10 times per week as candidates for tuning.

## Block 67: Capacity Planning and Performance Budgets

### 67.1 Capacity Estimation

1. Write a capacity calculator in Go: given requests per second, average response time, and target CPU utilization (70%), calculate the number of pods/instances needed.
1. Implement a load model: given a traffic pattern (peak hours, growth rate), project the required infrastructure for 3, 6, and 12 months using exponential growth formula.
1. Write a benchmark suite that determines the maximum RPS a single instance can handle at p99 < 200ms. Use binary search over RPS values with `vegeta` or custom load generation.

### 67.2 Performance Budgets

1. Write a CI check that runs benchmarks and fails if any benchmark regresses by more than 10% compared to the baseline stored in a file.
1. Implement a latency budget tracker: given a total budget of 500ms for an endpoint, track time spent in each layer (handler: 10ms, service: 50ms, database: 200ms, cache: 5ms) and log the breakdown.
1. Write a `LatencyBudget` middleware that sets a deadline on the context based on the remaining budget and passes it to downstream calls.

## Block 68: Zero-Downtime Deployments

### 68.1 Database Migration Strategies for Large Tables

1. Implement an online column addition: add a nullable column, backfill in batches of 1000 rows with `UPDATE ... WHERE id BETWEEN $1 AND $2`, then add a `NOT NULL` constraint.
1. Write a batch backfill function with rate limiting: update N rows per second to avoid overwhelming the database. Use `time.Ticker` to control the pace.
1. Implement a dual-write migration: write to both old and new columns during the transition period. Write a verification query that checks consistency between the two columns.
1. Write a safe column rename migration: add new column, start dual-writing in application code, backfill old rows in batches, switch reads to new column, drop old column. Implement each step as a separate migration.

### 68.2 Feature Toggles

1. Implement feature toggles backed by a database table: `feature_flags(name, enabled, rollout_percentage, created_at)`. Write `IsEnabled(ctx, flag) bool` that checks user hash against rollout percentage.
1. Implement a feature toggle that targets specific users: store a list of user IDs per flag and enable the feature only for those users plus the percentage rollout.
1. Write middleware that injects enabled feature flags into the response headers (`X-Features: new-checkout,dark-mode`) for client-side feature detection.
1. Implement feature flag caching: cache flags in memory for 30 seconds, refresh from database in the background, and serve stale data during refresh failures.

### 68.3 Canary Deployments

1. Write a traffic splitter middleware: route 5% of requests to the canary instance based on a consistent hash of the user ID.
1. Implement canary health monitoring: compare error rates and latencies between canary and stable instances. If the canary error rate is 2x the stable rate, log an alert.
1. Write a gradual rollout controller: increase canary traffic from 5% to 10%, 25%, 50%, 100% with a configurable interval and automatic rollback if error rate exceeds a threshold.

## Block 69: Service Discovery and DNS

### 69.1 DNS-Based Service Discovery

1. Implement DNS-based service discovery: resolve SRV records with `net.LookupSRV("", "", "myservice.service.consul")` and use the returned addresses for load balancing.
1. Write an HTTP client that re-resolves DNS every 30 seconds to pick up new service instances. Store resolved addresses in an `atomic.Value` for lock-free reads.
1. Implement DNS caching with TTL respect: cache resolved addresses for the duration specified in the DNS response, then re-resolve.

### 69.2 Client-Side Load Balancing

1. Implement round-robin load balancing: given a list of backend addresses, distribute requests evenly using an `atomic.Int64` counter.
1. Implement weighted round-robin: assign weights to backends (e.g., based on instance CPU capacity) and distribute requests proportionally.
1. Implement least-connections load balancing: track active request count per backend with `atomic.Int64` and send new requests to the backend with the fewest active connections.
1. Implement health-aware load balancing: periodically health-check backends and remove unhealthy ones from the rotation. Restore them when they recover.

### 69.3 gRPC Load Balancing

1. Implement a custom gRPC `resolver.Builder` that resolves service addresses from a configuration file and notifies the balancer of changes.
1. Use gRPC's built-in `round_robin` balancer with a custom resolver. Write a test with 3 server instances verifying even distribution.
1. Implement a gRPC health-checking resolver: use the gRPC health protocol (`grpc.health.v1.Health`) to remove unhealthy servers from the address list.

## Block 70: Conflict Resolution and CRDTs

### 70.1 Optimistic Concurrency in Distributed Systems

1. Implement a vector clock: each node maintains a counter, increment on local events, merge on receive. Write `Compare(a, b VectorClock) (Before | After | Concurrent)`.
1. Implement last-write-wins (LWW) conflict resolution: given two concurrent updates with timestamps, keep the one with the higher timestamp. Write tests for clock skew scenarios.
1. Implement a version vector for a distributed key-value store: detect conflicting writes and return both versions to the client for application-level resolution.

### 70.2 CRDT Implementations

1. Implement a G-Counter (grow-only counter) CRDT: each node maintains its own counter, the global value is the sum. Write `Increment(nodeID)`, `Value() int`, and `Merge(other GCounter)`.
1. Implement a PN-Counter (positive-negative counter): two G-Counters (one for increments, one for decrements). Support `Increment`, `Decrement`, `Value`, and `Merge`.
1. Implement an OR-Set (observed-remove set) CRDT: support `Add(element)`, `Remove(element)`, `Contains(element) bool`, and `Merge(other ORSet)`. Use unique tags to handle add/remove conflicts.
1. Write a property-based test for each CRDT: verify commutativity (merge order doesn't matter), associativity, and idempotency of the `Merge` operation.

## Block 71: HTTP Server Hardening

### 71.1 Security Headers Middleware

1. Write a middleware that sets security headers on every response: `Content-Security-Policy`, `X-Content-Type-Options: nosniff`, `Strict-Transport-Security`, `X-Frame-Options: DENY`, `Referrer-Policy`.
1. Write a middleware that validates the `Content-Type` header on POST/PUT/PATCH requests and rejects requests with unexpected content types with `415 Unsupported Media Type`.
1. Implement CORS middleware from scratch: handle preflight `OPTIONS` requests, validate `Origin` against an allowlist, and set appropriate `Access-Control-*` headers.

### 71.2 Input Validation and Sanitization

1. Write a middleware that limits the request body size using `http.MaxBytesReader`. Return `413 Payload Too Large` with a JSON error body when exceeded.
1. Implement input sanitization for a user profile endpoint: strip HTML tags from string fields using `bluemonday`, normalize Unicode with `golang.org/x/text/unicode/norm`, and trim whitespace.
1. Write a SQL injection demonstration: show a vulnerable query with string concatenation, then fix it with parameterized queries. Write a test that attempts injection and verifies safety.

### 71.3 IP-Based Access Control

1. Implement IP allowlisting middleware: parse a list of CIDR ranges from configuration, check `r.RemoteAddr` (accounting for `X-Forwarded-For` behind a proxy), and reject unauthorized IPs with `403`.
1. Implement IP-based rate limiting: maintain a `map[netip.Addr]*rate.Limiter` with a mutex, create new limiters on first request from an IP, and evict idle limiters after 10 minutes.
1. Write a middleware that extracts the real client IP from `X-Forwarded-For`, `X-Real-IP`, or `r.RemoteAddr` with a configurable trusted proxy list.
