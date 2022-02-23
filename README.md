# MarsRoverGolang

This is a version of the task written in Golang, as a challenge to myself to complete the task in a language I am largely unfamiliar with.

## Getting Started
1. `docker build --tag golang-wise-robot-test .`
2. `docker run --rm -it golang-wise-robot-test`
3. When prompted, enter the grid size in `x y` format
4. When prompted, enter the number of robots as an integer
5. When prompted, enter the initial state and commands of the robots in `(x, y, orientation) COMMANDS` format
    - You will be prompted multiple times to enter robot details if you entered a number of robots larger than 1 in step 4

## Improvements
Due to inexperience with Golang, I feel as if I could make many further improvements to the architecture and speed of this program. With more time I'm sure I could improve this by:
- Finding more efficient ways of storing/managing data and executing functions on that data
- Learning how to properly split the program into multiple files so as to make the code more human readable
- Learning how to represent data in Golang with certain datatypes that aren't available as they are in other languages, such as Enums
- Adding more type checking and input checking to ensure that erroneous values cannot be entered
