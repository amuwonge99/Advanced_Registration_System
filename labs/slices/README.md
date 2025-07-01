Slices
======

In this lab, we will be exploring slices and a few examples of how to use them.

WHAT IS A SLICE?
-----------------
An slice in Go is similar to an array, except slices don't need to specify a fixed size when created and they can grow/shrink in size.

CREATING SLICES
---------------
Slices are declared in a similar way to arrays without specifying the size. (https://go.dev/tour/moretypes/7).

For example:
```go
var ages = []int{1, 8, 7, 2, 9}
```
declares a variable `ages` which contains a slice with 5 integers.

There are lots of neat things you can do with slices, for example, you can refer to a smaller part of an array using a slice, e.g

```go
    var ages = []int{1, 8, 7, 2, 9}
    var first_three_ages = ages[:3]
    fmt.Println(first_three_ages)
```
The above should print out only the first 3 ages like this:

```go
[1 8 7]
```
To print the last 3 ages, use something like this. Note that you don't have to specify the last index:
```go
    var ages = []int{1, 8, 7, 2, 9}
    var last_three_ages = ages[2:]
    fmt.Println(last_three_ages)
```

The syntax for specifying the range is `[first_index:last_index]`. Slices are indexed similarly to arrays so they start at 0 and end at `(size of array - 1)`. In other words, our example `ages` array above starts at index 0 and ends at index 4. 

Another great thing about slices is that you can add/remove elements from them.

LAB TASK 1
----------
For this lab task, let's use another real-world example of adding and removing items from our shopping list.

Here's our shopping list:
```go
    var shoppingList = []string{"apples", "oranges", "bread", "tea", "milk"}
```

Write some code in Go that will add an extra item, "sugar", to the shopping list.

Your program should output:
```go
[apples oranges bread tea milk sugar]
```

LAB TASK 2
----------
For the same shopping list above, write some code that will remove the "third" item (bread) from the shopping list.

Your program should output:
```go
[apples oranges tea milk sugar]
```

Note that you will have to use a special feature in Go that "flattens" the arrays to allow you to merge them.