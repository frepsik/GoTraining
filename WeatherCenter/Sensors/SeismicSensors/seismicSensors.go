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
	transferPoint chan<- any,
	lat float64,
	lot float64,
	numberSensor int,
	stopChan <-chan struct{},
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сбор сейсмической активности датчиком №", numberSensor, "завершён")
			return
		case <-stopChan:
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
			fmt.Println("Время:", time.Now(), "\nДатчик Сейс.Актив №", numberSensor, "в координатах", lat, ",", lot, "считал атмосферное давление:", seismicActivity)
			fmt.Println()
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
func (ss *SeismicSensor) PoolSensor() <-chan any {

	if ss.stop == nil {
		panic("Use NewPressureSensor()")
	}

	ssTransferPoint := make(chan any)

	wg := &sync.WaitGroup{}

	for i := 1; i <= ss.countSensors; i++ {
		wg.Add(1)
		lat, lot := ss.generatingCoordinates()

		//Будем передавать отдельный канал, потому что если через map обращаться, будет гонка данных, и использовать mtx перед select там не целесообразно, через
		//пустой канал будет более эффективно по времени
		stopChan := make(chan struct{})

		ss.mtx.Lock()
		ss.stop[i] = stopChan
		ss.mtx.Unlock()

		go ss.collection(wg, ss.ctx, ssTransferPoint, lat, lot, i, stopChan)
	}

	//Как только завершится
	go func() {
		wg.Wait()
		close(ssTransferPoint)
	}()

	return ssTransferPoint
}
