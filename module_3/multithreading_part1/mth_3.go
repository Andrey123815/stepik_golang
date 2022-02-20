// Пакет и функция main уже объявлены.
// Выводить или вводить ничего не нужно!
func removeDuplicates(inputStream  chan string, outputStream chan string) {
defer close(outputStream)
var prev string
for {
v, ok := <- inputStream
if !ok {
break
}
if v != prev {
outputStream <- v
prev = v
}
}
}
