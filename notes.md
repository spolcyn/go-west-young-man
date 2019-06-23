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
    * An array var "denotes the entire array", not a pointer to the first element; thus, array arguments are copied, not referenced.
    * "Kind of like a struct but with indexed rather than named fields"
- Slices 
    * "Slices" are dynamically-sized, subsets of the elements of an array - you form it with a high and low bound on another array, e.g., var s []int = primes[1:4], which includes elements 1-3 of primes (the high bound is excluded, low bound is included)
    * Slices are only descriptions, they don't actually store any data -- changing slice elements changes underlying data, and other slices will have changes reflected in them
    * Slice literals create an array, then the slice that references it (e.g., []bool{true, true, false} creates then references an array of length 3 with that data)
    * Slice defaults for low bound is 0 and length of slice (read: array that it references) for high bound
    * Slice has length (number of elements it contains), capacity (number of elements of underlying array, counting from first element in slice)
    * Nil slice is the zero value of a slice
    * the 'make' function can make slices, usage make([]type, length, capacity) - this is the dynamic-array sizing method
    * Slice can't be grown beyond its capacity/can't be re-sliced below 0 to access earlier elements
    * For slice copy, destination argument comes before source
    * Can use "..." to expand a slice to its elements, e.g. "b := []int{1,2,3} // a = append(a, b...)
    * Since re-slicing doesn't create a new array, if working with a small part of big array, copy the slice and let the old one get garbage collected
- Range: iterate over slice or map -- two vars for each iteration, one the index, one a copy of the element at that index
    * Can skip using index or value by the " _ " character, e.g., for i, _ := range pow. If you just want the index, can just straight omit the second var (so it seems for i, _ := range pow and for i := range pow are the same)
- Maps: maps key to values; make() can make a map like for slices. map[key]Value
    * Map literals can be declared with their keys & values - "if top-level type is just a type name, you can omit it from the elements of the literal" - Question: What is a top-level type, and when would it not just be a type name? A slice, array?
    * Accessing done with array notation, deletion done w/ the delete(map, key) function.
    * Testing key present done w/ 2-value assignment (elem, ok = m[key] -> ok true iff key is in m, if ok=false, then elem is zero value for map's element type)
- Function values: Can pass functions like arguments. Can also define functions quickly within a method, e.g., func main() {hypot := func(x, y float64) float64 { return math.Sqrt(x*x + y*y)}
- Function closures: Go functions may be 'closures'
    * Closure: https://stackoverflow.com/questions/36636/what-is-a-closure
    * A closure is a function that retains access to all the local variables in scope when it was called. So, if you have a function funcParent and define some variable 'parentVar' there, then define a function funcChild in funcParent and you return funcChild, then when you call funcChild later, parentVar will still be accessible by it.
    * A closure is a "persistent local variable scope" (per SO) or the "function is 'bound' to the variables" (per tut) or the closure "encloses" the scope (per me?) 

## Methods

- Methods: A perfectly normal function, just with a special "receiver" argument. This appears between the func keyword and method name
    * e.g.: func (v Vertex) Abs() float64 { return math.Sqrt(v.X * v.X + v.Y * v.Y) (Vertex v is the receiver)
    * Equivalent: func Abs(v Vertex) float64 { ... }
    * Go doesn't have classes, so methods are defined on types. Then you can call the first example as v.Abs()
    * Can only declare method with receiver whose type is defined in same package as the method (so no using built-in types like int, though you can redefine to be like 'type MyInt int' and then use MyInt). Moreover, it can be used for non-struct types, like the just-named example for MyInt (=int)
    * Recievers can also be pointers, which allows you to modify the receiver. When it's a value receiver, the method operates on a copy.
    * For methods with receivers, the compiler correctly interprets how to treat a receiver if it's given as a value or as a pointer (e.g., if the receiver is a pointer, (&<var>).func() is implicitly called)
    * The same thing happens when calling the method on a pointer or on a value (e.g., v.func() and pV := &v pV.func() evaluate to the same thing)
- Choosing value vs. pointer receiver:
    * Pointer: Want method to modify value, and avoid copying value on each method call. (sounds like mostly want to use it this way)
    * In general, all methods on given type should be EITHER pointer OR value

## Interfaces

- Definition: A set of method signatures
    * Value of Interface type can hold any value implementing those methods
- Implementation is implicit -- no explicit "implements" keyword or anything of the source. 
- Implicit interfaces decouple the interface def & implementation (why is this useful? - one site suggests is that you can define an interfaec that's automatically implemented by types already written)
- An interface is in fact a value: tuple as such: (value, type) so an interface value holds a value of a "specific underlying concrete type")
- Interestingly, Go variables seem to have the function they were in as part of their type (see methods part 11 for an example)
- You can call methods with nil receivers (executes just fine) - note "an interface value that holds a nil concrete value is, itself, non-nil"
- Nil-interface value holds NEITHER value nor concrete type -- thus, calling a method on this is a run-time error (no type w/in interface, so don't know concretely which method to call)
- The empty interface specifies 0 methods, and thus may hold values of any type (b/c every type implements at least 0 methods)
    * Use case: Code that handle values of any unknown type, e.g., fmt.Print takes arguments of type interface{}
- Type Assertion: Get access to the underlying concrete value. If there's no underlying concrete value of the type requested, panic! 
    * You can use a 2-return value type though, which returns the underlying value and a boolean (underlying is the 0 value for that type if failure)
- Type Switch: Use to run different cases depending on which type a variable is
- Stringers: Like a toString() Java method, usually looked for when wanting to describe something as a string

## Errors
- Errors are also interfaces. nil error is success, non-nil denotes some type of error
- When you print an error, it calls a method that implements the interface (I guess in that file?)
- The "Exercise: Errors" asks why you get an infinite loop if calling fmt.Sprint(e) inside the error method, but works fine if you cast e to a float64 first. The reason is because the func Error() implements the String interface, so will just be repeatedly called when the program attempts to convert the function to a string, either from Sprint() or from a format string.
- Go does some weird things with semi-colon insertion, so any statements that go together can't allow a brace to end the compoundness

## Goroutines

- A lightweight thread managed by the Go runtime
    * Threads run in the same address space, access to shared memory must be synchronized
- Channels: "Typed conduit through which you can send and receive values with the channel operator, `<-`"
    * ch <- v // send v to channel ch
    * v := <-ch // receive from ch & assign value to v
    * Must be created before use: ch := make(chan int) (<- this is an int channel)
    * Sends & receives block until other side is ready (enables syncing w/o explicit work)
- Channel Buffering: provide buffer length as second arg, and then sends & receives only block when buffer is full/empty, respectively 
    * If you modify the concurrency/3 example to overfill the buffer, you get a "all goroutines are asleep - deadlock!" error. neat.
