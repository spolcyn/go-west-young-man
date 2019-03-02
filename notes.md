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
 
