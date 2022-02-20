// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task2(c chan string, str string) {
for i := 0; i < 5; i++ {
c <- str + " "
}
//c <- " "
}