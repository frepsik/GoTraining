package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	numberPostman int,
	mail string,
) {

	defer wg.Done()

	//Если бы время выполнения работы у почтальонов было больше, соответственно и писем бы они разнесли больше каждый, в связи с тем, что тут вечный цикл
	// (с количеством секунд выполнения в Main можно поиграться)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон ", numberPostman, "закончил свою работу")
			return
		default:
			fmt.Println("Я почтальон ", numberPostman, "начал работу")
			time.Sleep(1 * time.Second)

			transferPoint <- mail
			fmt.Println("Я почтальон ", numberPostman, "передал письмо. Письмо: ", mail)

		}
	}
}

func PoolPostman(ctx context.Context, amountPostman int) <-chan string {
	postmanTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= amountPostman; i++ {
		wg.Add(1)
		go postman(ctx, wg, postmanTransferPoint, i, postmanToMail(i))
	}

	//Добавили данный функционал в целях того, если мы будем возвращать в main уже закрытый канал, тогда мы не сможем прочитать его содержимое, но в связи с тем,
	//что нам это необходимо, мы запускаем отдельную горутину, как только доходим до wg.Wait() она блокируется, соответственно мы ждём как все wg.Done() отработают
	//и как только это случится, мы закроем канал и соответственно в main, при попытке считать новое значение из канала, цикл поймёт под капотом, что канал закрыт, считывается
	//значение по умолчанию, соответственно можно заканчивать с чтением от туда. Там поток заблокируется, если будет не в отдельной горутине, и будет ждать, пока у нас
	//тут всё отработает или же заблокируется, но в нашем сценарии, канал закроется, как только все горутины отпишут в него значения по очереди
	//Если этого не будет, основной main просто встанет (заблокируется), пока мы будем ждать завершения работ горутин всех
	go func() {
		wg.Wait()
		close(postmanTransferPoint)
	}()

	return postmanTransferPoint
}

// Вспомогающая функция позволяющая определить, какое письмо потащит почтальон
func postmanToMail(numberPostman int) string {
	mailForPostman := map[int]string{
		1: "Письмо из сервиса",
		2: "Приглашение на свадьбу",
		3: "Письмо со счтеами",
	}
	mail, ok := mailForPostman[numberPostman]
	if !ok {
		return "Вы выиграли в лотерею !"
	}
	return mail
}
