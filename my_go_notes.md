# My Go Notes

## Go Variable Types

>Go is a pass by value language like JavaScript and C#.
When we pass a value to a function, Go always creates a copy of that value and passes it to the function

> Value Types : Use pointers to change these things in a function

> Reference Types : Don't worry about pointers with these

| Value Types   | Reference Types |
| -------- | ------- |
| int  | slices    |
| float |  maps  |
| string    | channels   |
| bool    | pointers   |
| structs    | functions   |


## Usage of Pointer
> turn address into value with *address  
> turn value into address with &value

## Usage of Go Mod
> go mod init **<project_name>**

## Usage of Map

> Go Maps are like Objects in JavaScript. Both keys and values staticly typed so all keys must be exact type and all values must be exact type.

| Map  | Struct |
| -------- | ------- |
| All keys must be the same type  | Values can be of different type    |
| All values must be the same type    |  Keys don't support indexing  |
| Keys are indexed - we can iterate over them    | You need to know all the different fields at compile time   |
| Use to represent a collection of related properties    | Use to represent a "thing" with a lot of different properties   |

## Usage of Interface

> In the Go programming language, interfaces are used to define a set of methods that a type must implement. Interfaces enable you to write code that is more flexible, reusable, and testable by abstracting away the specific implementation details of types. 

* Interfaces are not generic types
* Interfaces are implicit
* Interfaces are a contract to help us manage types

### Here are some common use cases for interfaces in Go:

1. Polymorphism: Interfaces allow you to write functions and methods that can work with different types as long as those types satisfy the interface. This enables you to achieve polymorphism in Go, where you can treat different types that implement the same interface in a uniform way.

2. API Design: When designing libraries or packages, interfaces provide a way to define a contract that users of your code should follow. By specifying interfaces, you can document how types should behave and what methods they must implement.

3. Dependency Injection: Interfaces are often used in dependency injection patterns. Instead of depending on concrete types, you depend on interfaces. This makes it easier to swap out implementations or mock dependencies for testing.

4. Testing: Interfaces make it easier to write unit tests by allowing you to create mock implementations of dependencies. This is particularly useful when you want to isolate the code you're testing from external systems, such as databases or web services.

5. Plugin Systems: Interfaces can be used to create plugin systems where external code can be dynamically loaded and used, as long as it adheres to the defined interface

> Here's a simple example of how an interface is defined and used in Go:

```go
package main

import "fmt"

// Define an interface named Writer with a single method, Write.
type Writer interface {
    Write([]byte) (int, error)
}

// Create a concrete type, FileWriter, that implements the Writer interface.
type FileWriter struct {
    FileName string
}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implement the Write method for FileWriter.
    // Write data to the specified file.
    // Return the number of bytes written and any errors encountered.
    // ...
    return len(data), nil
}

func main() {
    // Create a FileWriter instance and use it through the Writer interface.
    var w Writer
    w = FileWriter{FileName: "example.txt"}
    w.Write([]byte("Hello, World!\n"))

    // You can also create other types that implement the Writer interface
    // and use them interchangeably.
}
```


## Usage of Channels and Go Routines

> In the Go programming language, Go Routines and Channels are powerful features for concurrent and parallel programming. They make it easy to write efficient, concurrent programs that can perform tasks concurrently and communicate safely between goroutines. Here's an overview of their usage:

### Go Routines

* Lightweight Concurrency: Goroutines are lightweight threads of execution that allow you to run functions concurrently without the overhead of creating a full OS thread for each one. You can spawn thousands or even millions of goroutines in a single Go program.
* Concurrency: They enable concurrent execution of tasks, making it easier to write efficient, non-blocking code. You can use goroutines for tasks such as handling concurrent HTTP requests, processing data in parallel, or running background tasks concurrently.
* Syntax: Goroutines are started using the go keyword followed by a function call. For example: go myFunction().
* Example:
```go 
func main() {
    go sayHello()
    go sayWorld()
    time.Sleep(time.Second)
}

func sayHello() {
    fmt.Println("Hello")
}

func sayWorld() {
    fmt.Println("World")
}
```

> Biggest TAKEAWAY with Go Routine that we never, ever try to access the same variable from a different child routine wherever possible

### Channels
* Synchronization and Communication: Channels are used for communication and synchronization between goroutines. They provide a safe way for goroutines to send and receive data and signals.
* Blocking: Sending data to a channel blocks until another goroutine is ready to receive it, and vice versa. This blocking behavior helps coordinate the timing of tasks and avoids race conditions.
* Unidirectional and Buffered Channels: You can create unidirectional channels (send-only or receive-only), and you can also create buffered channels to control the capacity of data that can be sent without blocking.
* Example:
```go
func main() {
    ch := make(chan string)
    go sendData(ch)
    go receiveData(ch)
    time.Sleep(time.Second)
}

func sendData(ch chan string) {
    ch <- "Hello"
    ch <- "World"
    close(ch)
}

func receiveData(ch chan string) {
    for msg := range ch {
        fmt.Println(msg)
    }
}
```

### Select Statement
* The select statement is used to choose between multiple communication operations on channels. It allows you to handle multiple channels concurrently and non-blocking.
* It's often used for handling timeout scenarios, graceful shutdown, and multiplexing multiple channels.
* Example:
```go
select {
case msg1 := <-ch1:
    fmt.Println("Received", msg1)
case ch2 <- "Hello":
    fmt.Println("Sent a message to ch2")
case <-time.After(time.Second):
    fmt.Println("Timed out")
}
```
> In summary, goroutines and channels are fundamental to Go's concurrency model. They enable you to write concurrent programs that are both efficient and easy to reason about. Go Routines allow you to execute functions concurrently, and channels provide a safe way for go routines to communicate and synchronize their actions. This combination makes Go well-suited for tasks involving parallelism, concurrency, and distributed systems.




