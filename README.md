# Revenge of the Pancakes
## Weave coding challenge

# summary
In this challenge a series of pancake possibilities are provided via CSV. I used a CSV because it is a very common method of storing data and I could use it to store and run the example scenario as well as several others. 
_The downside to using csv is that it expects all rows to be of the same length._

The purpose of this challenge is to get all of the pancakes face up in the fewest flips possible. A *face up* pancake is represented by a _+_ while the face down is represented by a _-_.

I used a series of for loops to implement a simple recursive algorithm that flips the stack or pancakes starting with the bottom pancake and working toward the top.

The example scenario provide is the first line in the CSV.

# to run

This code has been precompiled. It can be run by entering either
``` ./weave_code_example```
or 
``` go run main.go```
in the command prompt. There are no flags nor arguments.