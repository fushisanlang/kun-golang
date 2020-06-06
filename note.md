## 开发笔记



1. 颜色显示

```go
color.Set(color.FgRed)
defer color.Unset()
fmt.Println(a)
color.Set(color.FgGreen, color.Bold)
fmt.Println(b)
```


2. 结构体跨包调用
属性名首字母也要是大写


3. chan使用
```go
func main() {
	ch1 := make(chan int)
    go func1(ch1)
    go func2(ch1)
}
func func1(ch1 chan int) {
	for i:=1 ;1<10;i++ {
		ch1 <- i
	}
}
func func2(ch1 chan int) {
	for i:=1 ; i<10 ; i++ {
		j := <- ch1
	}
}

//main函数中make，在调用子函数时，将chan作为变量传入子函数
```
