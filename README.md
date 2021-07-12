## Basic calculator to evaluate expression

Expression consists of single digits (like '2', '7' etc), `'+'`, `'-'` and `' '`)

## The calculator package has a function

```func Evaluate(input string)(int, error)```

 which takes ```string``` and returns ```(int, error)```
## Installation:

```go get -u github.com/borichevskiy/expression_calculator```
## Usage:

```go
str := "1 + 2 - 3"
res, err := Evaluate(str)
```
## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
