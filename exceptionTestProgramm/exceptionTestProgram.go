package exceptiontestprogramm

import (
	"errors"
	entities "goTraining/exceptionTestProgramm/Entities"
)

type DataBase struct {
}

func (db *DataBase) AddUser(user *entities.User) error {
	if user != nil {
		if user.Id == 0 {
			return errors.New("Поле Id было передано пустым")
		}
		if user.Name == "" {
			return errors.New("Поле Name было передано пустым")
		}
		if user.Email == "" {
			return errors.New("Поле Email было передано пустым")
		}
		if user.Number == "" {
			return errors.New("Поле Number было передано пустым")
		}
		if user.Age < 3 || user.Age > 90 {
			return errors.New("Поле Age было передано пустым")
		}
		return nil
	} else {
		return errors.New("Данный объект пустой, необходимо заполнить")
	}

}

func (db DataBase) TestFunc() {}
