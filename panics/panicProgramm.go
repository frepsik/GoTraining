package panics

import (
	"fmt"
)

func TestFunc() {
	defer func() {
		panic := recover()
		if panic != nil {
			fmt.Println("Сработала определённая паника, её нужно использовать, когда что то сильно отвалилось, но нам нужно это обработать.")
			fmt.Println("Сообщение паники: ", panic)
		}
	}()

	slice := []int{1, 3, 6}

	fmt.Println(slice[3])
}
