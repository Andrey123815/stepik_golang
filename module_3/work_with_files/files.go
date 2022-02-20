// package main уже объявлен.
func main() {
// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
sum := 0
scanner := bufio.NewScanner(os.Stdin)
//err := scanner.Scan()
for scanner.Scan() {
i, _ := strconv.Atoi(scanner.Text())
//io.WriteString(os.Stdout, string(i))
sum += i
//err = scanner.Scan()
}
io.WriteString(os.Stdout, strconv.Itoa(sum))
}