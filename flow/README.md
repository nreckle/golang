# NOTE

for 是 Go 中的 “while”
此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。

```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
```go
func main() {
	for {
	}
}
```
`if.go`
if
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。

`if-with-a-short-statement.go`
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
该语句声明的变量作用域仅在 if 之内。

`switch.go`
Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。
实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

**defer** 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

defer 栈
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。