## Go Notes

### On Strings
Strings in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it.

Go also has a type that represents a single code point. The `rune` type is an alias for the `int32` type, just as `byte` is an alias for `uint8`.

### On Truthiness
**Go doesn’t allow truthiness.**
In fact, no other type can be converted to a bool, implicitly or explicitly. If you want to convert from another data type to boolean, you must use one of the comparison operators (==, !=, >, <, <=, or >=).
For example, to check if variable x is equal to 0, the code would be x == 0. If you want to check if string s is empty, use s == "".

### On the ':=' Operator
Using := has one limitation. If you are declaring a variable at the package level, you must use var because := is not legal outside of functions.

### On Variable Names
Go doesn’t use **snake case** (names like `index_counter` or `number_tries`). Instead, idiomatic Go uses **camel case** (names like `indexCounter` or `numberTries`) when an identifier name consists of multiple words.

In many languages, constants are always written in all uppercase letters, with words separated by underscores (names like `INDEX_COUNTER` or `NUMBER_TRIES`). Go **does not** follow this pattern. This is because Go uses the case of the first letter in the name of a package-level declaration to determine if the item is accessible outside the package.

### On Multidimensional Arrays
Go has only one-dimensional arrays, but you can simulate multidimensional arrays:
```Go
var x [2][3]int
```
This declares `x` to be an array of length 2 whose type is an array of `ints` of length 3.
This sounds pedantic, but some languages have true matrix support, like Fortran or Julia; Go isn’t one of them.

### On Array Limitations
Go considers the size of the array to be part of the type
of the array. This makes an array that’s declared to be [3]int a different type from
an array that’s declared to be [4]int. This also means that you cannot use a variable
to specify the size of an array, because types must be resolved at compile time, not at
runtime.