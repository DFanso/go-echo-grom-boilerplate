### Go Concepts Quick Reference

This guide provides a comprehensive overview of essential Go concepts, syntax, and features. It serves as a quick reference for understanding the language, covering variables, data types, control structures, functions, error handling, structs, methods, interfaces, concurrency, and more. Each section includes brief explanations and code examples where applicable.

---

#### 1. **Variables and Data Types**
- **Variable Declaration**:
  - Use `var` for explicit declaration, with optional type and initial value.
    ```go
    var x int = 10
    var y int  // Defaults to zero value (0 for int)
    ```
  - Use `:=` for short variable declaration with type inference (only in functions).
    ```go
    x := 10  // int
    name := "Go"  // string
    ```
- **Basic Data Types**:
  - `int`, `float64`, `string`, `bool`, etc.
- **Zero Values**:
  - Default values for uninitialized variables:
    - `0` for `int`
    - `""` for `string`
    - `false` for `bool`
    - `nil` for pointers, slices, maps, channels, interfaces, and functions.

---

#### 2. **Constants**
- Defined with `const`, immutable.
  ```go
  const Pi = 3.14
  ```

---

#### 3. **Pointers**
- Hold memory addresses of variables.
- `*T` is a pointer to a `T` value, `&x` gets the address of `x`.
  ```go
  var p *int
  x := 42
  p = &x  // p points to x
  fmt.Println(*p)  // Prints 42
  ```

---

#### 4. **Control Structures**
- **If Statements**:
  - No parentheses needed.
  - Can include a short statement before the condition.
    ```go
    if x := 10; x > 5 {
        fmt.Println("x is greater than 5")
    }
    ```
- **For Loops**:
  - The only loop construct in Go.
  - Can be used as a traditional for loop, while loop, or infinite loop.
    ```go
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    // While-like loop
    x := 0
    for x < 5 {
        fmt.Println(x)
        x++
    }
    ```
- **Switch Statements**:
  - No fallthrough by default.
  - Can switch on any type.
    ```go
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("macOS")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Println("Other")
    }
    ```

---

#### 5. **Functions**
- Defined with `func`, can have multiple return values.
  ```go
  func add(a int, b int) int {
      return a + b
  }
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, errors.New("division by zero")
      }
      return a / b, nil
  }
  ```

---

#### 6. **Error Handling**
- Errors are values, often returned as the last return value.
- Use `if err != nil` to check for errors.
  ```go
  result, err := divide(10, 0)
  if err != nil {
      fmt.Println("Error:", err)
  } else {
      fmt.Println("Result:", result)
  }
  ```

---

#### 7. **Structs**
- User-defined types composed of fields.
  ```go
  type Person struct {
      Name string
      Age  int
  }
  p := Person{Name: "Alice", Age: 30}
  ```

---

#### 8. **Methods**
- Functions with a receiver, attached to a type.
  ```go
  func (p Person) Greet() {
      fmt.Println("Hello, my name is", p.Name)
  }
  p.Greet()  // Calls the method
  ```

---

#### 9. **Interfaces**
- Define a set of methods that a type must implement.
- Types implicitly satisfy interfaces.
  ```go
  type Shape interface {
      Area() float64
  }
  type Circle struct {
      Radius float64
  }
  func (c Circle) Area() float64 {
      return 3.14 * c.Radius * c.Radius
  }
  var s Shape = Circle{Radius: 5}
  fmt.Println(s.Area())  // Calls Circle's Area method
  ```

---

#### 10. **Concurrency**
- **Goroutines**: Lightweight threads managed by the Go runtime.
  ```go
  go func() {
      fmt.Println("Hello from goroutine")
  }()
  ```
- **Channels**: Communicate between goroutines.
  ```go
  ch := make(chan int)
  go func() {
      ch <- 42  // Send
  }()
  value := <-ch  // Receive
  ```

---

#### 11. **Packages and Imports**
- Code is organized into packages.
- Use `import` to access other packages.
  ```go
  import "fmt"
  import (
      "os"
      "time"
  )
  ```

---

#### 12. **Slices and Maps**
- **Slices**: Dynamic arrays.
  ```go
  s := []int{1, 2, 3}
  s = append(s, 4)
  ```
- **Maps**: Key-value pairs.
  ```go
  m := make(map[string]int)
  m["key"] = 1
  ```

---

#### 13. **nil**
- Represents the zero value for pointers, slices, maps, channels, interfaces, and functions.
- Often used to indicate "no value" or "absence."
  ```go
  var p *int  // p is nil
  var s []int  // s is nil
  var m map[string]int  // m is nil
  ```

---

#### 14. **Type Assertions and Type Switches**
- **Type Assertion**: Extract the underlying value of an interface.
  ```go
  var i interface{} = "hello"
  s, ok := i.(string)  // s = "hello", ok = true
  ```
- **Type Switch**: Switch on the type of an interface.
  ```go
  switch v := i.(type) {
  case string:
      fmt.Println("string:", v)
  case int:
      fmt.Println("int:", v)
  default:
      fmt.Println("unknown type")
  }
  ```

---

#### 15. **Defer**
- Schedules a function call to run after the surrounding function returns.
  ```go
  func main() {
      defer fmt.Println("world")
      fmt.Println("hello")
  }
  // Output:
  // hello
  // world
  ```
