func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
var c chan int = make(chan int)
go func() {
defer close(c)
select {
case v, _ := <- firstChan:
//fmt.Println("Получено значение из первого канала")
c <- v * v
case v, _ := <- secondChan:
//fmt.Println("Получено значение из первого канала")
c <- 3 * v
case _, ok := <- stopChan:
//fmt.Println("Получено значение из первого канала")
if ok == true {
break
}
// Блок default выполнится раньше блока case - 1 секунда слишком много для Go
default:
fmt.Print("")
}
}()
return c
}
