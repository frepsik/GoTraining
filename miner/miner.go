package miner

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func miner(
	wg *sync.WaitGroup,
	ctx context.Context,
	transferPoint chan<- int, //Канал только на запись
	numberMiner int,
	power int,
) {

	defer wg.Done()

	//Если бы время выполнения работы у шахтёров было больше, соответственно и угля бы они добыли больше каждый, в связи с тем, что тут вечный цикл
	// (с количеством секунд выполнения в Main можно поиграться)
	for {

		//2 конструкция - данная конструкция имеет смысл, в случае, если необходимо завершить работу сразу, без возможности закончить, что было начато, то есть шахтёры начали
		//добывать уголь, и как только контекст завершился они не начали передавать его, бросили и завершили данную операцию

		fmt.Println("Я шахтёр ", numberMiner, "начал добывать уголь")
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр ", numberMiner, "завершил свою работу")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтёр ", numberMiner, "добыл уголь. Количество: ", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтёр ", numberMiner, "завершил свою работу")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтёр ", numberMiner, "передал уголь. Количество: ", power)
		}

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Я шахтёр ", numberMiner, "завершил свою работу")
		// 	return
		// default:
		// 	fmt.Println("Я шахтёр ", numberMiner, "начал добывать уголь")

		// 	time.Sleep(1 * time.Second)

		// 	transferPoint <- power

		// 	fmt.Println("Я шахтёр ", numberMiner, "добыл уголь и передал. Количество: ", power)
		// }
	}
}

func PoolMiner(ctx context.Context, minerCount int) <-chan int { //<-chan int - канал только на чтение

	wg := &sync.WaitGroup{}

	minerTransferPoint := make(chan int)

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(wg, ctx, minerTransferPoint, i, rand.IntN(3)+10)
	}

	//Добавили данный функционал в целях того, если мы будем возвращать в main уже закрытый канал, тогда мы не сможем прочитать его содержимое, но в связи с тем,
	//что нам это необходимо, мы запускаем отдельную горутину, как только доходим до wg.Wait() она блокируется, соответственно мы ждём как все wg.Done() отработают
	//и как только это случится, мы закроем канал и соответственно в main, при попытке считать новое значение из канала, цикл поймёт под капотом, что канал закрыт, считывается
	//значение по умолчанию, соответственно можно заканчивать с чтением от туда. Там поток заблокируется, если будет не в отдельной горутине, и будет ждать, пока у нас
	//тут всё отработает или же заблокируется, но в нашем сценарии, канал закроется, как только все горутины отпишут в него значения по очереди
	go func() {
		wg.Wait()
		close(minerTransferPoint)
	}()

	return minerTransferPoint
}
