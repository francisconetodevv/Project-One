package model

import (
	"fmt"
	"strings"
	"time"
)

/* 1. Interface Carro */
type Car interface {
	CarModels() string
	CarCategory() string
	CarMaintainence() float64
	CarReview() bool
}

/* 2. Tipos concretos da Classe */
type EletricCar struct {
	CarModel       string
	Review         time.Time
	BatteryCapacit float64
	Autonomy       float64
}

/* 2. Tipos concretos da Classe */
type CombustionCar struct {
	CarModel   string
	Review     time.Time
	PowerHorse float64
	Fuel       string
}

/* 3. Implementação dos métodos da Interface para EletricCar */
func (c EletricCar) CarModels() string {
	return c.CarModel
}

func (c EletricCar) CarCategory() string {
	return "Eletric"
}

func (c EletricCar) CarMaintainence() float64 {
	/* It cost R$ 50 per 1000 Km of autonomy */
	return (c.Autonomy / 1000) * 50
}

func (c EletricCar) CarReview() bool {
	return time.Since(c.Review) > 365*24*time.Hour
}

/* 4. Implementando os métodos da Interface para CombustionCar */
func (c CombustionCar) CarModels() string {
	return c.CarModel
}

func (c CombustionCar) CarCategory() string {
	return "Combustion"
}

func (c CombustionCar) CarMaintainence() float64 {
	/* It cost R$ 0,5 per Power */
	return c.PowerHorse * 0.5 * 1.2
}

func (c CombustionCar) CarReview() bool {
	return time.Since(c.Review) > 180*24*time.Hour
}

/* 5. Struct Genérica para a Frota*/
/* Serve para: */
type Frota[T Car] struct {
	Vehicle []struct {
		Vehicle    T
		Bought     time.Time
		Kilometers float64
	}
}

/* 6. Métodos da Frota */
func (f *Frota[T]) AddVehicle(v T, data time.Time, km float64) {
	f.Vehicle = append(f.Vehicle, struct {
		Vehicle    T
		Bought     time.Time
		Kilometers float64
	}{v, data, km})
}

func (f *Frota[T]) CostTotalMaintainence() float64 {
	total := 0.0
	for _, value := range f.Vehicle {
		total += value.Vehicle.CarMaintainence()
	}

	return total
}

func (f *Frota[T]) FilterCategory(Category string) []T {
	var result []T
	for _, value := range f.Vehicle {
		if strings.EqualFold(value.Vehicle.CarCategory(), Category) {
			result = append(result, value.Vehicle)
		}
	}

	return result
}

/* 7. Função utilitária genérica */
func Report[T Car](f Frota[T]) string {
	report := strings.Builder{}
	countCategory := make(map[string]int)
	reviewDelay := 0

	for _, value := range f.Vehicle {
		cat := value.Vehicle.CarCategory()
		countCategory[cat]++

		if value.Vehicle.CarReview() {
			reviewDelay++
		}
	}

	report.WriteString("Relatório da frota:\n")

	for cat, total := range countCategory {
		report.WriteString(fmt.Sprintf("- %s: %d Veiculos\n", cat, total))
	}

	return report.String()
}
