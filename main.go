package main

import (
	"errors"
	"fmt"
	"time"
)


type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int 

func NewMedicine(name, manufacturer string, manufactureDate time.Time, shelfLife int) (Medicine, error) {
	m := Medicine{}
	if err := m.SetName(name); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return Medicine{}, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return Medicine{}, err
	}
	return m, nil
}

