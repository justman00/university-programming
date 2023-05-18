package models

import (
	"fmt"
)

// Table is a common interface for all kinds of tables
type Table interface {
	GetInfo() string
}

// RestaurantFactory is an abstract factory that can create different types of tables
type RestaurantFactory interface {
	CreateSquareTable() Table
	CreateRoundTable() Table
}

// DinerSquareTable
type DinerSquareTable struct{}

func (d *DinerSquareTable) GetInfo() string {
	return "Diner square table"
}

// DinerRoundTable
type DinerRoundTable struct{}

func (d *DinerRoundTable) GetInfo() string {
	return "Diner round table"
}

// DinerFactory
type DinerFactory struct{}

func (d *DinerFactory) CreateSquareTable() Table {
	return &DinerSquareTable{}
}

func (d *DinerFactory) CreateRoundTable() Table {
	return &DinerRoundTable{}
}

// FineDiningSquareTable
type FineDiningSquareTable struct{}

func (f *FineDiningSquareTable) GetInfo() string {
	return "Fine dining square table"
}

// FineDiningRoundTable
type FineDiningRoundTable struct{}

func (f *FineDiningRoundTable) GetInfo() string {
	return "Fine dining round table"
}

// FineDiningFactory
type FineDiningFactory struct{}

func (f *FineDiningFactory) CreateSquareTable() Table {
	return &FineDiningSquareTable{}
}

func (f *FineDiningFactory) CreateRoundTable() Table {
	return &FineDiningRoundTable{}
}

// https://refactoring.guru/design-patterns/abstract-factory/go/example#example-0
// the Abstract Factory pattern is about creating families of related or interdependent objects.
func GetRestaurantFactory(kind string) (RestaurantFactory, error) {
	if kind == "diner" {
		return &DinerFactory{}, nil
	} else if kind == "fine_dining" {
		return &FineDiningFactory{}, nil
	}

	return nil, fmt.Errorf("invalid restaurant kind: %s", kind)
}
