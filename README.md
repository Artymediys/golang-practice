# Tasks

1) Given a Human structure (with an arbitrary set of fields and methods).
Implement embedding methods into Action structure from parent Human structure (analog of inheritance).


2) Write a program that concurrently calculates the value of the squares of numbers
taken from the array (2,4,6,8,10) and outputs their squares in stdout.


3) Given a sequence of numbers: 2,4,6,8,10.
Find the sum of their squares(2^2 + 3^2 + 4^2....) using concurrent computation.


4) Realize constant data writing to the channel (main thread).
Implement a set of N workers which read arbitrary data from the channel and output to stdout.
You need to be able to choose the number of workers at startup.

    The program should end by pressing `Ctrl+C`. Select and justify the method of terminating all the workers.


5) Develop a program which sequentially sends values to the channel and reads from the other side of the channel.
After N seconds the program should stop.


6) Implement all possible ways of stopping the goroutine execution.


7) Implement concurrently writing of data to map.


8) A variable `int64` is given. Develop a program that sets the i-th bit to 1 or 0.


9) Develop a number pipeline. Given two channels: the first is to write numbers (x) from the array,
the second is to write the result of x*2 operation, then data from the second channel should be output to `stdout`.


10) A sequence of temperature fluctuations is given: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Combine these values into groups in 10 degree steps. The sequence in the subsets is not important.

    Example: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.


11) Implement the intersection of two disordered sets.


12) There is a sequence of strings - (cat, cat, dog, cat, tree) create an own set for it.


13) Swap two numbers without creating a temporary variable.


14) Develop a program which is able to determine the type of variable in runtime:
`int`, `string`, `bool`, `channel` from a variable of type `interface{}`.


15) What are the negative consequences of this code fragment and how to fix it?
Give a correct example of the implementation.

    ```
    var justString string
    func someFunc() {
        v := createHugeString(1 << 10)
        justString = v[:100]
    }
    
    func main() {
        someFunc()
    }
    ```

16) Implement quicksorting of the array (quicksort) with built-in language methods.


17) Implement binary search with built-in language methods.


18) Implement a counter structure to be incremented in a concurrent environment.
On completion, the program should output the final counter value.


19) Develop a program that flips the string (e.g.: "главрыба — абырвалг"). The characters can be unicode.


20) Develop a program that flips the words in a string.

    Example: "snow dog sun - sun dog snow".


21) Implement an "adapter" pattern on any example.


22) Develop a program that multiplies, divides, adds, subtracts two numeric variables a,b whose value > 2^20.


23) Remove the i-th element from the slice.


24) Develop a program to find the distance between two points that are represented as a Point structure
with encapsulated parameters x, y and a constructor.


25) Implement your own sleep function.


26) Develop a program that checks that all characters in a string are unique (true if unique, false etc.).
The check function must be case insensitive.

    For example:
    ```
    abcd - true
    abCdefAaf - false
    aabcd - false
    ```