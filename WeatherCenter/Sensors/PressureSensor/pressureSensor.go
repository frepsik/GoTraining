package pressuresensor

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PressureSensor struct {
	stop         map[int]chan struct{}
	mtx          sync.Mutex
	ctx          context.Context
	countSensors int
}

// Конструктор
func NewPressureSensor(_ctx context.Context, _countSensors int) *PressureSensor {
	return &PressureSensor{
		stop:         make(map[int]chan struct{}),
		ctx:          _ctx,
		countSensors: _countSensors,
	}
}

// Метод осуществляющий измерение атмосферного давления в определённой координате
func (ps *PressureSensor) collection(
	wg *sync.WaitGroup,
	ctx context.Context,
	transferPoint chan<- PressureSensorType,
	lat float64,
	lot float64,
	numberSensor int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сбор атмосферного давления датчиком №", numberSensor, "завершён")
			return
		case <-ps.stop[numberSensor]:
			fmt.Println("Сбор атмосферного давления датчиком №", numberSensor, "завершили вручную")
			return
		default:
			fmt.Println("Датчик атмосферного давления №", numberSensor, "начал производить замер в координатах:", lat, ",", lot)
			time.Sleep(1 * time.Second)

			atmosphericPressure := rand.Intn(751) + 100

			data := PressureSensorType{
				AtmosphericPressure: int64(atmosphericPressure),
				X:                   lat,
				Y:                   lot,
			}

			transferPoint <- data
			fmt.Println("Время:", time.Now(), "Датчик Атм.Дав №", numberSensor, "в координатах", lat, ",", lot, "считал атмосферное давление:", atmosphericPressure)

		}
	}
}

// Функция для остановки работы конкретного датчика
func (ps *PressureSensor) StopCollection(numberSensor int) {

	if ps.stop == nil {
		panic("Use NewPressureSensor()")
	}

	ps.mtx.Lock()
	ch, ok := ps.stop[numberSensor]
	ps.mtx.Unlock()
	if ok {
		close(ch)
		delete(ps.stop, numberSensor)
	}

}

// Произвольная генерация координат где делается замер
func (ps *PressureSensor) generatingCoordinates() (x float64, y float64) {
	return rand.Float64()*180 - 90, rand.Float64()*360 - 180
}

// Функция, где мы осуществляем работу датчиков и далее передаём значение в функцию вызова этой, также присутствует контроль над работой каждого датчика
func (ps *PressureSensor) PoolSensor() <-chan PressureSensorType {

	if ps.stop == nil {
		panic("Use NewPressureSensor()")
	}

	psTransferPoint := make(chan PressureSensorType)

	wg := &sync.WaitGroup{}

	for i := 1; i <= ps.countSensors; i++ {
		wg.Add(1)
		lat, lot := ps.generatingCoordinates()
		ps.stop[i] = make(chan struct{})

		go ps.collection(wg, ps.ctx, psTransferPoint, lat, lot, i)
	}

	//Как только завершится
	go func() {
		wg.Wait()
		close(psTransferPoint)
	}()

	return psTransferPoint
}
