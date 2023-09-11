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