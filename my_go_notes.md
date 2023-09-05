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
turn address into value with *address
turn value into address with &value

## Usage of Go Mod
go mod init <project_name>