package humiditysensor

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Humiditysensor struct {
	stop         map[int]chan struct{}
	mtx          sync.Mutex
	ctx          context.Context
	countSensors int
}

// Конструктор
func NewHumiditysensor(_ctx context.Context, _countSensors int) *Humiditysensor {
	return &Humiditysensor{
		stop:         make(map[int]chan struct{}),
		ctx:          _ctx,
		countSensors: _countSensors,
	}
}

// Метод осуществляющий измерение влажности воздуха в определённой координате
func (hs *Humiditysensor) collection(
	wg *sync.WaitGroup,
	ctx context.Context,
	transferPoint chan<- any,
	lat float64,
	lot float64,
	numberSensor int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сбор влажности воздуха датчиком №", numberSensor, "завершён")
			return
		case <-hs.stop[numberSensor]:
			fmt.Println("Сбор влажности воздуха датчиком №", numberSensor, "завершили вручную")
			return
		default:
			fmt.Println("Датчик влажности воздуха №", numberSensor, "начал производить замер в координатах:", lat, ",", lot)
			time.Sleep(1 * time.Second)

			humidity := rand.Intn(100) + 1

			data := HumiditysensorType{
				Humidity: int16(humidity),
				X:        lat,
				Y:        lot,
			}

			transferPoint <- data
			fmt.Println("Время:", time.Now(), "Датчик Влаж.Возд №", numberSensor, "в координатах", lat, ",", lot, "считал атмосферное давление:", humidity)

		}
	}
}

// Функция для остановки работы конкретного датчика
func (hs *Humiditysensor) StopCollection(numberSensor int) {

	if hs.stop == nil {
		panic("Use NewPressureSensor()")
	}

	hs.mtx.Lock()
	ch, ok := hs.stop[numberSensor]
	hs.mtx.Unlock()
	if ok {
		close(ch)
		delete(hs.stop, numberSensor)
	}

}

// Произвольная генерация координат где делается замер
func (hs *Humiditysensor) generatingCoordinates() (x float64, y float64) {
	return rand.Float64()*180 - 90, rand.Float64()*360 - 180
}

// Функция, где мы осуществляем работу датчиков и далее передаём значение в функцию вызова этой, также присутствует контроль над работой каждого датчика
func (hs *Humiditysensor) PoolSensor() <-chan any {

	if hs.stop == nil {
		panic("Use NewPressureSensor()")
	}

	hsTransferPoint := make(chan any)

	wg := &sync.WaitGroup{}

	for i := 1; i <= hs.countSensors; i++ {
		wg.Add(1)
		lat, lot := hs.generatingCoordinates()
		hs.stop[i] = make(chan struct{})

		go hs.collection(wg, hs.ctx, hsTransferPoint, lat, lot, i)
	}

	//Как только завершится
	go func() {
		wg.Wait()
		close(hsTransferPoint)
	}()

	return hsTransferPoint
}
