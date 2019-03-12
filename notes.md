# Notes

## Intro

## Basics
- Can use either factored or multiple import statements
- Names beginning with capital letter are 'exported' (i.e., can access from another package -- unexported names not accessible outside the package)
    * Package names can be lowercase though (and usually are?)
- In functions, arguments can have type at end of sequence of var names that are all the same type, e.g., (x, y int) 
- Crazy: Functions can return more than one value, e.g., func swap (x,y string) (string, string) and you can make two variables equal to the result of that, e.g., a,b := swap("hello", "world")
- Can name your return values, then "return" just returns that (that's a 'naked return')
- If an initializer is present, type is inferred from initializer 
- Can use 'short assignment' (:=) *only in a function*; has an implicit type (e.g., k := 3)
- Can use the %T in printing to output type of a variable, %v for the actual variable
- Variables w/o explicit initialized value given 'zero' value which is 0, false or "" as appropriate
- Casting requires explicit converstion -- like function call, e.g., Type(var)

### https://blog.golang.org/gos-declaration-syntax 
An interesting article about Go's declaration syntax. 
- Takeaways:
    * Read s.t. name comes first, then type (other than pointers)
    * f func(func(int,int) int, int) func(int, int) int means "f is a function that accepts a function of two ints, an int, and another int, and returns a function pointer, the function accepting two ints and returning an int"

## Flow Control
- Go only has for loops, and braces are always required
- while is just for, but with no init or increment statement and no semicolons around it 
- omitting condition for for loop is an infite loop (for {})
- Go's if statements also don't require parantheses, and always require braces
- If statements can have variable declarations in them (if v:= 5; v < 6 is true)
- Go switch statement runs only selected case, and cases can be whatever variable type you want
- Can use Switch to write clean long if-then-else chains
- 'defer' is a new statement -- defers execution of function until surrounding function returns (arguments evaluated immediatly, but function call not executed until surround returns)
    * Defer calls pushed onto stack, LIFO order execution
    * Good blog post on subject: 

### https://blog.golang.org/defer-panic-and-recover
A blog post about these new defer, panic, and recover features.
- Defer: pushes function call onto a stack. 3 simple rules:
    * Deferred function's arguments evaluated when defer statement is evaluated
    * Deferred functions executed in LIFO after surround function returns
    * Deferred function's may read and assign to returning function's named return values
- Panic: Stops ordinary flow control and starts "panicking"
    * Deferred functions executed normally
    * Next function up call stack gets panic called
    * Panic continues up stack until all functions in goroutine have returned, then program crashes
- Recover: Regains control of panicking goroutine
    * Useful only inside deferred functions (why? because only deferred functions will execute normally)
    * Normal execution (i.e., no panicking occurring): return nil, no other effects
    * Panicking Execution: Recover value given to panic, resume normal execution

## More Types
- Go has pointers, no pointer arithmetic 
- Go has structs ("type NAME struct {}")
    * Structs are initialized using brackets, e.g., Vertex{1,2}
    * Accessing member field is done with '.', whether or not the struct is raw or its a pointer to the struct
    * Can get address of struct using the & operator
- You can natively set a value using exponential notation, e.g., 1e9 will set something to a billion
- Go has arrays, and its length is part of its type so they can't be resized
    * "Slices" are dynamically-sized, subsets of the elements of an array - you form it with a high and low bound on another array, e.g., var s []int = primes[1:4], which includes elements 1-3 of primes (the high bound is excluded, low bound is included)
    * Slices are only descriptions, they don't actually store any data -- changing slice elements changes underlying data, and other slices will have changes reflected in them
