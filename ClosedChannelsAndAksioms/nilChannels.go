package closedchannelsandaksioms

import "fmt"

//Функция демонстрирующая аксиомы каналы, на примере nil канала
func TestFuncInNilChannel() {
	//Создаём канал со значением nil, без инициализации
	var chan1 chan int

	//Пробуем записать значение в nil канал, произойдёт deadlock
	go func() {
		chan1 <- 2
	}()

	//Проубем закрыть nil канал, произойдёт паника
	close(chan1)

	//Прробуем считать значение с nil, канала произойдёт deadlock
	value := <-chan1
	fmt.Println(value)
}
