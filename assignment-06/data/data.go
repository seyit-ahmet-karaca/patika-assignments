package data

import (
	internalError "karaca/error"
)

type IData interface {
	Insert(username string, balance int)
	Update(username string, balance int) error
	GetByUsername(username string) (int, error)
	GetAll() map[string]int
}

type Data struct {
	dt map[string]int
}


func (d *Data)  Insert(username string, balance int) {
	d.dt[username] = balance

}

func (d *Data) Update(username string, balance int) error {
	if _, err := d.GetByUsername(username); err == nil {
		d.dt[username] = balance
		return nil
	} else {
		return err
	}
}

func (d *Data) GetByUsername(username string) (int, error) {
	if userBalance, ok := d.dt[username]; ok {
		return userBalance, nil
	}

	return 0, internalError.GetUsernameNotFound()
}

func (d *Data) GetAll() map[string]int {
	return d.dt
}

func NewData() IData {
	return &Data{dt: map[string]int{}}
}