# Stack

This is a stack implementation with generics in golang.

## Desciption

A stack is a linear data structure that is used to store a collection of objects. It is based on Last-In-First-Out (LIFO). A good example of a stack is a pile of plates in a cafeteria. The plates are added to the top of the pile and removed from the top of the pile. The most recently added plate is the first one to be removed.

The LIFO (Last In First Out) principle entails inserting a single element at the end of a stack, i.e., the element inserted at the end of the stack is the first element to appear.

Operations:

- `Pop` operations are the insertion of an element into the stack and the deletion of an element from the stack
- `Push` operations are the insertion of elements into the stack

## Example Usage

First create a `stack.New` to use throughout your application:

```go

s := stack.New[string]()

s.Push("example-1")
s.Push("example-2")

fmt.Println(s.Pop())
fmt.Println(s.Pop())

```
