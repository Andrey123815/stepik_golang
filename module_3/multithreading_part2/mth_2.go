// вы уже внутри main()

wg := new(sync.WaitGroup)


for i := 0; i < 10; i++ {
wg.Add(1) // Увеличиваем счетчик горутин в группе
go func(wg *sync.WaitGroup) {
defer wg.Done()
work() // Вызываем функцию work в отдельной горутине
}(wg)
}

wg.Wait() // ожидаем завершения всех горутин в группе
