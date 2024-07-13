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