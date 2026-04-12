package main

import (
	"context"
	weathercenter "goTraining/WeatherCenter"
	humiditysensor "goTraining/WeatherCenter/Sensors/HumiditySensor"
	pressuresensor "goTraining/WeatherCenter/Sensors/PressureSensor"
	seismicsensors "goTraining/WeatherCenter/Sensors/SeismicSensors"
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
}
