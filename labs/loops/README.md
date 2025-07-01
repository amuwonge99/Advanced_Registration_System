LOOPS
=====

[Loops](https://go.dev/tour/flowcontrol/1) are always an important part of coding and here are some examples of how to set them up and use them.

WAYS OF MAKING LOOPS
--------------------

There are two mains ways of running loops in Go, the first is very typical in coding and should look very familiar if you have ever coded in C or Java.

```go
for i := 0; i < 10; i++ {
	// do something
}
```

In this example we create an integer called `i`, give it a value of `0` and then while `i` is under `10`, run the loop and at the end of each loop increment `i`. Very typical for-loop.

We also have for-range-loops which use [slices](https://go.dev/tour/moretypes/7), in Go you can think of slices the same way other languages use arrays, the only difference is that they do not have a fixed size so they can grow or contract as required.

```go
names := []string{"Daniel", "Rupert", "Emma"}
for i, name := range names {
    print(i)
    print(name)
}
```

When we use the [`range`](https://go.dev/tour/moretypes/16) keyword for loops, we get back two values: the element and the value. So, in the example above in the first loop we would get `i=0` and `name=Daniel`, the second would be `i=1` and `name=Rupert`, and finally `i=2` and `name=Emma`.

If we don't care about one of the two values, e.g. we only care about the names and we don't want to have the element number, we can replace it with an underscore to throw away the value.

```go
for _, name := range names {
  print(name)
}
```

LAB TASK
--------

```go
names := []string{"Daniel", "Rupert", "Emma"}
colours := []string{"Red", "Blue", "Green", "Yellow"}
```

Create a program using loops that creates following output using the two slices above.

```txt
Daniel
Red
Blue
Green
Yellow
Rupert
Red
Blue
Green
Yellow
Emma
Red
Blue
Green
Yellow
```


Enhance your understanding of loops by dynamically generating output based on user-defined slices and using nested loops with conditional logic.

🧪 Task:
Create two new slices:
```go
animals := []string{"Cat", "Dog", "Rabbit"}
sounds := []string{"Meow", "Woof", "Squeak", "Growl"}
```

Write a program that prints each animal followed by all the sounds except the one that doesn't match the animal. Use a for-range loop and an if statement to skip the incorrect sound.

Example output:
```txt
Cat
Meow
Growl
Dog
Woof
Growl
Rabbit
Squeak
```

💡 Hint: Use an if statement inside your inner loop to skip a sound that doesn't match the animal.

Bonus: Add a third slice called favourites:
```go
favourites := []string{"Cat", "Squeak"}
```

Modify your program to print a ⭐ next to any animal or sound that appears in the favourites slice.

Example output:
```txt
Cat ⭐
Meow ⭐
Growl
Dog
Woof
Growl
Rabbit
Squeak ⭐
```
