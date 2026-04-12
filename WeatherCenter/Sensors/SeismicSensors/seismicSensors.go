package seismicsensors

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SeismicSensor struct {
	stop         map[int]chan struct{}
	mtx          sync.Mutex
	ctx          context.Context
	countSensors int
}

// Конструктор
func NewSeismicSensor(_ctx context.Context, _countSensors int) *SeismicSensor {
	return &SeismicSensor{
		stop:         make(map[int]chan struct{}),
		ctx:          _ctx,
		countSensors: _countSensors,
	}
}

// Метод осуществляющий измерение сейсмической активности в определённой координате
func (ss *SeismicSensor) collection(
	wg *sync.WaitGroup,
	ctx context.Context,
	transferPoint chan<- SeismicSensorsType,
	lat float64,
	lot float64,
	numberSensor int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сбор сейсмической активности датчиком №", numberSensor, "завершён")
			return
		case <-ss.stop[numberSensor]:
			fmt.Println("Сбор сейсмической активности датчиком №", numberSensor, "завершили вручную")
			return
		default:
			fmt.Println("Датчик сейсмической активности №", numberSensor, "начал производить замер в координатах:", lat, ",", lot)
			time.Sleep(1 * time.Second)

			seismicActivity := rand.Intn(12) + 1

			data := SeismicSensorsType{
				SeismicActivity: int8(seismicActivity),
				X:               lat,
				Y:               lot,
			}

			transferPoint <- data
			fmt.Println("Время:", time.Now(), "Датчик Сейс.Актив №", numberSensor, "в координатах", lat, ",", lot, "считал атмосферное давление:", seismicActivity)

		}
	}
}

// Функция для остановки работы конкретного датчика
func (ss *SeismicSensor) StopCollection(numberSensor int) {

	if ss.stop == nil {
		panic("Use NewPressureSensor()")
	}

	ss.mtx.Lock()
	ch, ok := ss.stop[numberSensor]
	ss.mtx.Unlock()
	if ok {
		close(ch)
		delete(ss.stop, numberSensor)
	}

}

// Произвольная генерация координат где делается замер
func (ss *SeismicSensor) generatingCoordinates() (x float64, y float64) {
	return rand.Float64()*180 - 90, rand.Float64()*360 - 180
}

// Функция, где мы осуществляем работу датчиков и далее передаём значение в функцию вызова этой, также присутствует контроль над работой каждого датчика
func (ss *SeismicSensor) PoolHumiditysensor() <-chan SeismicSensorsType {

	if ss.stop == nil {
		panic("Use NewPressureSensor()")
	}

	ssTransferPoint := make(chan SeismicSensorsType)

	wg := &sync.WaitGroup{}

	for i := 1; i <= ss.countSensors; i++ {
		wg.Add(1)
		lat, lot := ss.generatingCoordinates()
		ss.stop[i] = make(chan struct{})

		go ss.collection(wg, ss.ctx, ssTransferPoint, lat, lot, i)
	}

	//Как только завершится
	go func() {
		wg.Wait()
		close(ssTransferPoint)
	}()

	return ssTransferPoint
}
