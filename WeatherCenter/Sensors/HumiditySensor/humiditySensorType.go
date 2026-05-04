package humiditysensor

// Тип необходимый для сбора данных по датчику влажности в воздухе
type HumiditysensorType struct {
	Humidity int16 //Процент, в теории можно было бы сделать проверку на процент не менее 0 к примеру
	X        float64
	Y        float64
}
