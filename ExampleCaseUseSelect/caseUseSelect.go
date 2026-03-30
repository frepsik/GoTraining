package examplecaseuseselect

import (
	"fmt"
	"time"
)

// Более корректный пример реализации механизма Select, в случае, если тебе пишет несколько друзей в чате
func TestFunc() {
	chanFirst := make(chan Message)
	chanSecond := make(chan Message)
	chanThird := make(chan Message)

	go func() {
		for {
			chanFirst <- Message{
				Author: "Киря",
				Text:   "Привет, пошли по шаве?",
			}
			time.Sleep(4 * time.Second)
		}
	}()

	go func() {
		for {
			chanSecond <- Message{
				Author: "Макс",
				Text:   "Димас, можно к виртуалке подключиться ?",
			}
			time.Sleep(2 * time.Second)
		}
	}()

	//Вот к примеру, вот здесь, человек ушёл на 10 секунд, и если к примеру, оно бы у нас всё выполнялось по порядку, и мы бы сидели ждали, пока он вернётся
	//то есть выполняли бы код по порядку, тогда, мы бы не увидили, что в этот момент, написали другие друзья
	go func() {
		for {
			chanThird <- Message{
				Author: "Андрюха",
				Text:   "Ну короче тут вообще, щас такое расскажу и пропал на пол дня",
			}
			time.Sleep(10 * time.Second)
		}
	}()

	//По итогу, в Select отработает первым, тот кейс, где канал к моменту, когда мы подойдём сюда, будет уже заполнен параллельно выполняющимся одним из потоков
	for {
		select {
		case msg1 := <-chanFirst:
			fmt.Println("Отправитель: ", msg1.Author, "\nСообщение: ", msg1.Text, "\n-------")
		case msg2 := <-chanSecond:
			fmt.Println("Отправитель: ", msg2.Author, "\nСообщение: ", msg2.Text, "\n-------")
		case msg3 := <-chanThird:
			fmt.Println("Отправитель: ", msg3.Author, "\nСообщение: ", msg3.Text, "\n-------")
		}
	}
}
