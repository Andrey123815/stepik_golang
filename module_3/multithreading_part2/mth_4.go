func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
var c chan int = make(chan int)
//defer close(c)
go func() {
sum := 0
defer close(c)
for ;; {
select {
case v, _ := <- arguments:
sum += v
case _, ok := <- done:
if ok == true {
c <- sum
return
}
default:
fmt.Print("")
}
}
}()
//defer close(c)
return c
}