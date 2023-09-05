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
