package service

import "errors"

//Файл с ошибками, в случае если бизнес логика стала очень сложной или надо скрыть ошибки работы сервиса
//которые нет необходимости показывать пользователю

var ErrInternalServer = errors.New("InternalServerErrorCustom")
