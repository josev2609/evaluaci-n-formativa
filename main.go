package main

import (
	"errors"
	"fmt"
	"time"
)


type Medicine struct {name string
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
func (m *Medicine) Name() string {
	return m.name
}

func (m *Medicine) SetName(name string) error {
    if name == "" {
        return errors.New("el nombre no puede estar vacío") 
    }
    m.name = name 
    return nil    
}
func (m *Medicine) Manufacturer() string {
	return m.manufacturer
}
func (m *Medicine) SetManufactureDate(manufactureDate time.Time) error {
	if manufactureDate.IsZero() {
		return errors.New("la fecha de fabricación no es válida")
	}
	m.manufactureDate = manufactureDate
	return nil
}

func (m *Medicine) ShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("la vida útil debe ser un número positivo mayor a 0")
	}
	m.shelfLife = shelfLife
	return nil
}
func (m Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}
