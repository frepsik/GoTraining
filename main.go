package main

import (
	"context"
	"fmt"
	weathercenter "goTraining/WeatherCenter"
	humiditysensor "goTraining/WeatherCenter/Sensors/HumiditySensor"
	pressuresensor "goTraining/WeatherCenter/Sensors/PressureSensor"
	seismicsensors "goTraining/WeatherCenter/Sensors/SeismicSensors"
	"time"
)

func main() {
	seismicSensorContext, seismicSensorCancell := context.WithCancel(context.Background())
	humiditysensorContext, humiditysensorCancell := context.WithCancel(context.Background())
	pressureSensorContext, pressureSensorCancell := context.WithCancel(context.Background())

	pressureSensor := pressuresensor.NewPressureSensor(pressureSensorContext, 3)
	seismicSensor := seismicsensors.NewSeismicSensor(seismicSensorContext, 3)
	humiditysensor := humiditysensor.NewHumiditysensor(humiditysensorContext, 3)

	sensors := map[string]weathercenter.Sensors{
		"Датчик давления":                pressureSensor,
		"Датчик влажности":               humiditysensor,
		"Датчик сейсмической активности": seismicSensor,
	}

	startTime := time.Now()

	//Завершим работу сейсмических датчиков через 5 секунды
	go func() {
		time.Sleep(5 * time.Second)
		seismicSensorCancell()
		fmt.Println("---------->>>>>>>Сейсмический датчики завершают свою работу")
		fmt.Println()
	}()

	//Завершим работу датчик влажности через 8 секнуды работы
	go func() {
		time.Sleep(8 * time.Second)
		humiditysensorCancell()
		fmt.Println("---------->>>>>>>Датчики влажности завершают свою работу")
		fmt.Println()
	}()

	//Завершим работу датчиков давления через 12 секнуды работы
	go func() {
		time.Sleep(12 * time.Second)
		pressureSensorCancell()
		fmt.Println("---------->>>>>>>Датчики давления завершают свою работу")
		fmt.Println()
	}()

	resultRunSensors := weathercenter.WeatherCenterRun(sensors)

	for _, sensor := range resultRunSensors {
		fmt.Println("Наименование:", sensor.NameSensor, "\nСобраная информация:", sensor.Value)
		fmt.Println()
	}

	fmt.Println("Затраченное время: ", time.Since(startTime))
}
