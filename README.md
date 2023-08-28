# golang-calculator

A CLI Calculator with math constants and variables, all written in Go.

## Usage

Run `./golang-calculator`, then enter a math expression with numbers, `()` as parenthesis, `+-*/` for basic operations, `%` for modulo, and `^` exponentiation. For example,

```
> ((5*76/19-3*3)^2+10)%89
2023/08/29 00:18:52 input:  "((5*76/19-3*3)^2+10)%89"
2023/08/29 00:18:52 tokens: [( ( 5 * 76 / 19 - 3 * 3 ) ^ 2 + 10 ) % 89]
2023/08/29 00:18:52 rpn:    [5 76 * 19 / 3 3 * - 2 ^ 10 + 89 %]
((5*76/19-3*3)^2+10)%89 = 42.000000
```

## Constants

The constants "e" and "pi" can be used in expressions. They are case insensitive. For example,

```
> pi
2023/08/29 00:13:12 input:  "pi"
2023/08/29 00:13:12 tokens: [pi]
2023/08/29 00:13:12 rpn:    [pi]
pi = 3.141593
> Pi
2023/08/29 00:13:16 input:  "Pi"
2023/08/29 00:13:16 tokens: [Pi]
2023/08/29 00:13:16 rpn:    [pi]
Pi = 3.141593
```

## Variables

You can save results as variables!

```
> phi = (1+5^0.5)/2
2023/08/29 00:09:00 input:  "phi = (1+5^0.5)/2"
2023/08/29 00:09:00 tokens: [( 1 + 5 ^ 0.5 ) / 2]
2023/08/29 00:09:00 rpn:    [1 5 0.5 ^ + 2 /]
2023/08/29 00:09:00 Key phi is new, will set it to 1.618034
phi = (1+5^0.5)/2 = 1.618034
> phi^2
2023/08/29 00:09:10 input:  "phi^2"
2023/08/29 00:09:10 tokens: [phi ^ 2]
2023/08/29 00:09:10 rpn:    [phi 2 ^]
phi^2 = 2.618034
```

### Rules for Variable Names

-   They must only be letters, and are case insensitive
-   They can't have the same name as existing constants (currently only `e` and `pi`)
-   They can't be the same name as commands (currently "help", "exit", and "quit")
