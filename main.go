package main

import (
	"context"
	"fmt"
	"goTraining/miner"
	"goTraining/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	minerContext, minerClosed := context.WithCancel(context.Background())
	postmanContext, postmanClosed := context.WithCancel(context.Background())

	mtx := sync.Mutex{}
	var coal atomic.Int64

	var mails []string

	startTime := time.Now()

	//Шахтёры работают три секунды, после все завершают свою работу
	go func() {
		time.Sleep(3 * time.Second)
		minerClosed()
		fmt.Println("---->>>Шахтёры завершили свою работу")
	}()

	//Почтальоны работают всего 6 секунд
	go func() {
		time.Sleep(7 * time.Second)
		postmanClosed()
		fmt.Println("---->>>Почтальоны завершили свою работу")
	}()

	//Запускаем весь функционал по их работе
	//при количество в 100, или более, менее, время работы одинаковое, примерно 6 секнуд (провенрял 3, 15, 100 одновременно выполняющих работу)
	minerTransferPoint := miner.PoolMiner(minerContext, 3)
	postmanTransferPoint := postman.PoolPostman(postmanContext, 3)

	wg := sync.WaitGroup{}

	//Добавляем считывание из каналов в отдельные горутины, чтобы одновременно пытались читать, также как и параллельно у нас все записывают в канал
	wg.Add(1)
	go func() {
		defer wg.Done()

		for value := range minerTransferPoint {
			coal.Add(int64(value))
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for value := range postmanTransferPoint {
			mtx.Lock()
			mails = append(mails, value)
			mtx.Unlock()
		}
	}()

	//Здесь мы будем ждать, то есть поток заблокируется до момента, когда все шахтёры и почтальоны закончат свою работу, канал закроется, цикл, завершится и горутина на считывание
	//завершит своё существование и только в этот момент, мы пойдём дальше выводить остальное содержимое.
	//Deadlock не будет потому что, будут потоки, которые работают, не все горутины заблокированы
	wg.Wait()

	fmt.Println("Количество добытого угля: ", coal.Load())

	mtx.Lock()
	fmt.Println("Количество доставленных писем: ", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время: ", time.Since(startTime))
}
