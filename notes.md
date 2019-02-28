# Notes

## Intro

## Basics
- Can use either factored or multiple import statements
- Names beginning with capital letter are 'exported' (i.e., can access from another package -- unexported names not accessible outside the package)
* Package names can be lowercase though (and usually are?)

### https://blog.golang.org/gos-declaration-syntax 
An interesting article about Go's declaration syntax. 
- Takeaways:
* Read s.t. name comes first, then type (other than pointers)
* f func(func(int,int) int, int) func(int, int) int means "f is a function that accepts a function of two ints, an int, and another int, and returns a function pointer, the function accepting two ints and returning an int"
 
